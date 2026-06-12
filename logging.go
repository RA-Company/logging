package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/google/uuid"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

// Context key type
type CtxKey string

const (
	CtxKeyUUID CtxKey = "process-uuid" // Context key for process UUID
	fixedTime  string = "2006/01/02 15:04:05.000"
)

var (
	Logs        Logging
	GraylogAddr string = ""               // Graylog address (when empty, Graylog is disabled)
	Application string = "go-application" // Application name
	Host        string = "localhost"      // Host name
	Facility    string = "facility"       // Facility name
)

type Logging struct {
	UUID       string
	LogLevel   int    // Log level (0 - debug, 1 - warning, 2 - error, 3 - fatal, default 0)
	ConsoleApp bool   // Console application flag (do not print logs in console app)
	ShowTime   bool   // Show time in logs
	DontStop   bool   // Do not stop service on fatal error
	title      string // Process title
	graylog    *gelf.UDPWriter
	mu         sync.Mutex
}

// Get level of logging by level and context if it's present
//
// Parameters:
//   - level - log level
//   - ctx - context
func (logger *Logging) GetLevel(level int, ctx any) (string, string, bool) {
	var uuid string
	withContext := false

	switch ctx.(type) {
	case context.Context:
		if ctx.(context.Context).Value(CtxKeyUUID) != nil {
			uuid = ctx.(context.Context).Value(CtxKeyUUID).(string)
		} else {
			uuid = logger.UUID
		}
		withContext = true
	default:
		uuid = logger.UUID
	}

	levels := []string{"DBG", "WRN", "ERR", "FTL", "INF"}

	if level < 0 || level > 4 {
		level = 4
	}

	if level < logger.LogLevel {
		return "", uuid, withContext
	}

	return levels[level], uuid, withContext
}

// Print logs to console
//
// Parameters:
//   - level - log level (0 - debug, 1 - warning, 2 - error, 3 - fatal, 4 - info)
//   - args - arguments to print
func (logger *Logging) Print(level int, args ...any) {
	if len(args) == 0 {
		return
	}
	lev, uuid, withContext := logger.GetLevel(level, args[0])
	if logger.ConsoleApp {
		if level == 2 || level == 3 {
			if withContext {
				fmt.Printf("%s\n", fmt.Sprint(args[1:]...))
			} else {
				fmt.Printf("%s\n", fmt.Sprint(args...))
			}
		}
		return // do not print logs in console app
	}

	t := time.Now()

	text := ""
	if lev != "" {
		if withContext {
			text = fmt.Sprint(args[1:]...)
		} else {
			text = fmt.Sprint(args...)
		}
		if logger.ShowTime {
			fmt.Printf("%s\t%v\t[%v]\t%v\n", t.Format(fixedTime), lev, uuid, text)
		} else {
			fmt.Printf("%v\t[%v]\t%v\n", lev, uuid, text)
		}
	}

	logger.sendGelfMessage(text, t, lev, uuid)
}

// Printf logs formatted output to console
//
// Parameters:
//   - level - log level (0 - debug, 1 - warning, 2 - error, 3 - fatal, 4 - info)
//   - args - arguments to print
//     # args[0] - format string
//     # args[1:] - arguments to format string
func (logger *Logging) Printf(level int, args ...any) {
	if len(args) < 2 {
		return
	}
	lev, uuid, withContext := logger.GetLevel(level, args[0])
	if logger.ConsoleApp {
		if level == 2 || level == 3 {
			if len(args) > 2 {
				fmt.Printf("%v\n", fmt.Sprintf(args[1].(string), args[2:]...))
			} else {
				fmt.Printf("%v\n", fmt.Sprint(args[1:]...))
			}
		}
		return // do not print logs in console app
	}

	t := time.Now()
	text := ""
	if lev != "" {
		if withContext {
			if len(args) > 2 {
				text = fmt.Sprintf(args[1].(string), args[2:]...)
			} else {
				text = fmt.Sprint(args[1:]...)
			}
		} else {
			text = fmt.Sprintf(args[0].(string), args[1:]...)
		}
		if logger.ShowTime {
			fmt.Printf("%s\t%v\t[%v]\t%v\n", t.Format(fixedTime), lev, uuid, text)
		} else {
			fmt.Printf("%v\t[%v]\t%v\n", lev, uuid, text)
		}
	}

	logger.sendGelfMessage(text, t, lev, uuid)
}

func (logger *Logging) sendGelfMessage(text string, t time.Time, lev, uuid string) {
	if GraylogAddr == "" || text == "" {
		return
	}

	logger.mu.Lock()
	if logger.graylog == nil {
		writer, err := gelf.NewUDPWriter(GraylogAddr)
		if err != nil {
			logger.mu.Unlock()
			return
		}
		logger.graylog = writer
	}
	graylogWriter := logger.graylog
	logger.mu.Unlock()

	extra, _ := json.Marshal(map[string]string{
		"_application": Application,
		"_uuid":        uuid,
		"_type":        lev,
	})

	message := gelf.Message{
		Short:    ansiRegexp.ReplaceAllString(text, ""),
		Version:  "1.1",
		Host:     Host,
		TimeUnix: float64(t.UnixNano()) / float64(time.Second),
		Facility: Facility,
		RawExtra: extra,
	}

	switch lev {
	case "DBG":
		message.Level = gelf.LOG_DEBUG
	case "WRN":
		message.Level = gelf.LOG_WARNING
	case "ERR":
		message.Level = gelf.LOG_ERR
	case "FTL":
		message.Level = gelf.LOG_CRIT
	case "INF":
		message.Level = gelf.LOG_INFO
	}

	graylogWriter.WriteMessage(&message)
}

// Info logs an informational message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or argument to print
//     # args[1:] - arguments to print
func (logger *Logging) Info(args ...any) {
	logger.Print(4, args...)
}

// Infof logs a formatted informational message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context) or argument to print
//     # args[2:] - arguments to format string
func (logger *Logging) Infof(args ...any) {
	logger.Printf(4, args...)
}

// Debug logs a debug message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or argument to print
//     # args[1:] - arguments to print
func (logger *Logging) Debug(args ...any) {
	logger.Print(0, args...)
}

// Debugf logs a formatted debug message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - arguments to format string
func (logger *Logging) Debugf(args ...any) {
	logger.Printf(0, args...)
}

// Warn logs a warning message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or argument to print
//     # args[1:] - arguments to print
func (logger *Logging) Warn(args ...any) {
	logger.Print(1, args...)
}

// Warnf logs a formatted warning message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - arguments to format string
func (logger *Logging) Warnf(args ...any) {
	logger.Printf(1, args...)
}

// Error logs an error message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or argument to print
//     # args[1:] - arguments to print
func (logger *Logging) Error(args ...any) {
	logger.Print(2, args...)
}

// Errorf logs a formatted error message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - arguments to format string
func (logger *Logging) Errorf(args ...any) {
	logger.Printf(2, args...)
}

// Fatal logs a fatal error message and exits the program.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or argument to print
//     # args[1:] - arguments to print
func (logger *Logging) Fatal(args ...any) {
	logger.Print(3, args...)
	if !logger.DontStop {
		os.Exit(1) // Exit with status code 1
	}
}

// Fatalf logs a formatted fatal error message and exits the program.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or format string
//     # args[1] - format string (if args[0] is context)
//     # args[2:] - arguments to format string
func (logger *Logging) Fatalf(args ...any) {
	logger.Printf(3, args...)
	if !logger.DontStop {
		os.Exit(1) // Exit with status code 1
	}
}

// Starting service
//
// Parameters:
//   - title - process title
func (logger *Logging) Starting(title string) {
	logger.title = title
	logger.Infof("%s service is starting...", title)
}

// Stopping service
func (logger *Logging) Stopping() {
	logger.Infof("%s service is stopping...", logger.title)
}

// Initialize default parameters
func init() {
	Logs.ShowTime = true
	Logs.ConsoleApp = false
	Logs.LogLevel = 0
	Logs.UUID = uuid.New().String()
}
