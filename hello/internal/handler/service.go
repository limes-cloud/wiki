package handler

import (
	"context"
	"github.com/limes-cloud/kratos"
	"hello/internal/logic"

	v1 "hello/api/hello/v1"
	"hello/config"
)

// Service is a greeter service.
type Service struct {
	v1.UnsafeServiceServer
	logic *logic.Logic
}

func New(conf *config.Config) *Service {
	return &Service{
		logic: logic.NewLogic(conf),
	}
}

// SayHello implements helloworld.Server.
func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return s.logic.SayHello(kratos.MustContext(ctx), in)
}
