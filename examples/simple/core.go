package main

import (
	"github.com/iyear/go-plugin-grpc/core"
	"log"
)

func main() {
	c := core.New("my token",
		core.WithInterfaces(map[string][]string{
			"Math": { // the name of the interface
				"Plus", "Multiply", // the names of the methods
			}, // you can add more interfaces, but plugin must implement only one interface
		}),
		core.WithPort(14000)) // default port: 13000

	// launch the core
	go func() {
		if err := c.Serve(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	c.Call("MyPlugin", "v1", "Plus", map[string]interface{}{
		"A": 2,
		"B": 3,
	})

	select {}
}
