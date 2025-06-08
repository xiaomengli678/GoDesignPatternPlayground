package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	file *os.File
}

var (
	instance *Logger
	once     sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		f, err := os.Create("logfile.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		instance = &Logger{file: f}
	})
	return instance
}

func (l *Logger) Log(message string) {
	t := time.Now()
	fmt.Fprintf(l.file, "[%s] %s\n", t, message)
}

func (l *Logger) Close() {
	l.file.Close()
}

func main() {
	logger := GetLogger()
	defer logger.Close()

	logger.Log("hello world")
	logger.Log("this is a simple singleton logger")
}
