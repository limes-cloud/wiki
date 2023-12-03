package logic

import (
	"fmt"
	"github.com/limes-cloud/kratos"
	"github.com/limes-cloud/kratos/metadata"
	v1 "hello/api/hello/v1"
	"hello/config"
)

type Logic struct {
	conf *config.Config
}

func NewLogic(conf *config.Config) *Logic {
	return &Logic{
		conf: conf,
	}
}

// SayHello implements helloworld.GreeterServer.
func (l *Logic) SayHello(ctx kratos.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if l.conf.OpenWeb {
		ctx.Logger().Warn("close web")
	}

	if md, ok := metadata.FromServerContext(ctx); ok {
		extra := md.Get("x-md-extra")
		fmt.Print(extra)
	}

	return &v1.HelloReply{
		Message: in.Name,
	}, nil
}
