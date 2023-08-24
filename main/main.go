package main

import (
	"fmt"
	"path"

	"github.com/gardusig/serverfix"
	"github.com/gardusig/serverfix/internal/application"
)

func main() {
	fmt.Println("Starting server...")
	server, err := serverfix.NewServer(
		path.Join("config", "fix.cfg"),
		application.ServerApp{},
	)
	if err != nil {
		panic(err)
	}
	server.Start()
	defer server.Stop()
	fmt.Println("Started server")
	select {}
}
