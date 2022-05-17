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
				"EchoMap2Map", "EchoMap2Bytes", "EchoBytes2Map", "EchoBytes2Bytes",
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

	//call(c, "Plus", map[string]interface{}{
	//	"A": 2,
	//	"B": 3,
	//})
	//call(c, "Multiply", map[string]interface{}{
	//	"A": 2,
	//	"B": 3,
	//})
	//call(c, "EchoMap2Map", map[string]interface{}{
	//	"Text": "hello",
	//})
	//call(c, "EchoMap2Bytes", map[string]interface{}{
	//	"Text": "hello",
	//})
	//call(c, "EchoBytes2Map", []byte("hello"))
	//call(c, "EchoBytes2Bytes", []byte("hello"))

	call(c, "Panic", nil)
	select {}
}

func call(c *core.Core, name string, args interface{}) {
	start := time.Now()
	fmt.Printf("\ncall %s, args: %v\n", name, args)
	// Call Plugin
	r, err := c.Call("MyPlugin", "v1", name, args)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("result: %v, err: %v, time: %v\n", r, err, time.Since(start))
}
