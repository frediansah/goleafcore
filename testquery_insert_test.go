package goleafcore

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/frediansah/goleafcore/gldb"
	"github.com/frediansah/goleafcore/glentity"
	"github.com/frediansah/goleafcore/glinit"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type TableTestInsertTx struct {
	glentity.BaseEntityTs

	TableId int64  `json:"tableId" gleaf:"pk"`
	Name    string `json:"name"`
}

type TableTestInsertTxWithSeq struct {
	glentity.BaseEntityTs

	TableId int64  `json:"tableId" gleaf:"pk seq"`
	Name    string `json:"name"`
}

func TestInsertQuery(t *testing.T) {
	glinit.InitLog()
	db := glinit.InitDb()
	defer db.Close()

	gldb.BeginTrx(func(trx pgx.Tx) error {

		tableName := initTableTestInsertTx(trx)
		err := insertTableTestInsertTx(trx, tableName)
		logrus.Debug("Error insert table ", err)

		var findBy TableTestInsertTx
		errFindBy := gldb.FindByPkTx(trx, &findBy, tableName, 2)
		logrus.Debug("Error find by ", errFindBy)

		// tableWithSeq := initTableTestInsertTxWithTx(trx)
		// errSeq := insertTableTestInsertTxWithSeq(trx, tableWithSeq)
		// logrus.Debug("Error insert table with seq", errSeq)

		// selectTestInsertTableWithSeq(trx, tableWithSeq)

		return nil
	})

}

func selectTestInsertTableWithSeq(tx pgx.Tx, tableName string) error {
	query := `SELECT * FROM ` + tableName

	var list []*TableTestInsertTxWithSeq

	err := gldb.SelectTx(tx, &gldb.ReturnSelect{
		Result: &list,
	}, query)

	jsonByte, _ := json.Marshal(list)
	logrus.Debug("Data : ", string(jsonByte))

	return err
}

func insertTableTestInsertTx(trx pgx.Tx, tableName string) error {
	err := gldb.InsertTx(trx, TableTestInsertTx{
		BaseEntityTs: glentity.BaseEntityTs{
			CreateTimestamp: time.Now(),
			UpdateTimestamp: time.Now(),
			CreateUserId:    -1,
			UpdateUserId:    -1,
			Version:         0,
		},
		TableId: 1,
		Name:    "Name 1",
	}, tableName)

	return err
}

func insertTableTestInsertTxWithSeq(trx pgx.Tx, tableName string) error {
	err := gldb.InsertTx(trx, TableTestInsertTxWithSeq{
		BaseEntityTs: glentity.BaseEntityTs{
			CreateTimestamp: time.Now(),
			UpdateTimestamp: time.Now(),
			CreateUserId:    -1,
			UpdateUserId:    -1,
			Version:         0,
		},
		Name: "Name 1",
	}, tableName)

	return err
}

func initTableTestInsertTx(trx pgx.Tx) string {
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

func initTableTestInsertTxWithTx(trx pgx.Tx) string {
	tableName := `t_test_tablet_seq_` + glutil.DateTimeNow()

	sql := `CREATE TABLE IF NOT EXISTS ` + tableName + `( ` +
		` table_id 	bigserial, ` +
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
