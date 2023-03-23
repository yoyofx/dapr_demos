package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello dapr service invocation server")
	service := daprd.NewService(":8223")
	// add a service to service invocation handler
	if err := service.AddServiceInvocationHandler("/echo", echoHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
	if err := service.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}

}

func echoHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf(
		"echo - ContentType:%s, Verb:%s, QueryString:%s, %s",
		in.ContentType, in.Verb, in.QueryString, in.Data,
	)
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
