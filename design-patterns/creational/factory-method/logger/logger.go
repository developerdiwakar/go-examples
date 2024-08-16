package logger

// Logger interface defines methods that all loggers must implement
type Logger interface {
	Log(message string)
}

// LoggerFactory is the creator interface that has the factory method
type LoggerFactory interface {
	CreateLogger() Logger
}
