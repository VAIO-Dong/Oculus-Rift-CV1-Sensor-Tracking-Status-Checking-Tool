package logger

import (
	"fmt"

	"os"
	"time"
)

const (
	INFO = iota
	DEBUG
	WARN
	ERROR
	FATAL
	Unexpected
)

var file bool
var logfile *os.File

func LogTOfile(path string) {
	if IsExist(path) {
		f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
		logfile = f
	} else {
		f, _ := os.Create(path)
		logfile = f
	}

	file = true

	// todo MultiWriter
}

func CloseFile() {
	file = false
	err := logfile.Close()
	if err != nil {
		PrintErr(err.Error())
	}
}

func LogToCmd() {
}

func CmdClose() {
}

func PrintLog(log string, level int) {
	var info string
	switch level {
	case DEBUG:
		info = "DEBUG"
	case INFO:
		info = "INFO"
	case WARN:
		info = "WARN"
	case ERROR:
		info = "ERROR"
	case FATAL:
		info = "FATAL"
	case Unexpected:
		info = "Unexpected"
	default:
		info = "Unexpected"
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	if file {
		fmt.Printf("[%s][%s]: %s \n", t, info, log)
		_, _ = fmt.Fprintf(logfile, "[%s][%s]: %s \n", t, info, log)
	} else {
		fmt.Printf("[%s][%s]: %s \n", t, info, log)
	}

	//cqp.AddLog(cqp.Debug,"microQ",log)
}

func PrintInfo(log string) {
	PrintLog(log, INFO)
}

func PrintErr(err string) {
	PrintLog(err, ERROR)
}

func PrintFatal(log string) {
	PrintLog(log, FATAL)
}

func IsExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	} else {
		return true
	}
}

// test
func TestLog() {
	time.Sleep(5 * time.Second)
	PrintInfo("1")
	time.Sleep(5 * time.Second)
	PrintInfo("test")
	PrintInfo("1")
}
