package main

import (
	"context"

	examples "github.com/binchencoder/ease-gateway/proto/examples"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Implements of EchoServiceServer

type echoServer struct{}

// NewEchoServer new echo server
func NewEchoServer() examples.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *echoServer) EchoBody(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
	glog.Info(msg)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}

func (s *echoServer) EchoDelete(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
	glog.Info(msg)
	return msg, nil
}
