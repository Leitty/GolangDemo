package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogfilePath() string{
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string{
	prefixPath := getLogFileFullPath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch  {
	case os.IsNotExist(err):
		mkDir(getLogFileFullPath())
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}


func mkDir(filename string) {
	dir, _ := os.Getwd()
	err := os.Mkdir(dir+"/"+filename, os.ModePerm)
	if err != nil {
		log.Fatalf("Fail to create directory: %v", err)
	}
}