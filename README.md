# goproxy
[![Build Status](https://travis-ci.org/ejunjsh/goproxy.svg?branch=master)](https://travis-ci.org/ejunjsh/goproxy)

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

supports http,socks4/5 at the same time.

## install

    go get github.com/ejunjsh/goproxy


## usage

    usage: goproxy <listen address>
    
## example

if you run below command in your local or remote machine

    goproxy :8090
    
then you can set your proxy address `<localhost or remote ip>:8090` on your browser or something that support socks4/5 or http proxy.

[![](http://idiotsky.top/images3/goproxy.png)](http://idiotsky.top/images3/goproxy.png)

## docker

    sudo docker run --net host  sjj050121014/goproxy:1.0 <listen address>


😄 enjoy
