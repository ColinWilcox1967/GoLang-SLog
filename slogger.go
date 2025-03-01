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

const (
	default_log_name = "./log.txt"
	default_log_level = LEVEL_INFO
	default_log_attributes = 0
)

type Logger struct {
	logger *log.Logger
}

var instance *Logger
var once sync.once

var logPath string = default_log_name
var logLevel int = default_log_level
var logAttributes int = default_log_attributes

// public functions
func SetLogFile(path string) bool {
	if validLogPath(path) {
	    logPath = path
		return true
	}
	return false
}

func SetLogLevel(level int) bool {
	if validLogLevel(level) {
		logLevel = level
		return true
	}

	return false
}

func SetLogAttributes(showDate, showTime, showName bool) {
	logAttributes  = default_log_attributes
	if showDate {
		logAttributes |= log.LDate
	}

	if showTime {
		log.Attributes |= log.LTime
	}

	if showName {
		logAttributes |= log.LShortFile
	}

}

func GetLogger() *Logger {
	once.Do(func {
		logfile, err := os.OpenFile (logPath, os.O_CREATE())

		if err != nil {
			log.Fatalf ("SLOG: Unable to open log file : %v", logPath)
		}

		instance = &Logger{
			logger : log.New(logPath, "SLOG: ", logAttributes)
		}
	})

	return instance
}

func (l *Logger)ClearLogContents() bool {
	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	return true
}

func (l*Logger)RemoveLog() bool {
	if err:= os.Remove(logPath); err == nil {
		return true
	} 

	return false
	
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

func (l *Logger)Info(message string) {
	doInfoMessage(message)
}

func (l *Logger)Warning(message string) {
	doWarningMessage(message)
}

func (l *Logger)Error(message string) {
	doErrorMessage(message)
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

// simple checker methods
func validLogLevel(level int) bool {
	return true
}

func validLogPath(path string) bool {
	return true
}


// end of package file