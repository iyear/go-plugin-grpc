package main

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	p := plugin.New("MyPlugin", "v1", "my token",
		plugin.WithOnPanic(func(plugin *plugin.Plugin, execID uint64, funcName string, err error) {
			fmt.Println("plugin panic", plugin.Name(), execID, funcName, err)
		}),
		plugin.WithDialOpts(grpc.WithTransportCredentials(insecure.NewCredentials())))
	_ = p
}
