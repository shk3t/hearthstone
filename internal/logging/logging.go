package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

var DebugLogger *log.Logger

func Init() {
	debugLogFile, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Can't open log file")
	}
	DebugLogger = log.New(debugLogFile, "DEBUG:", log.LstdFlags|log.Lshortfile)

}

func Deinit() {
	writer := DebugLogger.Writer()
	writeCloser, ok := writer.(io.WriteCloser)
	if ok {
		err := writeCloser.Close()
		if err != nil {
			panic("Can't close log file")
		}
		fmt.Println("CLOSED!")
	}
}