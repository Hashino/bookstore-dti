package logger

import (
	"log"
	"os"
)

var (
	logFile *os.File
	Logger  *log.Logger
)

func Init() error {
	var err error
	logFile, err = os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	Logger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func LogError(err error) {
	if Logger != nil && err != nil {
		Logger.Println(err)
	}
}

func Close() {
	if logFile != nil {
		logFile.Close()
	}
}
