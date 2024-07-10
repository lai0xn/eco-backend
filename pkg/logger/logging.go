package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type MyLogger struct {
	zerolog.Logger
}

var Logger MyLogger

func NewLogger() MyLogger {
	// create output configuration
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Format level: fatal, error, debug, info, warn
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	// format error
	output.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}

	zerolog := zerolog.New(output).With().Caller().Timestamp().Logger()
	Logger = MyLogger{zerolog}
	return Logger
}

func LogInfo() *zerolog.Event {
	return Logger.Info()
}

func LogError() *zerolog.Event {
	return Logger.Error()
}

func LogDebug() *zerolog.Event {
	return Logger.Debug()
}

func LogWarn() *zerolog.Event {
	return Logger.Warn()
}

func LogFatal() *zerolog.Event {
	return Logger.Fatal()
}
