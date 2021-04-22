package glinit

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func InitLog() {
	var logPath string = "log/goleaf.log"

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	os.MkdirAll(filepath.Dir(logPath), 0700)
	logFile, err := os.OpenFile("log/goleaf.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Panic("Failed to create log file ", err)
	}

	mw := io.MultiWriter(logFile, os.Stdout)
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(mw)

	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debug("Init log successfully setup")
}
