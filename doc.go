// Package logging provides a singleton logger with optional Graylog (GELF/UDP) output.
//
// # Calling convention
//
// All log methods accept a variadic ...any argument list. The first argument
// may optionally be a [context.Context]; when present, the UUID stored under
// [CtxKeyUUID] is used in the log line instead of the global [Logging.UUID].
//
// Plain (non-f) variants concatenate remaining arguments via [fmt.Sprint]:
//
//	Logs.Debug("connected to", addr)          // no context
//	Logs.Info(ctx, "request received")        // with context
//
// Formatted (f) variants treat the next argument as a [fmt.Sprintf] format
// string, followed by its arguments:
//
//	Logs.Debugf("attempt %d of %d", n, max)         // no context
//	Logs.Infof(ctx, "user %s logged in", username)  // with context
//
// Passing a non-string value where a format string is expected is a caller
// error; the call is silently discarded.
//
// # Singleton initialisation
//
// [Logs] is initialised in init() with ShowTime=true, LogLevel=0 and a random
// UUID. Set package-level variables before the first log call:
//
//	logging.GraylogAddr = "graylog.example.com:12201"
//	logging.Application = "my-service"
//	logging.Logs.LogLevel = 1  // suppress debug output
//	logging.Logs.Starting("my-service")
//	defer logging.Logs.Stopping()
package logging
