# goproxy
supports tcp,http, more proxy modes are on the road.

## intall
go get github.com/ejunjsh/goproxy


## run tcp
$GOPATH/bin/goproxy -a :8090 -t tcp -b [backend_ip:port]


## run http
$GOPATH/bin/goproxy -a :8090 -t http