# Simple Go Logger

Structured logger with per-request UUID propagation via `context.Context`.

Each application instance gets a random global UUID on startup. When a
`context.Context` carrying a `CtxKeyUUID` value is passed as the first
argument, that UUID is used in the log line instead — making it easy to
correlate all messages belonging to a single request or goroutine.

## Quick start

```go
package main

import (
	"context"

	"github.com/ra-company/logging"
)

func main() {
	logging.Logs.Starting("my-service")
	defer logging.Logs.Stopping()

	ctx := context.Background()
	logging.Logs.Info(ctx, "service initialised")

	// Attach a per-request UUID to the context
	requestCtx := context.WithValue(ctx, logging.CtxKeyUUID, "3d861cf8-ab1c-4d6d-b91e-ba17027a0045")
	logging.Logs.Debugf(requestCtx, "handling request for user %s", "alice")

	logging.Logs.Error("something went wrong (no context)")
}
```

Sample output (ShowTime=true, the default):

```
2025/06/17 18:17:42.016	INF	[f4d14d28-ae09-4aed-958a-c6dcb6da2a89]	my-service service is starting...
2025/06/17 18:17:42.016	INF	[f4d14d28-ae09-4aed-958a-c6dcb6da2a89]	service initialised
2025/06/17 18:17:42.017	DBG	[3d861cf8-ab1c-4d6d-b91e-ba17027a0045]	handling request for user alice
2025/06/17 18:17:42.018	ERR	[f4d14d28-ae09-4aed-958a-c6dcb6da2a89]	something went wrong (no context)
2025/06/17 18:17:42.018	INF	[f4d14d28-ae09-4aed-958a-c6dcb6da2a89]	my-service service is stopping...
```

## Calling convention

All log methods accept `...any`. The first argument may optionally be a
`context.Context`; remaining arguments follow one of two patterns:

| Variant | Signature | Behaviour |
|---------|-----------|-----------|
| Plain   | `Debug(msg, ...)` | arguments joined via `fmt.Sprint` |
| Formatted | `Debugf(format, args...)` | `fmt.Sprintf`-style substitution |

Both variants accept an optional leading context:

```go
Logs.Debug("value:", v)                        // plain, no context
Logs.Debug(ctx, "value:", v)                   // plain, with context

Logs.Debugf("attempt %d of %d", n, max)        // formatted, no context
Logs.Debugf(ctx, "user %s logged in", name)    // formatted, with context
```

Available levels: `Debug` / `Debugf`, `Info` / `Infof`, `Warn` / `Warnf`,
`Error` / `Errorf`, `Fatal` / `Fatalf`.

`Fatal`/`Fatalf` call `os.Exit(1)` after logging unless `DontStop` is set.

## Configuration

Set package-level variables before the first log call:

```go
logging.GraylogAddr = "graylog.example.com:12201" // empty = Graylog disabled
logging.Application = "my-service"
logging.Host        = "hostname"
logging.Facility    = "facility"

logging.Logs.LogLevel   = 0     // 0=debug, 1=warn, 2=error, 3=fatal
logging.Logs.ShowTime   = true  // prepend timestamp (default: true)
logging.Logs.ConsoleApp = false // true: only error/fatal reach stdout
logging.Logs.DontStop   = false // true: Fatal does not call os.Exit
```

## Custom logger

`CustomLogger` wraps any `Logger` interface implementation and falls back to
the global `Logs` when no inner logger is set. Use it to inject a logger into
structs without depending on the global singleton:

```go
type MyService struct {
	log logging.Logger
}

svc := &MyService{}
cl := &logging.CustomLogger{}
cl.SetLogger(&logging.Logs) // or any Logger implementation
svc.log = cl
```

## Graylog (GELF/UDP)

Set `GraylogAddr` to enable. Messages are sent over UDP using the GELF 1.1
protocol. The writer is initialised lazily on the first log call. Delivery is
best-effort; write errors are silently ignored so that Graylog unavailability
never affects the application.

## Staying up to date

```bash
go get -u github.com/ra-company/logging
```

## Supported Go versions

Go 1.22 and later.

## License

MIT — see [LICENSE](LICENSE).

## Dependencies

**Runtime**

| Library | Purpose | License |
|---------|---------|---------|
| [gopkg.in/Graylog2/go-gelf.v2](https://github.com/Graylog2/go-gelf) | GELF UDP writer | MIT |
| [github.com/google/uuid](https://github.com/google/uuid) | RFC 4122 UUID generation | BSD-3-Clause |

**Tests only**

| Library | Purpose | License |
|---------|---------|---------|
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | Test assertions | MIT |
