package cyberia

import (
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type HandleFunc func(ctx *Context)

type Handler struct {
	Path   string
	Method string
	Handle HandleFunc
}

type App struct {
	Opts
	Handlers []Handler
}

type Context struct {
	*fasthttp.RequestCtx
	app *App
}

func (c *Context) JSON(v interface{}) error {
	res, err := c.app.JSONMarshal(v)
	if err != nil {
		return err
	}
	c.Response.SetBodyRaw(res)
	c.Response.Header.SetContentType("application/json")
	return nil
}

func (a *App) handler(ctx *fasthttp.RequestCtx) {
	c := Context{ctx, a}
	for _, h := range a.Handlers {
		if h.Method == string(ctx.Method()) && h.Path == string(ctx.Path()) {
			h.Handle(&c)
		}
	}
}

func (a *App) GET(path string, handle HandleFunc) {
	a.Handlers = append(a.Handlers, Handler{
		Path:   path,
		Method: fasthttp.MethodGet,
		Handle: handle,
	})
}

func (a *App) POST(path string, handle HandleFunc) {
	a.Handlers = append(a.Handlers, Handler{
		Path:   path,
		Method: fasthttp.MethodPost,
		Handle: handle,
	})
}

func (a *App) Listen(address string) {
	fasthttp.ListenAndServe(address, a.handler)
}

func defaultOpts() Opts {
	return Opts{
		JSONMarshal:   sonic.Marshal,
		JSONUnmarshal: sonic.Unmarshal,
	}
}

// constructor

type OptFunc func(*Opts)

type Opts struct {
	JSONMarshal   func(interface{}) ([]byte, error)
	JSONUnmarshal func([]byte, interface{}) error
}

func New(opts ...OptFunc) *App {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &App{Opts: o}
}
