package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File
	DefaultPrefix = ""
	DefaultCallerDepth = 2
	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init(){
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(format string, args ...interface{}){
	setPrefix(DEBUG)
	logger.Printf(format, args)
}

func Info(format string, args ...interface{}){
	setPrefix(INFO)
	logger.Printf(format, args)
}

func Warn(format string, args ...interface{}){
	setPrefix(WARNING)
	logger.Printf(format, args)
}

func Error(format string, args ...interface{}){
	setPrefix(ERROR)
	logger.Printf(format, args)
}

func Fatalf(format string, args ...interface{}){
	setPrefix(FATAL)
	logger.Printf(format, args)
	os.Exit(1)
}

func setPrefix(level Level){
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]",levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}


