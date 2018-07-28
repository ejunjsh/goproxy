package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"strconv"
	"strings"
)


func  Run(addr string){
	l,err:= net.Listen("tcp",addr)
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("server listens on ",addr)
	for{
		client,err:=l.Accept()
		if err !=nil {
			log.Println(err)
			continue
		}
		go serve(client)
	}

}

func serve(client net.Conn){
    var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	if b[0] == 0x05 { //only for socks5
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
	} else if b[0] == 0x04 { //only for socks4
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
	} else if b[0] == 0x43{ //http
		s:=string(b[:])
		ss:=strings.Split(s," ")
		method:=ss[0]
		if method=="CONNECT"{
			host:=ss[1]
			server, err := net.Dial("tcp", host)
			if err != nil {
				log.Println(err)
				return
			}
			success:=[]byte("HTTP/1.1 200 Connection established\r\n\r\n")
			_,err=client.Write(success)
			if err != nil {
				log.Println(err)
				return
			}
			go func() {
				io.Copy(client, server)
				server.Close()
				client.Close()
			}()
			io.Copy(server, client)
			client.Close()
			server.Close()
		}else {
			return
		}
	} else {
		client.Close()
	}
}