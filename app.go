package cyberia

import (
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type App struct {
	AppOpts
	BaseRouter
}

func (a *App) handler(ctx *fasthttp.RequestCtx) {
	c := Context{ctx, a}
	for _, h := range a.handlers {
		if h.Method == string(ctx.Method()) && h.Path == string(ctx.Path()) {
			h.Handle(&c)
		}
	}

}

func (a *App) Listen(address string) {
	fasthttp.ListenAndServe(address, a.handler)
}

// constructor

type AppOptFunc func(*AppOpts)

type AppOpts struct {
	jsonMarshal   func(interface{}) ([]byte, error)
	jsonUnmarshal func([]byte, interface{}) error
}

func defaultAppOpts() AppOpts {
	return AppOpts{
		jsonMarshal:   sonic.Marshal,
		jsonUnmarshal: sonic.Unmarshal,
	}
}

func New(opts ...AppOptFunc) *App {
	o := defaultAppOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &App{AppOpts: o}
}
