package main

import (
	"fmt"
	"github.com/limes-cloud/kratos/contrib/config/configure"
	"hello/internal/handler"
	"os"

	"github.com/limes-cloud/kratos"
	"github.com/limes-cloud/kratos/config"
	"github.com/limes-cloud/kratos/log"
	"github.com/limes-cloud/kratos/middleware/tracing"
	"github.com/limes-cloud/kratos/transport/grpc"
	"github.com/limes-cloud/kratos/transport/http"
	_ "go.uber.org/automaxprocs"
	v1 "hello/api/hello/v1"
	srcConf "hello/config"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "test"
	// Version is the version of the compiled software.
	Version string

	//CHost  string = "101.34.229.39:7082"
	CHost string = "101.34.229.39:6082"

	CToken string = "8A62C35740D5817A3F8C6958C4BE6B2C"

	id, _ = os.Hostname()
)

func main() {
	//file.NewSource("config/config.yaml")
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Config(configure.New(CHost, Name, CToken)),
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

	c.Watch("business", func(key string, value config.Value) {
		fmt.Printf("配置变更 key：%v=>%v", key, value)
		value.Scan(conf)
	})

	srv := handler.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}
