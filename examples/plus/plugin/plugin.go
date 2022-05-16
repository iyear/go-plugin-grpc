package main

import (
	"github.com/iyear/go-plugin-grpc/plugin"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	p := plugin.New("MyPlugin", "v1", "123",
		plugin.WithLogLevel(plugin.LogLevelDebug),
		plugin.WithHeartbeat(10*time.Second),
		plugin.WithDialOpts(grpc.WithInsecure()))

	// if handle only plus, plugin can't bind to core because it doesn't impl any interface
	// if handle plus and multiply, plugin can bind to core
	// if handle plus, multiply and echo, plugin can't bind to core because it impls two interfaces
	p.Handle("plus", plus)
	p.Handle("multiply", multiply)
	p.Handle("echo", echo)

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

func multiply(ctx plugin.Context) (map[string]interface{}, error) {
	ctx.L().Info("enter math.v1.multiply")
	args := ctx.Args()
	ctx.L().Info("finish multiply func")

	return map[string]interface{}{
		"V": args["A"].(float64) * args["B"].(float64),
	}, nil
}

func echo(ctx plugin.Context) (map[string]interface{}, error) {
	text := ctx.Args()["Text"].(string)
	//ctx.L().Debugf("echo %s", text)
	return map[string]interface{}{
		"Text": text,
	}, nil
}
