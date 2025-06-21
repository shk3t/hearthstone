package log

import (
	"io"
	"log"
	"os"
)

var DebugLogger *log.Logger
var DLog func(...any)

func Init() {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		panic("Can't create \"logs\" directory")
	}
	debugLogFile, err := os.OpenFile("logs/debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Can't open \"debug.log\" file")
	}

	DebugLogger = log.New(debugLogFile, "", log.LstdFlags|log.Lshortfile)
	DLog = DebugLogger.Println
}

func Deinit() {
	writer := DebugLogger.Writer()
	writeCloser, ok := writer.(io.WriteCloser)
	if ok {
		err := writeCloser.Close()
		if err != nil {
			panic("Can't close log file")
		}
	}
}