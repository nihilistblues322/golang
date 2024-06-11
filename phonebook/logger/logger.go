package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	warnLogger = log.New(os.Stdout, "WARN: ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	infoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
)

func Info(msg string) {
	infoLogger.Println(msg)
}
func Warn(err error, messages ...string) {
	if err == nil {
		return
	}

	message := "Error"
	if len(messages) > 0 {
		message = fmt.Sprintf("%s: %s", message, strings.Join(messages, " "))
	}

	warnLogger.Printf("%s: %s", message, err)
}
