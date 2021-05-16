package goleafcore

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/frediansah/goleafcore/gldb"
	"github.com/frediansah/goleafcore/glinit"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/sirupsen/logrus"
)

var tableNameTestLockerWithDb = `t_test_locker_with_db_seq_` + glutil.DateTimeNow()

type EntityTestLockerWithDb struct {
	TableId int64  `json:"tableId" gleaf:"pk seq"`
	Name    string `json:"name"`
	Version int64  `json:"version"`
}

func (e EntityTestLockerWithDb) TableName() string {
	return tableNameTestLockerWithDb
}

func TestLockerWithDb(t *testing.T) {
	glinit.InitLog()
	db := glinit.InitDb()
	defer db.Close()

	logrus.Debug("DB POOL max conn : ", db.Config().Copy().MaxConns)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var wg sync.WaitGroup

	n := 1000
	m := 10
	tableName := initTableTestLocker()
	initDataTestLocker(tableName, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			wg.Add(1)
			go testLockerFetchDbAndUpdate(i, r1, &wg)
		}
	}

	wg.Wait()
	logrus.Debug("DONE all thread")

	dropTableTestLocker(tableName)

}

func initTableTestLocker() string {
	tableName := tableNameTestLockerWithDb

	sql := `DROP TABLE IF EXISTS ` + tableName + `; ` +
		`CREATE TABLE IF NOT EXISTS ` + tableName + `( ` +
		` table_id 	bigserial, ` +
		` name 	text, ` +
		` version 		bigint ` +
		`)`
	gldb.Exec(sql)

	return tableName
}

func dropTableTestLocker(tableName string) error {
	logrus.Debug("DROP TABLE : ", tableName)
	sql := `DROP TABLE IF EXISTS ` + tableName + `; `
	return gldb.Exec(sql)
}

func initDataTestLocker(tableName string, n int) string {
	values := ""
	for i := 0; i < n; i++ {
		if len(values) > 0 {
			values = values + ", "
		}
		values = values + "(" + strconv.Itoa(i+1) + ", 'name" + strconv.Itoa(i+1) + "', 0)"
	}

	sql := `INSERT INTO ` + tableName + ` ( table_id, name, version )` +
		`VALUES ` + values
	err := gldb.Exec(sql)
	if err != nil {
		logrus.Error("error insert data :", err.Error())
	}

	return tableName
}

func testLockerFetchDbAndUpdate(i int, random *rand.Rand, wg *sync.WaitGroup) {
	defer wg.Done()
	var max = 1000
	var min = 1

	sleepTime := random.Intn(max-min) + min
	logrus.Debug("Sleep ", sleepTime, " mili seconds")
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	logrus.Debug("Done sleep ", sleepTime, " mili second")

	pk := int64(i + 1)
	pkStr := glutil.ToString(pk)
	glutil.Lock.Get("dtpk" + pkStr).Lock()
	defer glutil.Lock.Get("dtpk" + pkStr).Unlock()

	var data EntityTestLockerWithDb
	err := gldb.FindByPk(&data, "", pk)
	if err != nil {
		logrus.Debug("err featch data: ", err)
		return
	}

	queryUpdate := ` UPDATE ` + tableNameTestLockerWithDb +
		` SET version = $1 ` +
		` WHERE table_id = $2 `

	err = gldb.Exec(queryUpdate, data.Version+1, pk)
	if err != nil {
		logrus.Debug("err update data: ", err)
		return
	}
}
