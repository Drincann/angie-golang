# angie-golang

一个 Web 后端开发框架 based on golang。

```go
package main

import (
    "github.com/Drincann/angie-golang/angie"
)

func main() {
    app := angie.New()
    app.Get("/hello", func(ctx *angie.Context) {
        ctx.ResString("<h1>world</he>")
    }).Listen(8080)
}
```
