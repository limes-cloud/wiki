package main

import (
	"h1/internal/handler"
	"os"

	"github.com/limes-cloud/kratos"
	v1 "h1/api/v1"
	srcConf "h1/config"
	"github.com/limes-cloud/kratos/config"
	"github.com/limes-cloud/kratos/config/file"
	"github.com/limes-cloud/kratos/log"
	"github.com/limes-cloud/kratos/middleware/tracing"
	"github.com/limes-cloud/kratos/transport/grpc"
	"github.com/limes-cloud/kratos/transport/http"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

func main() {
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Config(file.NewSource("config/config.yaml")),
		kratos.RegistrarServer(RegisterServer),
		kratos.LoggerWith(kratos.LogField{
			"id":      id,
			"name":    Name,
			"version": Version,
			"trace":   tracing.TraceID(),
			"span":    tracing.SpanID(),
		}),
	)

	if err := app.Run(); err != nil {
		log.Errorf("run service fail: %v", err)
	}
}

func RegisterServer(hs *http.Server, gs *grpc.Server, c config.Config) {
	conf := &srcConf.Config{}
	if err := c.ScanKey("business", conf); err != nil {
		panic("business config format error:" + err.Error())
	}

	srv := handler.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}
