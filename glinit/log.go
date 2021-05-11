package glinit

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/frediansah/goleafcore/glconstant"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/sirupsen/logrus"
)

func InitLog() {
	var logPath string = glutil.GetEnv(glconstant.ENV_LOG_FILE, glconstant.LOG_FILE_DEFAULT)
	var logLevel string = strings.ToLower(glutil.GetEnv(glconstant.ENV_LOG_LEVEL, glconstant.LOG_LEVEL_DEFAULT))

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	os.MkdirAll(filepath.Dir(logPath), 0700)
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Panic("Failed to create log file ", err)
	}

	mw := io.MultiWriter(logFile, os.Stdout)
	logrus.SetOutput(mw)

	lgLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		lgLevel = logrus.DebugLevel
	}
	logrus.SetLevel(lgLevel)
	logrus.Debug("Init log successfully setup")
}
