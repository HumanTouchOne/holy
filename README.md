# holy

A minimalist Go web framework inspired by Gin.

## Features

- **HTTP Router** - RESTful routing with method chaining
- **Middleware** - Before/after hooks with path matching (prefix, regex, contains, exact)
- **Context** - Request/response helpers (JSON, Query, Cookie, etc.)
- **Components** - Built-in utilities: Log, Cache, JWT, Crypto, ORM

## Quick Start

```go
package main

import (
    "github.com/HumanTouchOne/holy/components/log"
    "github.com/HumanTouchOne/holy/engine"
    "github.com/HumanTouchOne/holy/router"
)

func main() {
    e := engine.New(func(eng *engine.Engine) {
        eng.SetAddr(":3003")
    })
    log.InitLogger([]string{"./logs/output.log"}, []string{"./logs/error.log"})
    router.LoadRouter(e)
    if err := e.Run(); err != nil {
        panic(err)
    }
}
```

## Routing

### Basic Routes

```go
e.GET("/path", handler)
e.POST("/path", handler)
e.PUT("/path", handler)
e.DELETE("/path", handler)
```

### RESTful

```go
type UserHandler struct{}

func (h *UserHandler) GET(ctx *engine.Context) {
    ctx.JSON(200, map[string]any{"users": []string{}})
}
func (h *UserHandler) POST(ctx *engine.Context) {
    ctx.AjaxSuccess("created", nil)
}

e.Rest("/user", &UserHandler{})
```

### Middleware

```go
// Before hook
e.Use(engine.PosAhead, engine.Prefix("/api", []string{}), authMiddleware)

// After hook
e.Use(engine.PosBehind, engine.Prefix("/api", []string{}), loggingHook)
```

## Context Helpers

```go
// Response
ctx.JSON(200, data)
ctx.AjaxSuccess("success", data)
ctx.AjaxError(400, "error", nil)
ctx.String(200, "text")

// Request
ctx.Query("key")
ctx.QueryInt("page", 1)
ctx.PostForm("field")
ctx.Cookie("token")
ctx.JSONArgs(&schema)
```

## Components


| Component | Description |
|-----------|-------------|
| log | Uber-Zap JSON logger |
| cache | Redis caching |
| jwt | JWT (HS256) |
| crypto | AES, Base64, SHA256, HMAC |
| orm | MySQL ORM |

## License

MIT
