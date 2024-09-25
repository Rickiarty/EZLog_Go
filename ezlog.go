package ezlog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func EZLog(detailedMsg string, params ...interface{}) (bool, string) {
	now := time.Now()

	// default values
	selfDefinedCode := 0
	selfDefinedTag := "INFO"
	logDirPath, _ := filepath.Abs("./_EZLOG_")

	// Set customized parameters, if any.
	if len(params) > 0 {
		if code, ok := params[0].(int); ok {
			selfDefinedCode = code
		}
	}
	if len(params) > 1 {
		if tag, ok := params[1].(string); ok {
			selfDefinedTag = tag
		}
	}
	if len(params) > 2 {
		if path, ok := params[2].(string); ok {
			logDirPath, _ = filepath.Abs(path)
		}
	}

	if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
		err := os.MkdirAll(logDirPath, os.ModePerm)
		if err != nil {
			fmt.Println(logDirPath)
			fmt.Println("[EZLog] log failed.")
			fmt.Println(err)
			return false, logDirPath
		}
	}

	logFilePath := filepath.Join(logDirPath, fmt.Sprintf("%s_log.txt", now.Format("20060102")))
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(logFilePath)
		fmt.Println("[EZLog] log failed.")
		fmt.Println(err)
		return false, logDirPath
	}

	defer file.Close()

	logMessage := fmt.Sprintf("[%d：%s]\n%s\n== %s (UTC：%s) ==\n",
		selfDefinedCode,
		selfDefinedTag,
		detailedMsg,
		now.Format("2006/01/02T15:04:05.000000"),
		now.UTC().Format("2006/01/02T15:04:05.000000"))

	if _, err := file.WriteString(logMessage); err != nil {
		fmt.Println(logMessage)
		fmt.Println("[EZLog] log failed.")
		fmt.Println(err)
		return false, logFilePath
	}

	return true, logDirPath
}
