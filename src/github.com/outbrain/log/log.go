/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package log

import (
	"time"
	"fmt"
	"os"
	"errors"
)

// LogLevel indicates the severity of a log entry
type LogLevel int

func (this LogLevel) String() string {
	switch this {
		case FATAL: return "FATAL"
		case CRITICAL: return "CRITICAL"
		case ERROR: return "ERROR"
		case WARNING: return "WARNING"
		case NOTICE: return "NOTICE"
		case INFO: return "INFO"
		case DEBUG: return "DEBUG"
	}
	return "unknown"
}

const (
	FATAL LogLevel = iota
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

const timeFormat = "2006-01-02 15:04:05"

// globalLogLevel indicates the global level filter for all logs (only entries with level equals or higher 
// than this value will be logged)
var globalLogLevel LogLevel = DEBUG

// SetLevel sets the global log level. Only entries with level equals or higher than
// this value will be logged
func SetLevel(logLevel LogLevel) {
	globalLogLevel = logLevel
}

// GetLevel returns current global log level
func GetLevel() LogLevel {
	return globalLogLevel
}

// logFormattedEntry nicely formats and emits a log entry
func logFormattedEntry(logLevel LogLevel, message string, args ...interface{}) string {
	if logLevel > globalLogLevel {
		return ""
	} 
	entryString := fmt.Sprintf("%s %s %s", time.Now().Format(timeFormat), logLevel, fmt.Sprintf(message, args...))
	fmt.Fprintln(os.Stderr, entryString)
	return entryString
}

// logEntry emits a formatted log entry 
func logEntry(logLevel LogLevel, message string, args ...interface{}) string {
	entryString := message
	for _, s := range args {
		entryString += fmt.Sprintf(" %s", s)
	}
	return logFormattedEntry(logLevel, entryString)
}

// logErrorEntry emits a log entry based on given error object
func logErrorEntry(logLevel LogLevel, err error) error {
	if err == nil {
		// No error
		return nil;
	}
	entryString := fmt.Sprintf("%+v", err)
	logEntry(logLevel, entryString)
	return err	
}

func Debug(message string, args ...interface{}) string {
	return logEntry(DEBUG, message, args...)
}

func Debugf(message string, args ...interface{}) string {
	return logFormattedEntry(DEBUG, message, args...)
}

func Info(message string, args ...interface{}) string {
	return logEntry(INFO, message, args...)
}

func Infof(message string, args ...interface{}) string {
	return logFormattedEntry(INFO, message, args...)
}

func Notice(message string, args ...interface{}) string {
	return logEntry(NOTICE, message, args...)
}

func Noticef(message string, args ...interface{}) string {
	return logFormattedEntry(NOTICE, message, args...)
}

func Warning(message string, args ...interface{}) error {
	return errors.New(logEntry(WARNING, message, args...))
}

func Warningf(message string, args ...interface{}) error {
	return errors.New(logFormattedEntry(WARNING, message, args...))
}

func Error(message string, args ...interface{}) error {
	return errors.New(logEntry(ERROR, message, args...))
}

func Errorf(message string, args ...interface{}) error {
	return errors.New(logFormattedEntry(ERROR, message, args...))
}

func Errore(err error) error {
	return logErrorEntry(ERROR, err)
}

func Critical(message string, args ...interface{}) error {
	return errors.New(logEntry(CRITICAL, message, args...))
}

func Criticalf(message string, args ...interface{}) error {
	return errors.New(logFormattedEntry(CRITICAL, message, args...))
}

func Criticale(err error) error {
	return logErrorEntry(CRITICAL, err)
}

// Fatal emits a FATAL level entry and exists the program
func Fatal(message string, args ...interface{}) error {
	logEntry(FATAL, message, args...)
	os.Exit(1)
	return errors.New(logEntry(CRITICAL, message, args...))
}

// Fatalf emits a FATAL level entry and exists the program
func Fatalf(message string, args ...interface{}) error {
	logEntry(FATAL, message, args...)
	os.Exit(1)
	return errors.New(logFormattedEntry(CRITICAL, message, args...))
}

// Fatale emits a FATAL level entry and exists the program
func Fatale(err error) error {
	logErrorEntry(FATAL, err)
	os.Exit(1)
	return err
}

