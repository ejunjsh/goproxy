package tcp

import (
	"net"
	"log"
	"io"
	"fmt"
)

type Proxy struct {
   ProxyAdr    string
	BackendAdr string
	Listener   net.Listener
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

		server,err:=net.Dial("tcp",tcpproxy.BackendAdr)
		if err!=nil{
			log.Println(err)
			return
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


