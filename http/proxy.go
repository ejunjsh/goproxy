package http

import (
	"net"
	"log"
	"io"
	"fmt"
	"bytes"

	"github.com/ejunjsh/goproxy/tcp"
	"net/url"
)

type Proxy struct {
	tcp.Proxy
}

func (httpproxy *Proxy) Run(){
	l,err:= net.Listen("tcp",httpproxy.ProxyAdr)
	if err!=nil{
		log.Println(err)
		return
	}
	httpproxy.Listener=l
	fmt.Println("server listens on ",httpproxy.ProxyAdr)
	for{
		client,err:=httpproxy.Listener.Accept()
		if err !=nil {
			log.Println(err)
			continue
		}
		go httpproxy.serve(client)
	}

}

func (httpproxy *Proxy) serve(client net.Conn){
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &address)

	if method != "CONNECT" {
		u,err:=url.Parse(address)
		if err!=nil{
			log.Println(err)
			return
		}
		if u.Scheme=="http" && u.Port()==""{
			address=u.Host+":80"
		} else if u.Scheme=="https" && u.Port()==""{
			address=u.Host+":443"
		} else {
		    address=u.Host
		}


	}

	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		success:=[]byte("HTTP/1.1 200 Connection established\r\n\r\n")
		client.Write(success)
	} else {
		server.Write(b[:n])
	}

	go func() {
		io.Copy(client,server)
		server.Close()
		client.Close()
	}()
	io.Copy(server,client)
	client.Close()
	server.Close()
}
