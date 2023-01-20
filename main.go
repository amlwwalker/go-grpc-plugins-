package main

import (
	"context"
	"fmt"
	"github.com/amlwwalker/go-grpc-plugins/pb/greenfinch"
	"github.com/knqyf263/go-plugin/types/known/emptypb"
	"log"
	"os"
	"path/filepath"
)


type Ex struct {
	plugins map[string]greenfinch.Plugin
}
func main() {
	var e = Ex{
		plugins: make(map[string]greenfinch.Plugin),
	}
	if err := e.run(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	for _, v := range e.plugins {
		content, err := v.Content(ctx, emptypb.Empty{})
		if err != nil {
			fmt.Println("error retrieving content", err)
		}
		fmt.Println("render content ", content.GetContent())
	}

}

func (e Ex) run() error {
	ctx := context.Background()
	p, err := greenfinch.NewPluginPlugin(ctx, greenfinch.PluginPluginOption{})
	if err != nil {
		return err
	}
	defer p.Close(ctx)

	var plugins []string
	filepath.Walk("plugins", func(path string, info os.FileInfo, err error) error  {

		if !info.IsDir() && filepath.Ext(path) == ".wasm" {
			plugins = append(plugins, path)
		}

		return nil
	})

	for _, v := range plugins {
		fmt.Println("v ", v)
		// Pass my host functions that are embedded into the plugin.
		p, err := p.Load(ctx, v, PluginFunctions{})
		e.plugins[v] = p
		init, err := p.Init(ctx, greenfinch.PluginInit{})
		if err != nil {
			return err
		}
		fmt.Println("response from init", init.PluginId, init.Message)
		content, err := p.Content(ctx, emptypb.Empty{})
		if err != nil {
			return err
		}
		fmt.Println("render content ", content.GetContent())
	}
	return nil
}


// myHostFunctions implements greeting.HostFunctions
type PluginFunctions struct{}

// HttpGet is embedded into the plugin and can be called by the plugin.
func (PluginFunctions) RequestSign(ctx context.Context, request greenfinch.Payload) (greenfinch.SignResponse, error) {
	sign := request.GetPayload()

	fmt.Println("request to sign ", sign)
	return greenfinch.SignResponse{SignedPayload: sign}, nil
}

// Log is embedded into the plugin and can be called by the plugin.
func (PluginFunctions) Log(ctx context.Context, request greenfinch.LogRequest) (emptypb.Empty, error) {
	// Use the host logger
	log.Println("logging ", request.GetMessage())
	return emptypb.Empty{}, nil
}
