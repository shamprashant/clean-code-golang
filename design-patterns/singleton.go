/*
Implement a thread-safe Logger singleton in Go that prints logs only if log level is "DEBUG".
You should be able to call it from anywhere, but only one instance must exist.
*/

package main

import (
	"fmt"
	"sync"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[l]
}

type Logger struct {
	level LogLevel
	mu    sync.RWMutex
}

func (l *Logger) Log(level LogLevel, message string) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if level >= l.level {
		fmt.Println("Console Log: ", message)
	}
}

func (l *Logger) setLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.level = level
}

var (
	instance *Logger
	once     sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		fmt.Println("Creating new Logger instance")
		instance = &Logger{}
	})
	return instance
}

func main() {
	logger1 := GetLogger()
	logger2 := GetLogger()

	logger1.Log(DEBUG, "Log from logger1")
	logger2.Log(DEBUG, "Log from logger2")

	logger1.setLevel(INFO)

	logger1.Log(DEBUG, "Log DEBUG level from logger1")
	logger1.Log(INFO, "Log INFO level from logger1")

	logger2.Log(DEBUG, "Log DEBUG level from logger2")
	logger2.Log(INFO, "Log INFO level from logger2")

	fmt.Println(DEBUG)
}
