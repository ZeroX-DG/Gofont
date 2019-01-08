package logger

// LogData data received in each logger
type LogData struct {
	FontNames []string
}

type loggerFunc func(LogData)

var loggers []loggerFunc

// Log log information after downloading the fonts
func Log(fontNames []string) {
	data := LogData{FontNames: fontNames}
	for _, logger := range loggers {
		logger(data)
	}
}

// Register register a logger to be called
func Register(fn loggerFunc) {
	loggers = append(loggers, fn)
}
