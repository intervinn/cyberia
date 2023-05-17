package cyberia

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type App struct {
	Opts
}

func (a *App) handler(ctx *fasthttp.RequestCtx) {
	fmt.Println(ctx.Method(), string(ctx.Method()))
}

func (a *App) Listen(address string) {
	fasthttp.ListenAndServe(address, a.handler)
}
