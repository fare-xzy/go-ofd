package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
)

var logger zerolog.Logger

var (
	logFolder = filepath.Join(os.TempDir(), "ofd")
	logFile   = "ofd.log"
	logLevel  = "debug"
)

func Init() {

	// log level
	switch logLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	default:
		panic(fmt.Sprintf("unknown log level: %s", "debug"))
	}
	_, err := os.Stat(logFolder)
	if err != nil {
		os.Mkdir(logFolder, 0666)
	}
	remoteLogPath := filepath.Join(logFolder, "remote")
	_, err = os.Stat(remoteLogPath)
	if err != nil {
		os.Mkdir(remoteLogPath, 0666)
	}
	// log file
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	fileWriter, err := os.OpenFile(filepath.Join(logFolder, logFile), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("open log file failed: %s", err))
	}
	multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)
	logger = zerolog.New(multi).With().Timestamp().Logger()
}
