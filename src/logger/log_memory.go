package logger

import (
	"log"
	"strings"
)

type Logger struct {
	history []string
}

var Log = &Logger{make([]string, 0)}

func (l *Logger) Println(message string) {
	log.Println(message)
	l.history = append(l.history, message)
}

func (l *Logger) CheckMessage(message string) bool {
	for _, m := range l.history {
		if strings.EqualFold(m, message) {
			return true
		}
	}
	return false
}
