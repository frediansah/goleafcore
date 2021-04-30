package goleafcore_test

import (
	"encoding/json"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/frediansah/goleafcore/gldb"
	"github.com/frediansah/goleafcore/glentity"
	"github.com/frediansah/goleafcore/glinit"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/sirupsen/logrus"
)

type TestBaseEntityTs struct {
	CreateTimestamp time.Time `json:"createTimestamp"`
	UpdateTimestamp time.Time `json:"updateTimestamp"`
	CreateUserId    int64     `json:"createUserId"`
	UpdateUserId    int64     `json:"updateUserId"`
	Version         int64     `json:"version"`
}

type TestBaseEntity struct {
	glentity.BaseEntityTs

	TableId int64  `json:"id" gleaf:"pk"`
	Name    string `json:"name"`
}

func TestInsertUsingBaseEntity(t *testing.T) {
	glinit.InitLog()
	glinit.InitDb()

	tableName := initTableBaseEntity()
	defer dropTableBaseEntity(tableName)

	insertDataBaseEntity(tableName, 10)

	resultList := selectDataBaseEntity(tableName)

	result := struct {
		DataList []*TestBaseEntity `json:"dataList"`
	}{
		DataList: resultList,
	}

	json, err := json.Marshal(&result)
	if err != nil {
		logrus.Error("Error json : ", err)

	}
	log.Println("Result select : ", string(json))

}

func selectDataBaseEntity(tableName string) []*TestBaseEntity {
	var resultList []*TestBaseEntity

	query := `SELECT * FROM ` + tableName
	err := gldb.Select(&gldb.ReturnSelect{
		Result: &resultList,
	}, query)
	if err != nil {
		logrus.Error("Error select with query ", query, " error ", err)
	}

	return resultList
}
func insertDataBaseEntity(tableName string, count int) {
	for i := 0; i < count; i++ {
		name := `Data ` + strconv.Itoa(i+1)
		now := time.Now()
		data := TestBaseEntity{
			BaseEntityTs: glentity.BaseEntityTs{
				CreateTimestamp: now,
				UpdateTimestamp: now,
				CreateUserId:    -1,
				UpdateUserId:    -1,
				Version:         0,
			},

			TableId: int64(i + 1),
			Name:    name,
		}

		err := gldb.Insert(data, tableName)
		if err != nil {
			logrus.Error("Error insert ", err)
		}
	}
}

func initTableBaseEntity() string {
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

func dropTableBaseEntity(tableName string) error {
	sql := `DROP TABLE ` + tableName

	return gldb.Exec(sql)
}
