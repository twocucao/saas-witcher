package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/twocucao/saas-witcher/internal/conf"
	"github.com/twocucao/saas-witcher/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	http.Middleware(
		middleware.Chain(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(),
		),
	)
	//srv.HandlePrefix("/", v1.NewGreeterHandler(greeter, m))
	return srv
}
