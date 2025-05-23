package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Logger is the main logging structure.
type Logger struct {
	minLevel LogLevel
	output   io.Writer
}

// New creates a new Logger instance.
// It defaults to the INFO level and writes to os.Stdout.
func New(minLevel LogLevel, out io.Writer) *Logger {
	if out == nil {
		out = os.Stdout
	}
	return &Logger{
		minLevel: minLevel,
		output:   out,
	}
}

// Default creates a new Logger with INFO level and stdout output.
func Default() *Logger {
	return New(INFO, os.Stdout)
}

// SetLevel sets the minimum log level for the logger.
func (l *Logger) SetLevel(level LogLevel) {
	l.minLevel = level
}

// SetOutput sets the output destination for the logger.
func (l *Logger) SetOutput(out io.Writer) {
	l.output = out
}

// log is the internal logging function.
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.minLevel {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	levelStr := level.String()
	message := fmt.Sprintf(format, args...)

	// Get caller info
	pc, file, line, ok := runtime.Caller(2) // 2 skips this log function and the public Debug/Info/... function
	var callerDetails string
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		// Get just the function name, not the full path
		shortFuncName := funcName[strings.LastIndex(funcName, ".")+1:]
		callerDetails = fmt.Sprintf("%s:%d %s()", filepath.Base(file), line, shortFuncName)
	} else {
		callerDetails = "???:0"
	}

	logEntry := fmt.Sprintf("%s [%s] [%s] %s\n", timestamp, levelStr, callerDetails, message)

	// Write to output
	// Using Fprintf to handle potential errors during writing, though we ignore them here for simplicity.
	// In a production library, you might want to handle write errors.
	_, _ = fmt.Fprint(l.output, logEntry)

	if level == FATAL {
		os.Exit(1)
	}
}

// Debug logs a message at DEBUG level.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info logs a message at the INFO level.
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn logs a message at WARN level.
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error logs a message at ERROR level.
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Fatal logs a message at FATAL level and then calls os.Exit(1).
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
}
