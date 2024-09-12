package utils

import (
	"goApp/config"
	"io"
	"log"
	"os"
	"path/filepath"
)

func LoggingSettings(config *config.ConfigList) {
	logFile := config.LOG_FILE + ".log"
	// ディレクトリを作成（存在しない場合）
	dir := filepath.Dir(logFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	//ログファイルを開く
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %v", logFile, err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
