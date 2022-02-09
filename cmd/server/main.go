package main

import (
	"os"
	"os/signal"
	"syscall"

	perkbox "github.com/joshuaAllday/perkbox/pkg/server"
)

func main() {
	server, err := perkbox.NewServer()
	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
	defer server.Stop()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan
}
