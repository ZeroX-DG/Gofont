package generatecsslogger

import (
	"fmt"
	"logger"
)

// Log log the info
func Log(data logger.LogData) {
	if data.CSSFile != "" {
		message := "Generated a css file at: " + data.CSSFile
		fmt.Println(message)
	}
}
