package glinit

import (
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"gitlab.com/stsgoleaf/goleafcore/glutil"
)

var config fiber.Config
var app *fiber.App = nil
var doOnceServer sync.Once
var serverHost *string

var SERVER_HOST string = glutil.GetEnv("SERVER_HOST", "0.0.0.0")
var SERVER_PORT string = glutil.GetEnv("SERVER_PORT", "5005")

func InitServer() *fiber.App {

	doOnceServer.Do(func() {
		logrus.Debug("Do once Init server")

		readTimeoutSecondsCount, _ := strconv.Atoi(glutil.GetEnv("SERVER_READ_TIMEOUT", "120"))

		config = fiber.Config{
			ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		}

		app = fiber.New(config)
	})

	return app
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	logrus.Debug("Call start server")
	serverUrl := SERVER_HOST + `:` + SERVER_PORT

	logrus.Info("Application will be running on ", serverUrl)
	if err := a.Listen(serverUrl); err != nil {
		logrus.Panic("Oops... Server is not running! Reason: %v", err)
	}

}

func EndServer(db *pgxpool.Pool) {
	logrus.Info("Ending server")
	db.Close()
	logrus.Info("Database closed")
}
