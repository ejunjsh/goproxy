package main

import (
	"github.com/ejunjsh/goproxy/http"
	"flag"
	"github.com/ejunjsh/goproxy/tcp"
	"github.com/ejunjsh/goproxy/https"
	"github.com/ejunjsh/goproxy/socket5"
	"github.com/ejunjsh/goproxy/socket4"
)


type proxy interface {
	Run()
}


func main(){
	var(
		proxyAdr   string
		proxyType  string
		backendAdr string
		cert string
		key string
	)
	flag.StringVar(&proxyAdr, "addr", "", "Network host to listen on.")
	flag.StringVar(&proxyAdr, "a", "", "Network host to listen on.")
	flag.StringVar(&proxyType, "type", "", "Network host to listen on.")
	flag.StringVar(&proxyType, "t", "", "Network host to listen on.")
	flag.StringVar(&backendAdr, "backend", "", "Network host to listen on.")
	flag.StringVar(&backendAdr, "b", "", "Network host to listen on.")
	flag.StringVar(&cert, "cert", "", "Network host to listen on.")
	flag.StringVar(&cert, "c", "", "Network host to listen on.")
	flag.StringVar(&key, "key", "", "Network host to listen on.")
	flag.StringVar(&key, "k", "", "Network host to listen on.")

	flag.Parse()

	var p proxy

	switch proxyType {
	case "http":
		p=&http.Proxy{tcp.Proxy{ProxyAdr: proxyAdr}}
	case "socket5":
		p=&socket5.Proxy{tcp.Proxy{ProxyAdr: proxyAdr}}
	case "socket4":
		p=&socket4.Proxy{tcp.Proxy{ProxyAdr: proxyAdr}}
	case "https":
		p=&https.Proxy{tcp.Proxy{ProxyAdr: proxyAdr},cert,key}
	case "tcp":
		fallthrough
	default:
		p=&tcp.Proxy{ProxyAdr: proxyAdr,BackendAdr: backendAdr}
	}

	p.Run()
}
