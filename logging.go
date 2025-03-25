package logging

import (
	"context"
	"fmt"
	"time"
)

// Context key type
type CtxKey string

var (
	Logs       *Logging
	CtxKeyUUID CtxKey = "process-uuid" // Context key for process UUID
)

type Logging struct {
	uuid         string
	LogLevel     int    // Log level (0 - debug, 1 - warning, 2 - error, 3 - fatal, default 0)
	ConsoleApp   bool   // Console application flag (do not print logs in console app)
	DontShowTime bool   // Do not show time in logs
	title        string // Process title
}

// Get level of logging by level and context if it's present
func (logger *Logging) GetLevel(level int, ctx interface{}) (string, string, bool) {
	var uuid string
	withContext := false

	switch ctx.(type) {
	case context.Context:
		if ctx.(context.Context).Value(CtxKeyUUID) != nil {
			uuid = ctx.(context.Context).Value(CtxKeyUUID).(string)
		} else {
			uuid = logger.uuid
		}
		withContext = true
	default:
		uuid = logger.uuid
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

func (logger *Logging) Print(level int, args ...interface{}) {
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
		if logger.DontShowTime {
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

func (logger *Logging) Printf(level int, args ...interface{}) {
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
			if logger.DontShowTime {
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
			if logger.DontShowTime {
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

func (logger *Logging) Info(args ...interface{}) {
	logger.Print(4, args...)
}

func (logger *Logging) Infof(args ...interface{}) {
	logger.Printf(4, args...)
}

func (logger *Logging) Debug(args ...interface{}) {
	logger.Print(0, args...)
}

func (logger *Logging) Debugf(args ...interface{}) {
	logger.Printf(0, args...)
}

func (logger *Logging) Warn(args ...interface{}) {
	logger.Print(1, args...)
}

func (logger *Logging) Warnf(args ...interface{}) {
	logger.Printf(1, args...)
}

func (logger *Logging) Error(args ...interface{}) {
	logger.Print(2, args...)
}

func (logger *Logging) Errorf(args ...interface{}) {
	logger.Printf(2, args...)
}

func (logger *Logging) Fatal(args ...interface{}) {
	logger.Print(3, args...)
}

func (logger *Logging) Fatalf(args ...interface{}) {
	logger.Printf(3, args...)
}

// Starting logs
//
// Parameters:
//   - uuid - default process UUID
//   - title - process title
func (logger *Logging) Starting(uuid, title string) {
	logger.title = title
	logger.Infof("%s service is starting...", title)
}

// Stopping logs
func (logger *Logging) Stopping() {
	logger.Infof("%s service is stopping...", logger.title)
}
