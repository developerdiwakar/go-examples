package main

import (
	"fmt"
	"go-exercises/software-design-patterns/creational/factory-method/logger"
	"time"
)

func main() {
	var factory logger.LoggerFactory

	message := fmt.Sprintf("%v %s\n", time.DateTime, "Error: Logging from console.")
	// Example create a Console Logger
	factory = &logger.FileLoggerFactory{}
	fileLogger := factory.CreateLogger()
	fileLogger.Log(message)

	message = fmt.Sprintf("%v %s", time.DateTime, "Info: Logging from console.")
	// Example create a Console Logger
	factory = &logger.ConsoleLoggerFactory{}
	consoleLogger := factory.CreateLogger()
	consoleLogger.Log(message)
}
