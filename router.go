package cyberia

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

func connectPaths(paths ...string) string {
	for i, v := range paths {
		if !strings.HasPrefix(v, "/") {
			paths[i] = "/" + v
		}
		if paths[i] == "/" && paths[i+1] == "/" {
			paths[i] = ""
		}
	}
	return strings.Join(paths, "")
}

type HandleFunc func(ctx *Context)

type Handler struct {
	Path   string
	Method string
	Handle HandleFunc
}

type BaseRouter struct {
	RouterOpts
	handlers []Handler
}

func (b *BaseRouter) GET(path string, handle HandleFunc) {
	b.handlers = append(b.handlers, Handler{
		Path:   path,
		Method: fasthttp.MethodGet,
		Handle: handle,
	})
}

func (b *BaseRouter) POST(path string, handle HandleFunc) {
	b.handlers = append(b.handlers, Handler{
		Path:   path,
		Method: fasthttp.MethodPost,
		Handle: handle,
	})
}

func (b *BaseRouter) RegisterRouter(r *BaseRouter) {
	fmt.Println("registering routes...")
	for _, handler := range r.handlers {
		path := connectPaths(b.prefix, r.prefix, handler.Path)
		b.handlers = append(b.handlers, Handler{
			Path:   path,
			Method: handler.Method,
			Handle: handler.Handle,
		})
	}
}

type RouterOpts struct {
	prefix string
}

type RouterOptFunc func(*RouterOpts)

func defaultRouterOpts() RouterOpts {
	return RouterOpts{
		prefix: "",
	}
}

func WithCustomPrefix(prefix string) RouterOptFunc {
	return func(o *RouterOpts) {
		o.prefix = prefix
	}
}

func (b *BaseRouter) Router(opts ...RouterOptFunc) *BaseRouter {
	o := defaultRouterOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &BaseRouter{RouterOpts: o}
}
