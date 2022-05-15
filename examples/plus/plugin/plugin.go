package main

import (
	"github.com/iyear/go-plugin-grpc/plugin"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	p := plugin.New("math", "v1", "123",
		plugin.WithLogLevel(plugin.LogLevelDebug),
		plugin.WithHeartbeat(10*time.Second),
		plugin.WithDialOpts(grpc.WithInsecure()))

	p.Handle("plus", plus)

	if err := p.Mount("localhost", 13001); err != nil {
		log.Println(err)
		return
	}

	select {}
}

func plus(ctx plugin.Context) (map[string]interface{}, error) {
	ctx.L().Info("enter math.v1.plus")
	args := ctx.Args()
	ctx.L().Info("finish plus func")

	return map[string]interface{}{
		"V": args["A"].(float64) + args["B"].(float64),
	}, nil
}
