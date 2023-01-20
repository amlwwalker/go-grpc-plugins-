//go:build tinygo.wasm

package main

import (
	"context"

	greeting "github.com/amlwwalker/go-grpc-plugins/pb/greeting";
)

//protoc --go-plugin_out=./pb/greeting --go-plugin_opt=paths=source_relative --proto_path=./proto ./proto/greeting.proto
//protoc --go-plugin_out=./pb --go-plugin_opt=paths=source_relative --proto_path=./proto ./proto/greeting.proto --go-grpc_out=./pb

//protoc --go-plugin_out=. --go-plugin_opt=paths=source_relative greeting/greet.proto

//protoc --go-plugin_out=. --go_opt=paths=source_relative --go-plugin_opt=paths=source_relative greeting/greet.proto

//doesn't work but... protoc --proto_path=../proto ../proto/lookup.proto --go_out=. --go-grpc_out=.


//tinygo build -o plugin.wasm -scheduler=none -target=wasi --no-debug plugin.go
//tinygo build -o plugins/host/host.wasm -scheduler=none -target=wasi --no-debug plugins/host/host.go
// main is required for TinyGo to compile to Wasm.
func main() {
	greeting.RegisterGreeter(GreetingPlugin{})

}

type GreetingPlugin struct{}

func (m GreetingPlugin) Greet(ctx context.Context, request greeting.GreetRequest) (greeting.GreetReply, error) {
	hostFunctions := greeting.NewHostFunctions()

	// Logging via the host function
	hostFunctions.Log(ctx, greeting.LogRequest{
		Message: "Sending a HTTP request...",
	})

	// HTTP GET via the host function
	resp, err := hostFunctions.HttpGet(ctx, greeting.HttpGetRequest{Url: "http://ifconfig.me"})
	if err != nil {
		return greeting.GreetReply{}, err
	}

	return greeting.GreetReply{
		Message: "Hello, " + request.GetName() + " from " + string(resp.Response),
	}, nil
}
