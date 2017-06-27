package socket5

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
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	if b[0] == 0x05 { //only for socket5
		//response to client: no need to validation
		client.Write([]byte{0x05, 0x00})
		n, err = client.Read(b[:])
		var host, port string
		switch b[3] {
		case 0x01: //IP V4
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03:                   //domain name
			host = string(b[5: n-2]) //b[4] length of domain name
		case 0x04: //IP V6
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
		}
		port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))
		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			log.Println(err)
			return
		}
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //response to client connection is done.
		go func() {
			io.Copy(client, server)
			server.Close()
			client.Close()
		}()
		io.Copy(server, client)
		client.Close()
		server.Close()
	}
}