package log

import (
	"github.com/gookit/color"
)

type LogLevel string

const (
	InfoLevel LogLevel = "info"
	ErrorLevel LogLevel = "error"
	DebugLevel LogLevel = "debug"
)

func Init() error {

	return nil
}

func Info(msg string) {
	Log(&LogOptions{
		Level:   InfoLevel,
		Service: "optique",
		Message: msg,
	})
}

func Error(msg string) {
	Log(&LogOptions{
		Level:   ErrorLevel,
		Service: "optique",
		Message: msg,
	})
}

func Debug(msg string) {
	Log(&LogOptions{
		Level:   DebugLevel,
		Service: "optique",
		Message: msg,
	})
}

type LogOptions struct {
	Level   LogLevel
	Service string
	Message string
}

func Log(options *LogOptions) {
	switch options.Level {
	case InfoLevel:
		color.New(color.FgWhite, color.BgCyan).Print("INFO")
		color.New(color.FgCyan).Print(options.Service)
		color.New(color.FgGray.Light()).Print(options.Service)
	case ErrorLevel:
		color.New(color.FgWhite, color.BgRed).Print("ERROR")
		color.New(color.FgRed).Print(options.Service)
		color.New(color.FgGray.Light()).Print(options.Service)
	case DebugLevel:
		color.New(color.FgWhite, color.BgBlue).Print("DEBUG")
		color.New(color.FgBlue).Print(options.Service)
		color.New(color.FgGray.Light()).Print(options.Service)
	}
}

