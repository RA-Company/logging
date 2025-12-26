package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/google/uuid"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

// Context key type
type CtxKey string

var (
	Logs        Logging
	CtxKeyUUID  CtxKey = "process-uuid"   // Context key for process UUID
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
	lev, uuid, withContext := logger.GetLevel(level, args[0])
	if logger.ConsoleApp {
		if level == 2 || level == 3 {
			if withContext {
				fmt.Print(fmt.Sprint(args[1:]...))
			} else {
				fmt.Print(fmt.Sprint(args...))
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
			fmt.Printf("%s\t%v\t[%v]\t%v\n", logger.TimeToStr(t), lev, uuid, text)
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
			fmt.Printf("%s\t%v\t[%v]\t%v\n", logger.TimeToStr(t), lev, uuid, text)
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

	if logger.graylog == nil {
		if writer, err := gelf.NewUDPWriter(GraylogAddr); err != nil {
			return
		} else {
			logger.graylog = writer
		}
	}

	var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

	message := gelf.Message{
		Short:    ansiRegexp.ReplaceAllString(text, ""),
		Version:  "1.1",
		Host:     Host,
		TimeUnix: float64(t.UnixNano()) / float64(time.Second),
		Facility: Facility,
		RawExtra: json.RawMessage(`{"_application":"` + Application + `", "_uuid":"` + uuid + `", "_type":"` + lev + `"}`),
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

	logger.graylog.WriteMessage(&message)
}

// TimeToStr converts time.Time to string in format "2006/01/02 15:04:05.999"
// It ensures that the string is always 23 characters long by appending "00" or "0" as needed.
//
// Parameters:
//   - t - time.Time object to convert
//
// Returns:
//   - string: representation of the time in the specified format
func (logger *Logging) TimeToStr(t time.Time) string {
	str := t.Format("2006/01/02 15:04:05.999")

	if len(str) == 19 {
		return str + ".000"
	} else if len(str) == 21 {
		return str + "00"
	} else if len(str) == 22 {
		return str + "0"
	}

	return str
}

// Info logs an informational message.
//
// Parameters:
//   - args - arguments to print
//     # args[0] - context (optional) or argument to print
//     # args[1:] - arguments to print
func (logger *Logging) Info(args ...any) {
	logger.Printf(4, args...)
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
	logger.Printf(0, args...)
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
	logger.Printf(1, args...)
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
	logger.Printf(2, args...)
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
	logger.Printf(3, args...)
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
