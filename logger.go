package logging

import (
	"context"
	"os"
)

// Logger Interface defines the methods for logging at different levels.
type Logger interface {
	// Debug logs a message at Debug level.
	// The first argument is the context for the log entry,
	Debug(args ...any)

	// Info logs a message at Info level.
	// The first argument is the context for the log entry,
	Info(args ...any)

	// Warn logs a message at Warning level.
	// The first argument is the context for the log entry,
	Warn(args ...any)

	// Error logs a message at Error level.
	// The first argument is the context for the log entry,
	Error(args ...any)

	// Fatal logs a message at Fatal level
	// and process will exit with status set to 1.
	// The first argument is the context for the log entry,
	Fatal(args ...any)
}

// CustomLogger is a wrapper around a Logger interface that allows
// for custom logging implementations. It provides methods to log messages
type CustomLogger struct {
	logger Logger
}

// SetLogger allows setting a custom logger.
// This is useful when you want to use a different logging mechanism
// instead of the default one provided by the package.
//
// Parameters:
//   - logger: An instance of a type that implements the Logger interface.
func (dst *CustomLogger) SetLogger(logger Logger) {
	dst.logger = logger
}

// Debug logs a debug message using the provided logger or the default logging mechanism.
// It is useful for debugging purposes and can be used to log detailed information about the API calls
// and their responses.
//
// Parameters:
//   - ctx: The context for the logging operation, allowing for cancellation and timeouts.
//   - text: The debug message to log.
func (dst *CustomLogger) Debug(ctx context.Context, text string) {
	if dst.logger != nil {
		dst.logger.Debug(ctx, text)
	} else {
		Logs.Debugf(ctx, text)
	}
}

// Info logs an informational message using the provided logger or the default logging mechanism.
// It is useful for logging general information about the API calls and their responses.
//
// Parameters:
//   - ctx: The context for the logging operation, allowing for cancellation and timeouts.
//   - text: The informational message to log.
func (dst *CustomLogger) Info(ctx context.Context, text string) {
	if dst.logger != nil {
		dst.logger.Info(ctx, text)
	} else {
		Logs.Infof(ctx, text)
	}
}

// Warn logs a warning message using the provided logger or the default logging mechanism.
// It is useful for logging potential issues or unexpected behavior in the API calls.
//
// Parameters:
//   - ctx: The context for the logging operation, allowing for cancellation and timeouts.
//   - text: The warning message to log.
func (dst *CustomLogger) Warn(ctx context.Context, text string) {
	if dst.logger != nil {
		dst.logger.Warn(ctx, text)
	} else {
		Logs.Warnf(ctx, text)
	}
}

// Error logs an error message using the provided logger or the default logging mechanism.
// It is useful for logging errors encountered during API calls or other operations.
//
// Parameters:
//   - ctx: The context for the logging operation, allowing for cancellation and timeouts.
//   - text: The error message to log.
func (dst *CustomLogger) Error(ctx context.Context, text string) {
	if dst.logger != nil {
		dst.logger.Error(ctx, text)
	} else {
		Logs.Errorf(ctx, text)
	}
}

// Fatal logs a fatal error message using the provided logger or the default logging mechanism.
// It is useful for logging critical errors that require immediate attention and may cause the application to exit.
//
// Parameters:
//   - ctx: The context for the logging operation, allowing for cancellation and timeouts.
//   - text: The fatal error message to log.
func (dst *CustomLogger) Fatal(ctx context.Context, text string) {
	if dst.logger != nil {
		dst.logger.Fatal(ctx, text)
		os.Exit(1) // Exit with status code 1
	} else {
		Logs.Fatalf(ctx, text)
	}
}
