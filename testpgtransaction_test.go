package goleafcore

import (
	"encoding/json"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/frediansah/goleafcore/gldb"
	"github.com/frediansah/goleafcore/glentity"
	"github.com/frediansah/goleafcore/glinit"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type TableTestTx struct {
	glentity.BaseEntityTs

	TableId int64  `json:"tableId" gleaf:"pk"`
	Name    string `json:"name"`
}

func TestTransactionInsert(t *testing.T) {
	glinit.InitLog()
	glinit.InitDb()

	err := gldb.BeginTrx(func(trx pgx.Tx) error {
		tableName := initTableTransation(trx)

		if errInit := initDataTransaction(trx, tableName, 10); errInit != nil {
			return errInit
		}

		resultList, errSelect := selectDataTransaction(trx, tableName)
		if errSelect != nil {
			return errSelect
		}
		jsonByte, _ := json.Marshal(resultList)

		logrus.Debug("Result list : ", string(jsonByte))

		return errors.New("test rollback")
	})

	logrus.Debug("Result trx overall :", err)
}

func initDataTransaction(tx pgx.Tx, tableName string, count int) error {
	query := `INSERT INTO ` + tableName + `( table_id, name, create_timestamp, update_timestamp, create_user_id, update_user_id, version ) ` +
		` VALUES ($1, $2, $3, $4, $5, $6, $7)`

	for i := 0; i < count; i++ {
		err := gldb.ExecTx(tx, query,
			(i + 1), "Name "+strconv.Itoa(i+1),
			time.Now(), time.Now(), -1, -1, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func selectDataTransaction(tx pgx.Tx, tableName string) ([]*TableTestTx, error) {
	query := `SELECT * FROM ` + tableName
	var resultList []*TableTestTx

	err := gldb.SelectTx(tx, &gldb.ReturnSelect{
		Result: &resultList,
	}, query)

	return resultList, err
}

func initTableTransation(trx pgx.Tx) string {
	tableName := `t_test_tabletx_` + glutil.DateTimeNow()

	sql := `CREATE TABLE IF NOT EXISTS ` + tableName + `( ` +
		` table_id 	bigint, ` +
		` name 		text, ` +
		` create_timestamp 		timestamptz, ` +
		` update_timestamp 		timestamptz, ` +
		` create_user_id 		bigint, ` +
		` update_user_id 		bigint, ` +
		` version 		bigint ` +
		`)`

	gldb.ExecTx(trx, sql)

	return tableName
}
