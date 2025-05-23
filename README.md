# Go Logging Library

A quite simple, flexible logging library for Go apps.

## Overview

This library provides a straightforward way to implement logging in your Go projects. It supports multiple log levels, customizable output destinations, and detailed log messages that include timestamps, severity levels, and caller information.

## Features

*   **Multiple Log Levels**: Supports `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`, and `OFF` levels.
*   **Customizable Output**: Log output can be directed to any `io.Writer` (e.g., `os.Stdout`, a file).
*   **Configurable Log Level**: Set a minimum log level to control verbosity.
*   **Detailed Log Format**: Logs include timestamp, log level, caller function name, file name, and line number.
*   **Easy to Use**: Simple API for creating and using loggers.
*   **Fatal Logging**: `Fatal` level logs a message and then terminates the application.

## Installation

To use this library, you can import it into your Go project. (Assuming it's hosted on a VCS like GitHub, you'd typically use `go get`):

```sh
  go get github.com/melih-gulerb/logging
```

## Example Usage

```go
package main

import (
	"os"
	"github.com/melih-gulerb/logging"
)

func main() {
	logger := logging.New(logging.INFO, os.Stdout)
	logger.Info("This is an informational message.")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal("Failed to open log file: %v", err)
		return
	}
	defer file.Close()

	fileLogger := logging.New(logging.DEBUG, file)
	fileLogger.Debug("This is a debug message for the file.")
}
```

## Log Format

```
YYYY-MM-DD HH:MM:SS.mmm [LEVEL] [filename.go:lineNo funcName()] Your log message
2023-10-27 10:00:00.123 [INFO] [main.go:15 main()] This is an informational message.
```