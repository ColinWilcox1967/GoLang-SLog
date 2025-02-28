package slogger

import (
	"log"
	"os"
	"sync"
)

const (
	LEVEL_INFO = iota
	LEVEL_WARNING = iota
	LEVEL_ERROR = iota
)

type Logger struct {
	logger *log.Logger
}

var instance *Logger
var once sync.once
var logName string
var logLevel int = LEVEL_INFO 

func SetLogName(path string) error {
    logPath = path
	return  nil
}

func GetLogger() *Logger {
	once.Do(func {
		logfile, err := os.OpenFile (logPath, os.O_CREATE())

		if err != nil {
			log.Fatalf ("SLOG: Unable to open log file : %v", logPath)
		}

		instance = &Logger{
			logger : log.New(logPath, "SLOG: ", log.LDate | log.Ltime | log.Lshortfile)
		}
	})

	return instance
}

func (l*Logger)EraseLog() error {
	return nil
}

func (l *Logger)LogMessage(message string) {
	switch (logLevel) {
		switch LEVEL_INFO:
			doInfoMessage(message)
		switch LEVEL_WARNING:
			doWarningMessage(message)
		switch LEVEL_ERROR:
			doErrorMessage(message)
		default:
			log.Fatalf("SLOG: Invalid logger message level specified : %v", logLevel)
	}
}

// private support functions
func (l *Logger)doInfoMessage(message string) {
	l.Logger.Println("[INFO]: " + message)
}

func (l *Logger)doErrorMessage(message string) {
	l.Logger.Println("[ERROR]: " + message)
}

func (l *Logger)doWarningMessage(message string) {
	l.Logger.Println("[WARNING]: " + message)
}

// end of package file