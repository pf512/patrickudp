package main

import (
	"github.com/pf512/patudp/server"
	"golang.org/x/net/context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	ch := make(chan string)
	controlSignals := make(chan string)
	server.Server(ctx, "0.0.0.0:4444", ch, controlSignals)
}
