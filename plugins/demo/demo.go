//go:build tinygo.wasm

package main

import (
	"context"
	"github.com/knqyf263/go-plugin/types/known/emptypb"
	"github.com/amlwwalker/go-grpc-plugins/pb/greenfinch";
	"fmt"
)

func main() {
	greenfinch.RegisterPlugin(GreenfinchPlugin{})
}

type GreenfinchPlugin struct{}

func (m GreenfinchPlugin) Init(ctx context.Context, request greenfinch.PluginInit) (greenfinch.PluginReply, error) {
	hostFunctions := greenfinch.NewGreenfinchFunctions()

	// Logging via the host function
	hostFunctions.Log(ctx, greenfinch.LogRequest{
		Message: "Configuring the plugin",
	})

	resp, err := hostFunctions.RequestSign(ctx, greenfinch.Payload{Payload: "abc123", Meta: "some more data"})
	if err != nil {
		return greenfinch.PluginReply{PluginId: "demo", Message: "failed"}, err
	}


	fmt.Printf("request %+v\r\nresponse %+v\r\n", request, resp)
	return greenfinch.PluginReply{
		PluginId: "demo",
		Message: resp.SignedPayload,
	}, nil
}

func (m GreenfinchPlugin) Content(ctx context.Context, e emptypb.Empty) (greenfinch.PluginContent, error) {
	return greenfinch.PluginContent{
		Content: "<html><body></body></html>",
	}, nil
}
func (m GreenfinchPlugin) Setting(ctx context.Context, e emptypb.Empty) (greenfinch.PluginContent, error) {
	return greenfinch.PluginContent{}, nil
}
