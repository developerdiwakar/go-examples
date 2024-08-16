package logger

import (
	"errors"
	"log"
	"os"
)

type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	// Create Logger file if not exist
	logFile := createLoggerFile()

	defer logFile.Close()

	// Writing Log data
	if _, err := logFile.WriteString(message); err != nil {
		log.Fatalf("Error writing contents to the file: %v", err)
	}
}

type FileLoggerFactory struct{}

func (fl *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

func createLoggerFile() *os.File {
	dir := "logs"
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(dir, os.ModePerm); err != nil {
			log.Println(err)
		}
	}

	fileName := dir + string(os.PathSeparator) + "logger.log"
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatalf("Error Openning and Creating file: %v", err)
	}

	return logFile
}
