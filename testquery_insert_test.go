package goleafcore_test

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

	TableId int64  `json:"tableId" gleaf:"pk, seq"`
	Name    string `json:"name"`
}

func TestInsertQuery(t *testing.T) {
	glinit.InitLog()
	db := glinit.InitDb()
	defer db.Close()

	gldb.BeginTrx(func(trx pgx.Tx) error {

		tableName := initTableTestInsertTx(trx)
		inserted, err := insertTableTestInsertTx(trx, tableName)
		logrus.Debug("RESULT err INSERT TABLE NO SEQ ", err)
		if err != nil {
			return err
		}

		var findBy TableTestInsertTx
		errFindBy := gldb.FindByPkTx(trx, &findBy, tableName, inserted.TableId)
		logrus.Debug("FIND TABLE NO SEQ eer : ", errFindBy)
		if errFindBy != nil {
			return errFindBy
		}

		logrus.Debug("START TEST WITH SEQUENCE ")
		tableWithSeq := initTableTestInsertTxWithTx(trx)
		data, errSeq := insertTableTestInsertTxWithSeq(trx, tableWithSeq)
		logrus.Debug("errSeq INSERT TABLE WITH SEQ", errSeq)

		selectTestInsertTableWithSeq(trx, tableWithSeq, data.TableId)

		logrus.Debug("Insert twice ?   :")
		data2, errSeq2 := insertTableTestInsertTxWithSeq(trx, tableWithSeq)
		logrus.Debug("Insert result ?   :", data2)
		logrus.Debug("Insert err2 ?   :", errSeq2)

		return nil
	})

}

func selectTestInsertTableWithSeq(tx pgx.Tx, tableName string, tableId int64) error {
	query := `SELECT * FROM ` + tableName + ` WHERE table_id = $1 `

	var result TableTestInsertTxWithSeq

	err := gldb.SelectOneTx(tx, &result, query, tableId)

	jsonByte, _ := json.Marshal(result)
	logrus.Debug("Data : ", string(jsonByte))

	return err
}

func insertTableTestInsertTx(trx pgx.Tx, tableName string) (*TableTestInsertTx, error) {
	data := TableTestInsertTx{
		BaseEntityTs: glentity.BaseEntityTs{
			CreateTimestamp: time.Now(),
			UpdateTimestamp: time.Now(),
			CreateUserId:    -1,
			UpdateUserId:    -1,
			Version:         0,
		},
		TableId: 1,
		Name:    "Name 1",
	}

	err := gldb.InsertTx(trx, &data, tableName)

	return &data, err
}

func insertTableTestInsertTxWithSeq(trx pgx.Tx, tableName string) (*TableTestInsertTxWithSeq, error) {
	data := TableTestInsertTxWithSeq{
		BaseEntityTs: glentity.BaseEntityTs{
			CreateTimestamp: time.Now(),
			UpdateTimestamp: time.Now(),
			CreateUserId:    -1,
			UpdateUserId:    -1,
			Version:         0,
		},
		Name: "Name 1",
	}

	err := gldb.InsertTx(trx, &data, tableName)

	jsonByte, _ := json.Marshal(data)

	logrus.Debug("Data inserted with seq : ", string(jsonByte))

	return &data, err
}

func initTableTestInsertTx(trx pgx.Tx) string {
	tableName := `t_test_tabletx_` + glutil.DateTimeNow()

	sql := `DROP TABLE IF EXISTS ` + tableName + `; ` +
		`CREATE TABLE IF NOT EXISTS ` + tableName + `( ` +
		` table_id 	bigint, ` +
		` name 		text, ` +
		` create_timestamp 		timestamptz, ` +
		` update_timestamp 		timestamptz, ` +
		` create_user_id 		bigint, ` +
		` update_user_id 		bigint, ` +
		` version 		bigint ` +
		`); `

	gldb.ExecTx(trx, sql)

	return tableName
}

func initTableTestInsertTxWithTx(trx pgx.Tx) string {
	tableName := `t_test_tablet_seq_` + glutil.DateTimeNow()

	sql := `DROP TABLE IF EXISTS ` + tableName + `; ` +
		`CREATE TABLE IF NOT EXISTS ` + tableName + `( ` +
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
