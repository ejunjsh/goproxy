# goproxy
[![Build Status](https://travis-ci.org/ejunjsh/goproxy.svg?branch=master)](https://travis-ci.org/ejunjsh/goproxy)

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

a proxy with go,supports tcp,http,socks4/5

## install
````
go get github.com/ejunjsh/goproxy
````
## run with tcp
````
$GOPATH/bin/goproxy -a :8090 -t tcp -b [backend_ip:port]
````
## run with http
````
$GOPATH/bin/goproxy -a :8090 -t http
````
## run with socket4
````
$GOPATH/bin/goproxy -a :8090 -t socks4
````
## run with socket5
````
$GOPATH/bin/goproxy -a :8090 -t socks5
````
