package logging

// Logger Interface defines the methods for logging at different levels.
type Logger interface {
	// Debug logs a message at Debug level.
	// The first argument is the context for the log entry,
	Debug(args ...any)

	// Debugf logs a formatted message at Debug level.
	// The first argument is the context for the log entry,
	Debugf(args ...any)

	// Info logs a message at Info level.
	// The first argument is the context for the log entry,
	Info(args ...any)

	// Infof logs a formatted message at Info level.
	// The first argument is the context for the log entry,
	Infof(args ...any)

	// Warn logs a message at Warning level.
	// The first argument is the context for the log entry,
	Warn(args ...any)

	// Warnf logs a formatted message at Warning level.
	// The first argument is the context for the log entry,
	Warnf(args ...any)

	// Error logs a message at Error level.
	// The first argument is the context for the log entry,
	Error(args ...any)

	// Errorf logs a formatted message at Error level.
	// The first argument is the context for the log entry,
	Errorf(args ...any)

	// Fatal logs a message at Fatal level
	// and process will exit with status set to 1.
	// The first argument is the context for the log entry,
	Fatal(args ...any)

	// Fatalf logs a formatted message at Fatal level
	// and process will exit with status set to 1.
	// The first argument is the context for the log entry,
	Fatalf(args ...any)
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
//   - args: The debug message to log, which can be any number of arguments.
//     # args[0] - context (optional) or argument to print
//     # args[1:] - additional arguments to print
func (dst *CustomLogger) Debug(args ...any) {
	if dst.logger != nil {
		dst.logger.Debug(args...)
		return
	}

	Logs.Debug(args...)
}

// Debugf logs a formatted debug message using the provided logger or the default logging mechanism.
// It is useful for debugging purposes and can be used to log detailed information about the API calls
// and their responses with formatted strings.
//
// Parameters:
//   - args: The formatted debug message to log, which can be any number of arguments.
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - additional arguments to format string
func (dst *CustomLogger) Debugf(args ...any) {
	if dst.logger != nil {
		dst.logger.Debugf(args...)
		return
	}

	Logs.Debugf(args...)
}

// Info logs an informational message using the provided logger or the default logging mechanism.
// It is useful for logging general information about the API calls and their responses.
//
// Parameters:
//   - args: The informational message to log, which can be any number of arguments.
//     # args[0] - context (optional) or argument to print
//     # args[1:] - additional arguments to print
func (dst *CustomLogger) Info(args ...any) {
	if dst.logger != nil {
		dst.logger.Info(args...)
		return
	}

	Logs.Info(args...)
}

// Infof logs a formatted informational message using the provided logger or the default logging mechanism.
// It is useful for logging general information about the API calls and their responses with formatted strings.
//
// Parameters:
//   - args: The formatted informational message to log, which can be any number of arguments.
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - additional arguments to format string
func (dst *CustomLogger) Infof(args ...any) {
	if dst.logger != nil {
		dst.logger.Infof(args...)
		return
	}

	Logs.Infof(args...)
}

// Warn logs a warning message using the provided logger or the default logging mechanism.
// It is useful for logging potential issues or unexpected behavior in the API calls.
//
// Parameters:
//   - args: The warning message to log, which can be any number of arguments.
//     # args[0] - context (optional) or argument to print
//     # args[1:] - additional arguments to print
func (dst *CustomLogger) Warn(args ...any) {
	if dst.logger != nil {
		dst.logger.Warn(args...)
		return
	}

	Logs.Warn(args...)
}

// Warnf logs a formatted warning message using the provided logger or the default logging mechanism.
// It is useful for logging potential issues or unexpected behavior in the API calls with formatted strings.
//
// Parameters:
//   - args: The formatted warning message to log, which can be any number of arguments.
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - additional arguments to format string
func (dst *CustomLogger) Warnf(args ...any) {
	if dst.logger != nil {
		dst.logger.Warnf(args...)
		return
	}

	Logs.Warnf(args...)
}

// Error logs an error message using the provided logger or the default logging mechanism.
// It is useful for logging errors encountered during API calls or other operations.
//
// Parameters:
//   - args: The error message to log, which can be any number of arguments.
//     # args[0] - context (optional) or argument to print
//     # args[1:] - additional arguments to print
func (dst *CustomLogger) Error(args ...any) {
	if dst.logger != nil {
		dst.logger.Error(args...)
		return
	}

	Logs.Error(args...)
}

// Errorf logs a formatted error message using the provided logger or the default logging mechanism.
// It is useful for logging errors encountered during API calls or other operations with formatted strings.
//
// Parameters:
//   - args: The formatted error message to log, which can be any number of arguments.
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - additional arguments to format string
func (dst *CustomLogger) Errorf(args ...any) {
	if dst.logger != nil {
		dst.logger.Errorf(args...)
		return
	}

	Logs.Errorf(args...)
}

// Fatal logs a fatal error message using the provided logger or the default logging mechanism.
// It is useful for logging critical errors that require immediate attention and may cause the application to exit.
//
// Parameters:
//   - args: The fatal error message to log, which can be any number of arguments.
//     # args[0] - context (optional) or argument to print
//     # args[1:] - additional arguments to print
func (dst *CustomLogger) Fatal(args ...any) {
	if dst.logger != nil {
		dst.logger.Fatal(args...)
		return
	}

	Logs.Fatal(args...)
}

// Fatalf logs a formatted fatal error message using the provided logger or the default logging mechanism.
// It is useful for logging critical errors that require immediate attention and may cause the application to exit with formatted strings.
//
// Parameters:
//   - args: The formatted fatal error message to log, which can be any number of arguments.
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - additional arguments to format string
func (dst *CustomLogger) Fatalf(args ...any) {
	if dst.logger != nil {
		dst.logger.Fatalf(args...)
		return
	}

	Logs.Fatalf(args...)
}
