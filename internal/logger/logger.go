package logger

import (
	"log"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *SystemLogger
)

type SystemLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	panicLogger *log.Logger
}

func NewSystemLogger() *SystemLogger {
	return &SystemLogger{
		infoLogger:  log.New(os.Stdout, "[ INF ] ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, "[ WRN ] ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "[ ERR ] ", log.Ldate|log.Ltime|log.Lshortfile),
		panicLogger: log.New(os.Stderr, "[ PNC ] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (c *SystemLogger) Info(msg string) {
	c.infoLogger.Println(msg)
}

func (c *SystemLogger) Warn(msg string) {
	c.warnLogger.Println(msg)
}

func (c *SystemLogger) Error(msg string) {
	c.errorLogger.Println(msg)
}

func (c *SystemLogger) Panic(msg string) {
	c.panicLogger.Println(msg)
}

func Get() *SystemLogger {
	once.Do(func() {
		logger = NewSystemLogger()
	})

	return logger
}
