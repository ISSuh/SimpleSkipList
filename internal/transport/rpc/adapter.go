// MIT License

// Copyright (c) 2024 ISSuh

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package rpc

import (
	"context"
	"net"

	"github.com/ISSuh/skvs/internal/transport/handler"
	"github.com/ISSuh/skvs/internal/transport/rpc/proto"

	"google.golang.org/grpc"
)

type RegisterFunc func(server *grpc.Server)

type ServerAdaptor struct {
	proto.UnimplementedStorageServer

	handler handler.Storage
	engine  *grpc.Server
}

func NewServerAdaptor(handler handler.Storage) ServerAdaptor {
	e := grpc.NewServer()

	return ServerAdaptor{
		handler: handler,
		engine:  e,
	}
}

func (s *ServerAdaptor) Register() RegisterFunc {
	return func(server *grpc.Server) {
		proto.RegisterStorageServer(server, s)
	}
}

func (s *ServerAdaptor) Run(listener net.Listener) error {
	if err := s.engine.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *ServerAdaptor) Heartbeat(context.Context, *proto.Ping) (*proto.Pong, error) {
	return nil, nil
}

func (s *ServerAdaptor) Get(context.Context, *proto.Request) (*proto.Response, error) {
	return nil, nil
}

func (s *ServerAdaptor) Set(context.Context, *proto.Request) (*proto.Response, error) {
	return nil, nil
}

func (s *ServerAdaptor) Delete(context.Context, *proto.Request) (*proto.Response, error) {
	return nil, nil
}
