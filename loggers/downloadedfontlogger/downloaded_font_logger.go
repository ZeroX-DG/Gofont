package downloadedfontlogger

import (
	"fmt"
	"logger"
	"strings"
)

// Log log the info
func Log(data logger.LogData) {
	message := "Downloaded font: "
	if len(data.FontNames) > 1 {
		message = "Downloaded fonts: "
	}
	fmt.Println(message, strings.Join(data.FontNames, ", "))
}
