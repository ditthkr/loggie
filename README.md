# loggie 🧠⚡️ — Context-Aware Logger for Go

Context-aware, pluggable logger for Go web applications.  
No more passing logger through every function — just use `context.Context`.

---

## 🚀 What is loggie?

`loggie` helps you embed a structured logger inside `context.Context`, so you can log from any layer — service, repository, or handler — with a consistent `trace_id`, `user_id`, or any custom field.

It supports Zap (now), and is extensible to Logrus, Slog, and more.

---

## ✨ Features

✅ Structured logging with `context.Context`  
✅ Auto-generated `trace_id` per request  
✅ Custom fields via `loggie.WithCustomField()`  
✅ Middleware for Fiber / Gin / Echo  
✅ Pluggable backends (Zap, Logrus, Slog, etc.)  
✅ Fallback logger included (safe in any context)  
✅ Ready for Fx lifecycle and timeout-aware `context.WithTimeout`

---

## 📦 Installation

```bash
go get github.com/ditthkr/loggie
````

---

## 🧱 Architecture

```txt
               ┌────────────────────┐
               │ context.Context    │
               └────────┬───────────┘
                        │
                ┌───────▼────────┐
                │ loggie.Logger  │◄── (interface)
                └───────┬────────┘
                        │
        ┌───────────────┼────────────────┐
        ▼               ▼                ▼
  ZapLogger        LogrusLogger       SlogLogger
(implemented)     (planned)          (planned)

```

---

## 🔌 Logger Interface

```go
type Logger interface {
    Info(msg string, fields ...any)
    Error(msg string, fields ...any)
    With(fields ...any) Logger
}
```

> You can plug any logger backend by implementing this interface.

---

## ⚙️ Middleware Usage

### Fiber + Zap (fully working)

```go
import (
    "github.com/ditthkr/loggie"
    "github.com/ditthkr/loggie/middleware/fiberlog"
    "go.uber.org/zap"
    "github.com/gofiber/fiber/v3"
)

func main() {
    rawLogger, _ := zap.NewProduction()
    defer rawLogger.Sync()

    adapter := &loggie.ZapLogger{L: rawLogger}
    app := fiber.New()
    app.Use(fiberlog.Middleware(adapter))

    app.Get("/ping", func(c *fiber.Ctx) error {
        log := loggie.FromContext(c.UserContext())
        log.Info("Ping received", "path", c.Path())
        return c.SendString("pong")
    })

    app.Listen(":8080")
}
```

🧪 Sample Log Output:

```json
{
  "level": "info",
  "msg": "Ping received",
  "trace_id": "b4a3f5a0...",
  "path": "/ping"
}
```

---

## ✍️ Custom Fields (e.g. user\_id)

```go
ctx = loggie.WithCustomField(ctx, "user_id", 42)

log := loggie.FromContext(ctx)
log.Info("Order created")
```

📤 Output:

```json
{
  "msg": "Order created",
  "trace_id": "abc-xyz",
  "user_id": 42
}
```

---

## 🧰 Utilities

| Function                           | Purpose                       |
|------------------------------------| ----------------------------- |
| `FromContext(ctx)`                 | Retrieves logger from context |
| `WithLogger(ctx, logger)`          | Injects a logger              |
| `WithTraceId(ctx)`                 | Adds `trace_id` to context    |
| `TraceId(ctx)`                     | Retrieves trace\_id           |
| `WithCustomField(ctx, key, value)` | Adds any structured field     |
| `DefaultLogger()`                  | Returns no-op fallback logger |

---

## 📁 Available Middleware

| Framework | Import Path                                     | Function                |
| --------- | ----------------------------------------------- | ----------------------- |
| Fiber     | `github.com/ditthkr/loggie/middleware/fiberlog` | `fiberlog.Middleware()` |
| Gin       | `github.com/ditthkr/loggie/middleware/ginlog`   | `ginlog.Middleware()`   |
| Echo      | `github.com/ditthkr/loggie/middleware/echolog`  | `echolog.Middleware()`  |

> All middlewares are generic and accept any `loggie.Logger`.

---

## 🔌 Current and Planned Logger Adapters

| Logger | Package / Status        |
| ------ | ----------------------- |
| Zap    | ✅ `loggie.ZapLogger`    |
| Logrus | 🕓 In progress          |
| Slog   | 🕓 Planned for Go 1.21+ |

---

## 🧪 Testing & Fallbacks

Even without injecting a logger, `loggie` will still work with a **safe no-op fallback**:

```go
log := loggie.FromContext(context.Background())
log.Info("This is safe even without a logger")
```

---

## 📃 License

MIT © 2025 [@ditthkr](https://github.com/ditthkr)