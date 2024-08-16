package logger

// Logger interface defines methods that all loggers must implement
type Logger interface {
	Log(message string)
}

type LoggerFactory interface {
	CreateLogger() Logger
}
