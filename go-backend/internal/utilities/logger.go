package utilities

import (
	"log"
	"os"
)

type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	WarnLogger  *log.Logger
}

func SetUpLogger() {
	file, _ := os.Create("log.log")
	log.SetOutput(file)
}

func (l *Logger) Info(message string) {
	l.InfoLogger.Println(message)
}

func (l *Logger) Error(message string) {
	l.ErrorLogger.Println(message)
}

func (l *Logger) Warn(message string) {
	l.WarnLogger.Println(message)
}
