package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG level = iota
	INFO
	WARNING
	ERROR
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setProfix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setProfix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setProfix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setProfix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setProfix(ERROR)
	logger.Fatalln(v)
}

func setProfix(level level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
