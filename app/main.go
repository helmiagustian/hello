package main

import (
	"github.com/helmiagustian/hello"
	"github.com/helmiagustian/hello/pkg/handlers"
)

func main() {
	serverStart()
}

func serverStart() {
	helloServer := hello.NewHello()
	defaultHandler := new(handlers.DefaultHandler)
	helloServer.SetupRouters(defaultHandler)
	helloServer.ListenAndServe()
}
