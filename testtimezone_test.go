package goleafcore

import (
	"encoding/json"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/frediansah/goleafcore/gldb"
	"github.com/frediansah/goleafcore/glinit"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/sirupsen/logrus"
)

type TestTable struct {
	TableId         int64     `json:"id" gleaf:"pk"`
	Name            string    `json:"name"`
	CreateTimestamp time.Time `json:"createTimestamp"`
	UpdateTimestamp time.Time `json:"updateTimestamp"`
	CreateUserId    int64     `json:"createUserId"`
	UpdateUserId    int64     `json:"updateUserId"`
	Version         int64     `json:"version"`
}

func TestInsertUsingTimezone(t *testing.T) {
	glinit.InitLog()
	glinit.InitDb()

	tableName := initTable()
	defer dropTable(tableName)

	insertData(tableName, 10)

	resultList := selectData(tableName)

	result := struct {
		DataList []*TestTable `json:"dataList"`
	}{
		DataList: resultList,
	}

	json, err := json.Marshal(&result)
	if err != nil {
		logrus.Error("Error json : ", err)

	}
	log.Println("Result select : ", string(json))

}

func selectData(tableName string) []*TestTable {
	var resultList []*TestTable

	query := `SELECT * FROM ` + tableName
	err := gldb.Select(&gldb.ReturnSelect{
		Result: &resultList,
	}, query)
	if err != nil {
		logrus.Error("Error select with query ", query, " error ", err)
	}

	return resultList
}
func insertData(tableName string, count int) {
	for i := 0; i < count; i++ {
		name := `Data ` + strconv.Itoa(i+1)
		now := time.Now()
		data := TestTable{
			TableId:         int64(i + 1),
			Name:            name,
			CreateTimestamp: now,
			UpdateTimestamp: now,
			CreateUserId:    -1,
			UpdateUserId:    -1,
			Version:         0,
		}

		err := gldb.Insert(data, tableName)
		if err != nil {
			logrus.Error("Error insert ", err)
		}
	}
}

func initTable() string {
	tableName := `TestTable` + glutil.DateTimeNow()

	sql := `CREATE TABLE ` + tableName + `( ` +
		` table_id 	bigint, ` +
		` name 		text, ` +
		` create_timestamp 		timestamptz, ` +
		` update_timestamp 		timestamptz, ` +
		` create_user_id 		bigint, ` +
		` update_user_id 		bigint, ` +
		` version 		bigint ` +
		`)`

	gldb.Exec(sql)

	return tableName
}

func dropTable(tableName string) error {
	sql := `DROP TABLE ` + tableName

	return gldb.Exec(sql)
}
