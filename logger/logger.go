package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	withTimestamp bool
	timeFormat    string
}

const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

// New creates a new Logger instance
func New(withTimestamp bool) *Logger {
	return &Logger{
		withTimestamp: withTimestamp,
		timeFormat:    "2006-01-02 15:04:05",
	}
}

// SetTimeFormat allows customizing the timestamp format
func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

func (l *Logger) colorPrint(color string, format string, a ...any) {
	if l.withTimestamp {
		timestamp := time.Now().Format(l.timeFormat)
		format = fmt.Sprintf("[%s] %s", timestamp, format)
	}
	coloredFormat := fmt.Sprintf("%s%s%s", color, format, reset)
	fmt.Printf(coloredFormat, a...)
}

func (l *Logger) Success(format string, a ...any) {
	l.colorPrint(green, format, a...)
}

func (l *Logger) Changed(format string, a ...any) {
	l.colorPrint(yellow, format, a...)
}

func (l *Logger) Failed(format string, a ...any) {
	l.colorPrint(red, format, a...)
}

func (l *Logger) Fatal(format string, a ...any) {
	l.colorPrint(red, format, a...)
	os.Exit(1)
}

func (l *Logger) Ignore(format string, a ...any) {
	l.colorPrint(yellow, format, a...)
}

func (l *Logger) Printf(format string, a ...any) {
	if l.withTimestamp {
		timestamp := time.Now().Format(l.timeFormat)
		format = fmt.Sprintf("[%s] %s", timestamp, format)
	}
	fmt.Printf(format, a...)
}

func (l *Logger) Println(a ...any) {
	if l.withTimestamp {
		timestamp := time.Now().Format(l.timeFormat)
		fmt.Printf("[%s] ", timestamp)
	}
	fmt.Println(a...)
}
