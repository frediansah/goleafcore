package glinit

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/frediansah/goleafcore/glutil"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var DB_CTX context.Context = context.Background()
var pool *pgxpool.Pool = nil
var doOnce sync.Once

func InitDb() *pgxpool.Pool {
	db := GetDB()
	err := db.Ping(DB_CTX)
	if err != nil {
		logrus.Fatal("Database failed to connect : ", err)
	}

	return db
}

func GetDB() *pgxpool.Pool {
	doOnce.Do(func() {
		DB_USER := glutil.GetEnv("DB_USER", "sts")
		DB_PASSWORD := glutil.GetEnv("DB_PASSWORD", "Awesome123!")
		DB_NAME := glutil.GetEnv("DB_NAME", "erp_cloud")
		DB_PORT := glutil.GetEnv("DB_PORT", "5432")
		DB_HOST := glutil.GetEnv("DB_HOST", "172.17.0.1")
		DB_POOL_MAX_CONNS := glutil.GetEnv("DB_POOL_MAX_CONNS", "5")
		DB_POOL_MIN_CONNS := glutil.GetEnv("DB_POOL_MIN_CONNS", "1")
		DB_APP_NAME := glutil.GetEnv("DB_APP_NAME", "GoleafBoilerPlate")

		dbUrl := "postgresql://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT +
			"/" + DB_NAME +
			"?pool_max_conns=" + DB_POOL_MAX_CONNS +
			"&pool_min_conns=" + DB_POOL_MIN_CONNS +
			"&application_name=" + DB_APP_NAME

		logrus.Debug("Connecting database to ", dbUrl)

		var err error
		pool, err = pgxpool.Connect(DB_CTX, dbUrl)
		logrus.Println("Starting pool with max connection ", pool.Stat().MaxConns())

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			logrus.Fatal("Database connection error")
		}
	})

	return pool
}
