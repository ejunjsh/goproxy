package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("usage: goproxy <listen address>")
		return
	}

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGHUP)
		<-c
	}()

	Run(os.Args[1])
}
