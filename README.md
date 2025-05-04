# loggie 🧠⚡️ — Context-Aware Logger for Go

Context-aware, pluggable logger for Go web applications.  
No more passing logger through every function — just use `context.Context`.

---

## 🚀 What is loggie?

`loggie` helps you embed a structured logger inside `context.Context`, so you can log from any layer — service, repository, or handler — with a consistent `trace_id`, `user_id`, or any custom field.

It supports Zap, Logrus (new!), and is ready for OpenTelemetry (OTEL).

---

## ✨ Features

✅ Structured logging via `context.Context`  
✅ Auto-generated `trace_id` per request  
✅ OTEL-compatible (detects `trace_id` from OpenTelemetry spans)  
✅ Custom fields (e.g. `user_id`, `order_id`)  
✅ Middleware for Fiber / Gin / Echo  
✅ Pluggable backends: Zap, Logrus, Slog (soon)  
✅ Fallback logger included (safe anywhere)  
✅ Works with Fx lifecycle & `context.WithTimeout`

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
(✅ ready)         (✅ ready)         (🕓 planned)
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

### Fiber + Zap

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

### Fiber + Logrus

```go
import (
    "github.com/sirupsen/logrus"
    "github.com/ditthkr/loggie/logruslogger"
    "github.com/ditthkr/loggie/middleware/fiberlog"
)


func main() {
    logger := logrus.New()
	
    adapter := &logruslogger.LogrusLogger{L: logrus.NewEntry(logger)}
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

---

## 🌐 OTEL Support (OpenTelemetry)

If you're using OpenTelemetry, `loggie` will **automatically extract `trace_id` from span context**
via `go.opentelemetry.io/otel/trace`.

No config required — just pass the OTEL-injected `context.Context`.

```go
ctx := r.Context() // contains OTEL span
log := loggie.FromContext(ctx)

log.Info("Received payment webhook")
// trace_id will match what's in OTEL system (Jaeger, Tempo, etc)
```

---

## ✍️ Custom Fields

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

| Function                         | Purpose                           |
| -------------------------------- | --------------------------------- |
| `FromContext(ctx)`               | Retrieves logger from context     |
| `WithLogger(ctx, logger)`        | Injects a logger                  |
| `WithTraceID(ctx)`               | Adds `trace_id` to context        |
| `TraceID(ctx)`                   | Retrieves `trace_id` (OTEL-aware) |
| `WithCustomField(ctx, key, val)` | Adds structured field             |
| `DefaultLogger()`                | No-op fallback logger             |

---

## 📁 Middleware Support

| Framework | Import Path                                     | Function                |
| --------- | ----------------------------------------------- | ----------------------- |
| Fiber     | `github.com/ditthkr/loggie/middleware/fiberlog` | `fiberlog.Middleware()` |
| Gin       | `github.com/ditthkr/loggie/middleware/ginlog`   | `ginlog.Middleware()`   |
| Echo      | `github.com/ditthkr/loggie/middleware/echolog`  | `echolog.Middleware()`  |

All middlewares are generic and support any `loggie.Logger`.

---

## 🔌 Logger Adapters

| Logger | Status      | Package                        |
| ------ | ----------- | ------------------------------ |
| Zap    | ✅ Supported | `loggie.ZapLogger`             |
| Logrus | ✅ Supported | `loggie/logruslogger`          |
| Slog   | 🕓 Planned  | `loggie/slogger` (coming soon) |

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
