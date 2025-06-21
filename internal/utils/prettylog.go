package prettylog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ANSI color codes
const (
	gray   = "\033[90m"
	violet = "\033[35m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
	orange = "\033[38;5;208m"
	red    = "\033[31m"
	darkred = "\033[38;5;88m"
	reset  = "\033[0m"
)

type LogType string

const (
	LogMessage LogType = "MESSAGE"
	LogInfo    LogType = "INFO"
	LogWarning LogType = "WARNING"
	LogError   LogType = "ERROR"
	LogFatal   LogType = "FATAL"
)

var logColors = map[LogType]string{
	LogMessage: blue,
	LogInfo:    cyan,
	LogWarning: orange,
	LogError:   red,
	LogFatal:   darkred,
}

// PrintPretty prints a formatted log line with color and timestamp.
func PrintPretty(function, message string, logType LogType) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	dateColor := gray
	funcColor := violet
	typeColor, ok := logColors[logType]
	if !ok {
		typeColor = blue
	}

	var typeStr string
	if logType != LogMessage && logType != LogInfo {
		typeStr = fmt.Sprintf("[%s%s%s] ", typeColor, string(logType), reset)
	}

	// Print: date [function] [TYPE] message
	fmt.Printf("%s%s%s [%s%s%s] %s%s%s%s\n",
		dateColor, timestamp, reset,
		funcColor, function, reset,
		typeStr,
		typeColor, message, reset,
	)

	if logType == LogFatal {
		os.Exit(1)
	}
}

// PrintError prints an error message with error details.
func PrintError(function string, err error, message string, fatal bool) {
	logType := LogError
	if fatal {
		logType = LogFatal
	}
	msg := message
	if err != nil {
		msg = fmt.Sprintf("%s: %s", message, strings.TrimSpace(err.Error()))
	}
	PrintPretty(function, msg, logType)
}
