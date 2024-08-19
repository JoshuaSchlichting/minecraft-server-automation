package logger

import (
	"log"
	"os"
)

// Define log levels
const (
	TRACE = iota
	DEBUG
	INFO
	WARNING
	ERROR
)

var (
	logLevel = INFO // Default log level
	logger   = log.New(os.Stdout, "", log.LstdFlags)
)

// SetLogLevel sets the global log level
func SetLogLevel(level int) {
	logLevel = level
}

// Log functions for each level
func Trace(v ...interface{}) {
	if logLevel <= TRACE {
		logger.SetPrefix("TRACE: ")
		logger.Println(v...)
	}
}

func Debug(v ...interface{}) {
	if logLevel <= DEBUG {
		logger.SetPrefix("DEBUG: ")
		logger.Println(v...)
	}
}

func Info(v ...interface{}) {
	if logLevel <= INFO {
		logger.SetPrefix("INFO: ")
		logger.Println(v...)
	}
}

func Warning(v ...interface{}) {
	if logLevel <= WARNING {
		logger.SetPrefix("WARNING: ")
		logger.Println(v...)
	}
}

func Error(v ...interface{}) {
	if logLevel <= ERROR {
		logger.SetPrefix("ERROR: ")
		logger.Println(v...)
	}
}

func Fatal(v ...interface{}) {
	if logLevel <= ERROR {
		logger.SetPrefix("FATAL: ")
		logger.Fatalln(v...)
	}
}
