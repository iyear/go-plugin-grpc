package main

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	p := plugin.New("MyPlugin", "v1", "123",
		plugin.WithLogLevel(plugin.LogLevelDebug),
		plugin.WithHeartbeat(10*time.Second),
		plugin.WithDialOpts(grpc.WithTransportCredentials(insecure.NewCredentials())))

	// if handle only plus, plugin can't bind to core because it doesn't impl any interface
	// if handle plus and multiply, plugin can bind to core
	// if handle plus, multiply and echo, plugin can't bind to core because it impls two interfaces
	//p.Handle("Plus", Plus)
	//p.Handle("Multiply", Multiply)
	p.Handle("EchoMap2Map", EchoMap2Map)
	p.Handle("EchoMap2Bytes", EchoMap2Bytes)
	p.Handle("EchoBytes2Map", EchoBytes2Map)
	p.Handle("EchoBytes2Bytes", EchoBytes2Bytes)
	p.Handle("Panic", Panic)
	p.Handle("Nil", Nil)

	if err := p.Mount("localhost", 13001); err != nil {
		log.Println(err)
		return
	}

	select {}
}

func Plus(ctx plugin.Context) (interface{}, error) {
	ctx.L().Info("enter math.v1.plus")
	args := ctx.Map()
	ctx.L().Info("finish plus func")

	return map[string]interface{}{
		"V": args.GetInt("A") + args.GetInt("B"),
	}, nil
}

func Multiply(ctx plugin.Context) (interface{}, error) {
	ctx.L().Info("enter math.v1.multiply")
	args := ctx.Map()
	ctx.L().Info("multiply func finish")

	return map[string]interface{}{
		"V": args.GetInt("A") * args.GetInt("B"),
	}, nil
}

func EchoMap2Map(ctx plugin.Context) (interface{}, error) {
	text := ctx.Map().GetString("Text")
	ctx.L().Debugf("echo|arg:map|result:map|arg:%v", text)
	return map[string]interface{}{
		"Text": text,
	}, nil
}

func EchoMap2Bytes(ctx plugin.Context) (interface{}, error) {
	text := ctx.Map().GetString("Text")
	ctx.L().Debugf("echo|arg:map|result:bytes|arg:%v", text)
	return []byte(text), nil
}

func EchoBytes2Map(ctx plugin.Context) (interface{}, error) {
	text := ctx.Bytes()
	ctx.L().Debugf("echo|arg:bytes|result:map|arg:%v", text)
	return map[string]interface{}{
		"Text": string(text),
	}, nil
}

func EchoBytes2Bytes(ctx plugin.Context) (interface{}, error) {
	text := ctx.Bytes()
	ctx.L().Debugf("echo|arg:bytes|result:bytes|arg:%v", text)
	return text, nil
}

func Panic(ctx plugin.Context) (interface{}, error) {
	ctx.L().Debug("I will panic")
	panic(fmt.Errorf("panic info"))
}

func Nil(ctx plugin.Context) (interface{}, error) {
	ctx.L().Debug("I will return nil")
	return nil, nil
}
