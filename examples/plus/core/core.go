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
		core.WithInterfaces([]string{
			"plus",
			// "multiply", if add multiply,plugin can not bind to core
		}, []string{
			"minus", "divide",
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
	call(c)
	select {}
}

func call(c *core.Core) {
	start := time.Now()
	fmt.Println("start call: 2 + 3 = ?")
	// Call Plugin
	r, err := c.Call("math", "v1", "plus", map[string]interface{}{"A": 2, "B": 3})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("call finished. result map: %v, err: %v\n", r, err)
	fmt.Printf("result: %d ,time: %v\n", int(r["V"].(float64)), time.Since(start).String())
}
