package logger

/*
import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func Init(logFilePath string, stdoutFlg bool) {
	if len(logFilePath) > 0 && stdoutFlg == false {
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Failed to open log file:", err)
		}
		Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else if len(logFilePath) == 0 && stdoutFlg == true {
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else if len(logFilePath) > 0 && stdoutFlg == true {
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Failed to open log file:", err)
		}
		multiWriter := io.MultiWriter(os.Stdout, file)
		Info = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

*/
