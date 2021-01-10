package logger

import (
	"log"
	"os"
)

type Logger struct {
	Infolog *log.Logger
	Errlog  *log.Logger
}

func New() *Logger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		Infolog: infoLog,
		Errlog:  errLog,
	}
}
