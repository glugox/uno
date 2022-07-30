package uno

// At this point, this is only a facade around Go std lib "log"

import (
	"fmt"
	stdLog "log"
	"runtime"
	"sync"
)

const (
	// Useuall more than we need
	LogLevelVerbose = 1

	// This level helps identifying
	//  any issues
	LogLevelDebug = 2

	// Only the most usefull information
	// on normal usage
	LogLevelInfo = 3

	// Shows only messages that suggests something
	// is wrong, but program can still continue
	LogLevelWarn = 4

	// Shows only the error why the application
	// stopped working
	LogLevelError = 5
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37;02m"
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

var (

	// Default log level
	LogLevel = LogLevelInfo
)

var lock = &sync.Mutex{}

// Logger interface that objects accepts
type Logger interface {
	Verbose(format string, s ...any)
	Debug(format string, s ...any)
	Info(format string, s ...any)
	Success(format string, s ...any)
	Warn(format string, s ...any)
	Error(format string, s ...any) (err error)
}

// Log which reference should be passed through obejcts
// that needs logging
type Log struct {
	Level uint8
}

// To create multiple types of logs on multiple
// places, we use log.Factory.Create
type Factory struct{}

// NewFactory returns new Factory instance
func NewFactory() *Factory {
	return &Factory{}
}

// Keep track of our global log factory instance
var logFactory *Factory

// Returns singleton instance of our
// default log factory
func DefaultLogFactory() *Factory {
	if logFactory == nil {
		lock.Lock()
		defer lock.Unlock()
		if logFactory == nil {
			logFactory = NewFactory()
		}
	}
	return logFactory
}

func (f *Factory) NewLogger() *Log {
	return &Log{
		Level: LogLevelVerbose,
	}
}

// Verbose
func (l *Log) Verbose(format string, s ...any) {
	if l.Level > LogLevelVerbose {
		return
	}
	stdLog.Printf(Gray+format+Reset, s...) // TODO: Correct implementations
}

// Debug
func (l *Log) Debug(format string, s ...any) {
	if l.Level > LogLevelDebug {
		return
	}
	stdLog.Printf(Gray+format+Reset, s...)
}

// Info
func (l *Log) Info(format string, s ...any) {
	if l.Level > LogLevelInfo {
		return
	}
	stdLog.Printf(format, s...)
}

// Success prints the message with green color
func (l *Log) Success(format string, s ...any) {
	if l.Level > LogLevelInfo {
		return
	}
	stdLog.Printf(Green+format+Reset, s...)
}

// Warn
func (l *Log) Warn(format string, s ...any) {
	if l.Level > LogLevelWarn {
		return
	}
	stdLog.Printf(format, s...)
}

// Error
func (l *Log) Error(format string, s ...any) (err error) {
	if l.Level > LogLevelError {
		return
	}
	err = fmt.Errorf(Red+format+Reset, s...)
	stdLog.Print(err.Error())
	return
}
