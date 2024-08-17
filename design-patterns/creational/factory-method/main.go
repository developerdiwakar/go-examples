package main

import (
	"fmt"
	"go-exercises/design-patterns/creational/factory-method/logger"
	"time"
)

func main() {
	var factory logger.LoggerFactory

	message := fmt.Sprintf("%v %s\n", time.DateTime, "emerg: Logging from FileLogger.")
	// Example create a Console Logger
	factory = &logger.FileLoggerFactory{}
	fileLogger := factory.CreateLogger()
	fileLogger.Log(message)

	message = fmt.Sprintf("%v %s", time.DateTime, "emerg: Logging from ConsoleLogger.")
	// Example create a Console Logger
	factory = &logger.ConsoleLoggerFactory{}
	consoleLogger := factory.CreateLogger()
	consoleLogger.Log(message)
}
