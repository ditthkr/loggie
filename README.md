# loggie 🧠⚡️ — Context-Aware Logger for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/ditthkr/loggie.svg)](https://pkg.go.dev/github.com/ditthkr/loggie)
[![Go Report Card](https://goreportcard.com/badge/github.com/ditthkr/loggie)](https://goreportcard.com/report/github.com/ditthkr/loggie)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

> Simple, structured, and traceable logging via `context.Context` — for modern Go backend services.

---

## ✨ Features

- ✅ Context-aware logging (`logger.FromContext(ctx)`)
- 🧵 Auto generate `trace_id` per request
- 🏷 Attach dynamic custom fields (`user_id`, `order_id`, etc.)
- 🔌 Plug-and-play middleware for **Gin**, **Fiber**, and **Echo**
- 🔧 Compatible with `context.WithTimeout`, `WithCancel`
- 📦 Designed for use with Zap (Logrus / slog coming soon)

---

## 📦 Installation

```bash
go get github.com/ditthkr/loggie
````

---

## 🚀 Quick Start

### 🔗 Gin

```go
r := gin.Default()
r.Use(middleware.GinZapMiddleware(zapLogger))

r.GET("/ping", func(c *gin.Context) {
	ctx := loggie.WithCustomField(c.Request.Context(), "user_id", 123)
	log := loggie.FromContext(ctx)
	log.Info("hello gin")
	c.JSON(200, gin.H{"msg": "pong"})
})
```

---

### ⚡ Fiber

```go
app := fiber.New()
app.Use(middleware.FiberZapMiddleware(zapLogger))

app.Get("/ping", func(c fiber.Ctx) error {
	ctx := loggie.WithCustomField(c.Context(), "user_id", 123)
	log := loggie.FromContext(ctx)
	log.Info("hello fiber")
	return c.JSON(fiber.Map{"msg": "pong"})
})
```

---

### 🎯 Echo

```go
e := echo.New()
e.Use(middleware.EchoZapMiddleware(zapLogger))

e.GET("/ping", func(c echo.Context) error {
	ctx := loggie.WithCustomField(c.Request().Context(), "role", "admin")
	log := loggie.FromContext(ctx)
	log.Info("hello echo")
	return c.JSON(200, map[string]string{"msg": "pong"})
})
```

---

## 🔍 Log Output (structured JSON)

```json
{
  "level": "info",
  "msg": "hello gin",
  "trace_id": "123e4567-e89b-12d3-a456-426614174000",
  "user_id": 123
}
```

---

## 🧪 WithTimeout / Cancel

`loggie` works seamlessly with `context.WithTimeout` or `context.WithCancel`.

```go
ctx := loggie.WithCustomField(r.Context(), "user_id", 999)
ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()

log := loggie.FromContext(ctx)
log.Info("processing with timeout")
```

---

## 🧱 Custom Fields

Attach custom metadata across service layers without rewriting logger or trace logic.

```go
ctx = loggie.WithCustomField(ctx, "order_id", "ORD-789")
log := loggie.FromContext(ctx)
log.Info("step 1") // order_id will appear
```

---

## 📌 Roadmap

* [x] Gin Middleware
* [x] Fiber v3 Middleware
* [x] Echo Middleware
* [ ] Adapter for `logrus`
* [ ] Adapter for `slog` (Go 1.21+)
* [ ] Tracing integration (OpenTelemetry)
* [ ] Unit tests & benchmarks

---

## 📄 License

MIT License © 2025 [DITTHKR](https://github.com/ditthkr)