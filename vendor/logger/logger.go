package logger

// LogData data received in each logger
type LogData struct {
	FontNames []string
	CSSFile   string
}

type loggerFunc func(LogData)

var loggers []loggerFunc

// Log log information after downloading the fonts
func Log(data LogData) {
	for _, logger := range loggers {
		logger(data)
	}
}

// Register register a logger to be called
func Register(fn loggerFunc) {
	loggers = append(loggers, fn)
}
