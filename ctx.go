package cyberia

import "github.com/valyala/fasthttp"

type Context struct {
	*fasthttp.RequestCtx
	app *App
}

func (c *Context) JSON(v interface{}) error {
	res, err := c.app.jsonMarshal(v)
	if err != nil {
		return err
	}
	c.Response.SetBodyRaw(res)
	c.Response.Header.SetContentType("application/json")
	return nil
}
