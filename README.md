# loggie 🧠⚡️ — Context-Aware Logger for Go

Context-aware, pluggable logger for Go web applications.  
No more passing logger through every function — just use `context.Context`.

---

## 🚀 What is loggie?

`loggie` helps you embed a structured logger inside `context.Context`, so you can log from any layer — service, repository, or handler — with a consistent `trace_id`, `user_id`, or any custom field.

It supports Zap, Logrus, OpenTelemetry, and works with any web framework.

---

## ✨ Features

✅ Structured logging via `context.Context`  
✅ Auto-generated `trace_id` per request  
✅ OTEL-compatible (extracts `trace_id` from spans)  
✅ Custom fields via `loggie.WithCustomField()`  
✅ Framework-agnostic logger injection  
✅ Zap / Logrus support  
✅ No-op fallback logger  
✅ Fx lifecycle and context cancellation compatible

---

## 📦 Installation

```bash
go get github.com/ditthkr/loggie
````

---

## 🧠 Injecting Logger in Middleware

Use `loggie.Injection(ctx, logger)` inside any middleware, from any web framework.

```go
ctx, traceId := loggie.Injection(req.Context(), logger)
```

### ✅ Fiber

```go
app.Use(func(c *fiber.Ctx) error {
	ctx, traceId := loggie.Injection(c.UserContext(), logger)
	c.SetUserContext(ctx)
	c.Set("X-Trace-Id", traceId)
	return c.Next()
})
```

### ✅ Gin

```go
r.Use(func(c *gin.Context) {
	ctx, traceId := loggie.Injection(c.Request.Context(), logger)
	c.Request = c.Request.WithContext(ctx)
	c.Writer.Header().Set("X-Trace-Id", traceId)
	c.Next()
})
```

### ✅ Echo

```go
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, traceId := loggie.Injection(c.Request().Context(), logger)
		req := c.Request().WithContext(ctx)
		c.SetRequest(req)
		c.Response().Header().Set("X-Trace-Id", traceId)
		return next(c)
	}
})
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

### ✅ Supported Adapters

| Logger | Package               |
| ------ | --------------------- |
| Zap    | `loggie.ZapLogger`    |
| Logrus | `loggie/logruslogger` |

---

## 📄 Usage Example

```go
log := loggie.FromContext(ctx)
log.Info("Order created", "user_id", 123)
```

📤 Output (Zap or Logrus):

```json
{
  "msg": "Order created",
  "trace_id": "...",
  "user_id": 123
}
```

---

## 🧰 Utilities

| Function                         | Description                        |
| -------------------------------- | ---------------------------------- |
| `Injection(ctx, logger)`      | Inject logger + trace\_id into ctx |
| `FromContext(ctx)`               | Get logger with fields from ctx    |
| `WithTraceId(ctx)`               | Add trace\_id manually             |
| `TraceId(ctx)`                   | Get trace\_id (auto or OTEL)       |
| `WithLogger(ctx, logger)`        | Attach logger                      |
| `WithCustomField(ctx, key, val)` | Add any key-value field            |
| `DefaultLogger()`                | Get fallback no-op logger          |

---

## 🌐 OpenTelemetry Support

If a context already contains an OTEL span, loggie will extract its `trace_id`:

```go
span := trace.SpanFromContext(ctx)
traceId := span.SpanContext().TraceId().String()
```

You don’t need to do anything extra — `TraceId(ctx)` handles it automatically.

---

## 📃 License

MIT © 2025 [@ditthkr](https://github.com/ditthkr)
