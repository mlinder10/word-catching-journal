package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/mlinder10/wcj/config"
)

var logPath = config.Env.LogPath

func Write(msg string) {
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	currentTime := fmt.Sprintf("%v", time.Now())

	if _, err := f.Write([]byte(currentTime + ": " + msg + "\n")); err != nil {
		fmt.Println(err.Error())
	}

	if err := f.Close(); err != nil {
		fmt.Println(err.Error())
	}
}

func GetHistory() ([]byte, error) {
	return os.ReadFile(logPath)
}
