package main

import (
	"github.com/ejunjsh/goproxy/http"
	"flag"
	"github.com/ejunjsh/goproxy/tcp"
)


type proxy interface {
	Run()
}


func main(){
	var(
		proxyAdr   string
		proxyType  string
		backendAdr string
	)
	flag.StringVar(&proxyAdr, "addr", "", "Network host to listen on.")
	flag.StringVar(&proxyAdr, "a", "", "Network host to listen on.")
	flag.StringVar(&proxyType, "type", "", "Network host to listen on.")
	flag.StringVar(&proxyType, "t", "", "Network host to listen on.")
	flag.StringVar(&backendAdr, "backend", "", "Network host to listen on.")
	flag.StringVar(&backendAdr, "b", "", "Network host to listen on.")

	flag.Parse()

	var p proxy

	switch proxyType {
	case "http":
		p=&http.Proxy{tcp.Proxy{ProxyAdr: proxyAdr}}
	case "tcp":
		fallthrough
	default:
		p=&tcp.Proxy{ProxyAdr: proxyAdr,BackendAdr: backendAdr}
	}

	p.Run()
}
