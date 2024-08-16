package logger

import (
	"fmt"
)

// ConsoleLogger is a concrete logger for logging to the console
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Printf("%s\n", message)
}

type ConsoleLoggerFactory struct{}

func (cl *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
