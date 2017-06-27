# goproxy
supports tcp,http, more proxy modes are on the road.

## install
go get github.com/ejunjsh/goproxy


## run with tcp
$GOPATH/bin/goproxy -a :8090 -t tcp -b [backend_ip:port]

## run with http
$GOPATH/bin/goproxy -a :8090 -t http

## run with socket4
$GOPATH/bin/goproxy -a :8090 -t socket4

## run with socket5
$GOPATH/bin/goproxy -a :8090 -t socket5