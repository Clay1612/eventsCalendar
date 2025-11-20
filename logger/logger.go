package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	File             *os.File
	userInputLogger  *log.Logger
	userOutputLogger *log.Logger
	errorLogger      *log.Logger
)

func Init() error {
	var err error

	File, err = os.OpenFile("save/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("logger init func error: %w", err)
	}

	userInputLogger = log.New(File, "INPUT: ", log.Ldate|log.Ltime|log.Lshortfile)
	userOutputLogger = log.New(File, "OUTPUT: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(File, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

func Input(msg string) {
	_ = userInputLogger.Output(2, msg)
}

func Output(msg string) {
	_ = userOutputLogger.Output(2, msg)
}

func Error(msg string) {
	_ = errorLogger.Output(2, msg)
}
