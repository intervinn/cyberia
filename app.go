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
type marshalFunc func(interface{}) ([]byte, error)
type unmarshalFunc func([]byte, interface{}) error

type AppOpts struct {
	jsonMarshal   marshalFunc
	jsonUnmarshal unmarshalFunc
}

func defaultAppOpts() AppOpts {
	return AppOpts{
		jsonMarshal:   sonic.Marshal,
		jsonUnmarshal: sonic.Unmarshal,
	}
}

func WithCustomJSON(marshal marshalFunc, unmarshal unmarshalFunc) AppOptFunc {
	return func(o *AppOpts) {
		o.jsonMarshal = marshal
		o.jsonUnmarshal = unmarshal
	}
}

func New(opts ...AppOptFunc) *App {
	o := defaultAppOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &App{AppOpts: o}
}
