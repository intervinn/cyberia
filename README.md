# cyberia

This is a simple abstraction of <a href="https://github.com/valyala/fasthttp"></a> to simplify api development. 

Cyberia has been inspired by Fiber, a library inspired by Express that runs on fasthttp aswell.

This library uses <a href="https://github.com/bytedance/sonic">sonic</a>, a fast json library for Go, implemented in Assembly, but you can change it to any other library that you want.

```go
package main

import "github.com/intervinn/cyberia"

func main() {
    app := cyberia.New()

    router := app.Router(cyberia.WithCustomPrefix(""))

    router.GET("/", func(ctx *cyberia.Context) {
        ctx.JSON(map[string]string{
            "message": "hello, world!",
        })
    })

    app.RegisterRouter(router)

    app.Listen(":8080")
}
```
