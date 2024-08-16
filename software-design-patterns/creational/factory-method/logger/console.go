package logger

import (
	"fmt"
)

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Printf("%s\n", message)
}

type ConsoleLoggerFactory struct{}

func (cl *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
