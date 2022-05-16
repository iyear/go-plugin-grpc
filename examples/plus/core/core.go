package main

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/core"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	c := core.New("123",
		core.WithLogLevel(core.LogLevelDebug),
		core.WithPort(13001),
		core.WithInterfaces(map[string][]string{
			"Math": {
				"plus", "multiply",
			},
			"String": {
				"echo",
			},
		}),
		core.WithExecReqChSize(5),
		core.WithExecTimeout(time.Second*5),
		core.WithServerOpts(grpc.WriteBufferSize(64*1024), grpc.ReadBufferSize(64*1024)),
		core.WithHealthTimeout(time.Second*15),
	)
	go func() {
		if err := c.Serve(); err != nil {
			log.Fatal(err)
		}
	}()
	// wait for plugin to start
	time.Sleep(time.Second * 8)

	call(c, "plus", map[string]interface{}{
		"A": 2,
		"B": 3,
	})
	call(c, "multiply", map[string]interface{}{
		"A": 2,
		"B": 3,
	})
	call(c, "echo", map[string]interface{}{
		"Text": "hello",
	})
	select {}
}

func call(c *core.Core, name string, args map[string]interface{}) {
	start := time.Now()
	fmt.Printf("call %s, args: %v\n", name, args)
	// Call Plugin
	r, err := c.Call("MyPlugin", "v1", name, args)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("result: %v, err: %v, time: %v\n", r, err, time.Since(start))
}
