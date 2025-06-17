package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Context key type
type CtxKey string

var (
	Logs       Logging
	CtxKeyUUID CtxKey = "process-uuid" // Context key for process UUID
)

type Logging struct {
	UUID       string
	LogLevel   int    // Log level (0 - debug, 1 - warning, 2 - error, 3 - fatal, default 0)
	ConsoleApp bool   // Console application flag (do not print logs in console app)
	ShowTime   bool   // Show time in logs
	title      string // Process title
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

	if lev != "" {
		t := time.Now()
		if logger.ShowTime {
			if withContext {
				fmt.Printf("%s\t%v\t[%v]\t%v\n", logger.TimeToStr(t), lev, uuid, fmt.Sprint(args[1:]...))
			} else {
				fmt.Printf("%s\t%v\t[%v]\t%v\n", logger.TimeToStr(t), lev, uuid, fmt.Sprint(args...))
			}
		} else {
			if withContext {
				fmt.Printf("%v\t[%v]\t%v\n", lev, uuid, fmt.Sprint(args[1:]...))
			} else {
				fmt.Printf("%v\t[%v]\t%v\n", lev, uuid, fmt.Sprint(args...))
			}
		}

	}
}

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

	if lev != "" {
		t := time.Now()
		if withContext {
			if logger.ShowTime {
				if len(args) > 2 {
					fmt.Printf("%s\t%v\t[%v]\t%v\n", logger.TimeToStr(t), lev, uuid, fmt.Sprintf(args[1].(string), args[2:]...))
				} else {
					fmt.Printf("%s\t%v\t[%v]\t%v\n", logger.TimeToStr(t), lev, uuid, fmt.Sprint(args[1:]...))
				}
			} else {
				if len(args) > 2 {
					fmt.Printf("%v\t[%v]\t%v\n", lev, uuid, fmt.Sprintf(args[1].(string), args[2:]...))
				} else {
					fmt.Printf("%v\t[%v]\t%v\n", lev, uuid, fmt.Sprint(args[1:]...))
				}
			}
		} else {
			if logger.ShowTime {
				fmt.Printf("%s\t%v\t[%v]\t%v\n", logger.TimeToStr(t), lev, uuid, fmt.Sprintf(args[0].(string), args[1:]...))
			} else {
				fmt.Printf("%v\t[%v]\t%v\n", lev, uuid, fmt.Sprintf(args[0].(string), args[1:]...))
			}
		}
	}
}

func (logger *Logging) TimeToStr(t time.Time) string {
	str := t.Format("2006/01/02 15:04:05.999")
	if len(str) == 21 {
		str += "00"
	} else if len(str) == 22 {
		str += "0"
	}

	return str
}

func (logger *Logging) Info(args ...any) {
	logger.Print(4, args...)
}

func (logger *Logging) Infof(args ...any) {
	logger.Printf(4, args...)
}

func (logger *Logging) Debug(args ...any) {
	logger.Print(0, args...)
}

func (logger *Logging) Debugf(args ...any) {
	logger.Printf(0, args...)
}

func (logger *Logging) Warn(args ...any) {
	logger.Print(1, args...)
}

func (logger *Logging) Warnf(args ...any) {
	logger.Printf(1, args...)
}

func (logger *Logging) Error(args ...any) {
	logger.Print(2, args...)
}

func (logger *Logging) Errorf(args ...any) {
	logger.Printf(2, args...)
}

func (logger *Logging) Fatal(args ...any) {
	logger.Print(3, args...)
}

func (logger *Logging) Fatalf(args ...any) {
	logger.Printf(3, args...)
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
