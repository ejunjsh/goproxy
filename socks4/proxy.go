package socks4

import (
	"github.com/ejunjsh/goproxy/tcp"
	"net"
	"log"
	"fmt"
	"io"
	"strconv"
)

type Proxy struct {
	tcp.Proxy
}

func (tcpproxy *Proxy) Run(){
	l,err:= net.Listen("tcp",tcpproxy.ProxyAdr)
	if err!=nil{
		log.Println(err)
		return
	}
	tcpproxy.Listener=l
	fmt.Println("server listens on ",tcpproxy.ProxyAdr)
	for{
		client,err:=tcpproxy.Listener.Accept()
		if err !=nil {
			log.Println(err)
			continue
		}
		go tcpproxy.serve(client)
	}

}

func (tcpproxy *Proxy) serve(client net.Conn){
    var b [1024]byte
	_, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	if b[0] == 0x04 { //only for socks4
		var host, port string
		host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		port = strconv.Itoa(int(b[2])<<8 | int(b[3]))
		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			log.Println(err)
			return
		}
		client.Write([]byte{0x00, 0x5a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //response to client connection is done.
		go func() {
			io.Copy(client, server)
			server.Close()
			client.Close()
		}()
		io.Copy(server, client)
		client.Close()
		server.Close()
	} else {
	    client.Close()
	}
}