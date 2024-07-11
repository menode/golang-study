package server

import (
	"context"
	v1 "kratos-test/api/realworld/v1"
	"kratos-test/internal/conf"
	"kratos-test/internal/pkg/middleware/auth"
	"kratos-test/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/realworld.v1.Realworld/Login":        {},
		"/realworld.v1.Realworld/Register":     {},
		"/realworld.v1.Realworld/GetArticle":   {},
		"/realworld.v1.Realworld/ListArticles": {},
		"/realworld.v1.Realworld/GetComments":  {},
		"/realworld.v1.Realworld/GetTags":      {},
		"/realworld.v1.Realworld/GetProfile":   {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, jwt *conf.JWT, greeter *service.RealworldService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ErrorEncoder(errorEncoder),
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(jwt.Secret)).Match(NewSkipRoutersMatcher()).Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		)),
	}
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
	v1.RegisterRealworldHTTPServer(srv, greeter)
	return srv
}
