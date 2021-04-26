package glutil

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var doOnceLoadEnv sync.Once

func GetEnv(key, fallback string) string {
	doOnceLoadEnv.Do(func() {
		if _, err := os.Stat(".env"); err == nil {
			logrus.Debug(".env file available load it into system")
			err := godotenv.Load()
			if err != nil {
				logrus.Warn(".env is failed to load ", err)
			}
			logrus.Debug("Load .env done")
		}
	})

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
