package gldb

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/frediansah/goleafcore/glinit"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const TAG_GLEAF string = "gleaf"
const TAG_VALUE_PK string = "pk"
const TAG_VALUE_SEQUENCE string = "seq"

var cacheQueryI map[string]string = make(map[string]string)
var cacheIndexI map[string][]int = make(map[string][]int)

type ReturnSelect struct {
	Error  error
	Result interface{}
}

type TrxFunc = func(trx pgx.Tx) error

func Insert(obj interface{}, tableName string) error {

	values := make([]interface{}, 0)
	query := queryInsertWithValues(obj, tableName, &values)
	db := glinit.GetDB()

	logrus.Debug("Query insert : ", query)
	logrus.Debug("Values size : ", len(values))

	arr := fmt.Sprintf("%v", values...)
	logrus.Debug("Isi values from insert : ", arr)

	result, err := db.Exec(glinit.DB_CTX, query, values...)
	logrus.Debug("Result insert : ", result)
	if err != nil {
		logrus.Error("Error on insert : ", err)
	}

	return err
}

func InsertTx(tx pgx.Tx, obj interface{}, tableName string) error {

	values := make([]interface{}, 0)
	query := queryInsertWithValues(obj, tableName, &values)

	logrus.Debug("Query insert : ", query)
	logrus.Debug("Values size : ", len(values))

	arr := fmt.Sprintf("%v", values...)
	logrus.Debug("Isi values from insert : ", arr)

	result, err := tx.Exec(glinit.DB_CTX, query, values...)
	logrus.Debug("Result insert : ", result)
	if err != nil {
		logrus.Error("Error on insert : ", err)
	}

	return err
}

func Select(result *ReturnSelect, query string, params ...interface{}) error {
	db := glinit.GetDB()
	logrus.Debug("Select query : ", query)

	err := pgxscan.Select(glinit.DB_CTX, db, result.Result, query, params...)
	if err != nil {
		result.Error = err
		logrus.Error("Error detail : ", err)
	}

	return err
}

func SelectOne(result interface{}, query string, params ...interface{}) error {
	db := glinit.GetDB()
	logrus.Debug("Select query : ", query)

	err := pgxscan.Select(glinit.DB_CTX, db, result, query, params...)

	return err
}

func SelectOneTx(tx pgx.Tx, result interface{}, query string, params ...interface{}) error {
	logrus.Debug("Select query : ", query)

	err := pgxscan.Select(glinit.DB_CTX, tx, result, query, params...)

	return err
}

func FindByPkTx(tx pgx.Tx, result interface{}, tableName string, pk interface{}) error {
	logrus.Debug("Find by PK tx")

	columns := GetColumnNames(result, "")
	pkColumn := FindPkColumn(result, "")

	logrus.Debug("Columns : ", columns)
	logrus.Debug("pk column : ", pkColumn)

	if len(pkColumn) <= 0 {
		return errors.New("result struct does not have tag gleaf:\"pk\"")
	}

	query := "SELECT " + AppendColumnNames(columns) + " FROM " + tableName +
		` WHERE ` + pkColumn + ` = $1 `

	logrus.Debug("Query find by pk : ", query)

	err := pgxscan.Get(glinit.DB_CTX, tx, result, query, pk)

	logrus.Debug("error query : ", err)
	logrus.Debug("Resuls query list : ", result)

	return err
}

func FindByPk(result *interface{}, tableName string, pk interface{}) error {
	logrus.Debug("Find by PK")
	db := glinit.GetDB()

	columns := GetColumnNames(result, "")
	pkColumn := FindPkColumn(result, "")

	logrus.Debug("Columns : ", columns)
	logrus.Debug("pk column : ", pkColumn)

	if len(pkColumn) <= 0 {
		return errors.New("result struct does not have tag gleaf:\"pk\"")
	}

	query := "SELECT " + AppendColumnNames(columns) + " FROM " + tableName +
		` WHERE ` + pkColumn + ` = $1 `

	logrus.Debug("Query find by pk : ", query)

	err := pgxscan.Get(glinit.DB_CTX, db, result, query, pk)

	logrus.Debug("error query : ", err)
	logrus.Debug("Resuls query list : ", result)

	return err
}

func SelectTx(tx pgx.Tx, result *ReturnSelect, query string, params ...interface{}) error {
	//db := glinit.GetDB()
	logrus.Debug("SelectTx query : ", query)

	err := pgxscan.Select(glinit.DB_CTX, tx, result.Result, query, params...)
	if err != nil {
		result.Error = err
		logrus.Error("Error detail : ", err)
	}

	return err
}

func Exec(query string, params ...interface{}) error {
	db := glinit.GetDB()
	logrus.Debug("Exec query : ", query)

	command, err := db.Exec(glinit.DB_CTX, query, params...)
	logrus.Debug("Result Command : ", command)
	if err != nil {
		logrus.Debug("Error command exec : ", err)
	}

	return err
}

func ExecTx(trx pgx.Tx, query string, params ...interface{}) error {
	//db := glinit.GetDB()
	logrus.Debug("ExecTx query : ", query)

	command, err := trx.Exec(glinit.DB_CTX, query, params...)
	logrus.Debug("Result ExecTx Command : ", command)
	if err != nil {
		logrus.Debug("Error ExecTx : ", err)
	}

	return err
}

func BeginTrx(trxFun TrxFunc) error {
	db := glinit.InitDb()

	trx, err := db.Begin(glinit.DB_CTX)
	if err != nil {
		return err
	}

	errTrx := trxFun(trx)

	if errTrx == nil {
		logrus.Debug("Transaction success, commited")
		trx.Commit(glinit.DB_CTX)
	} else {
		logrus.Debug("Transaction error, rolled back : ", errTrx)

		trx.Rollback(glinit.DB_CTX)
	}

	return errTrx
}

func queryInsertWithValues(obj interface{}, tableName string, values *[]interface{}) string {
	cacheKey := genCacheKey(obj)

	query, queryExists := cacheQueryI[cacheKey]
	ignoreIndex, ignoreIndexExists := cacheIndexI[cacheKey]

	if queryExists && ignoreIndexExists {
		logrus.Debug("Use cache : ", query)
		logrus.Debug("ignoreIdx : ", ignoreIndex)

		retriveValueExeptAt(obj, ignoreIndex, values)
		return query
	}

	columnNames := getColumnNamesWithValues(obj, "", values, TAG_VALUE_SEQUENCE)

	logrus.Debug("Length from get column name with values : ", len(*values))
	result := ""
	result = `INSERT INTO ` + tableName + ` ( ` +
		AppendColumnNames(columnNames) + ` ) ` +
		` VALUES (` + generateDolar(len(columnNames)) + `) `

	cacheQueryI[cacheKey] = result

	return result
}

func retriveValueExeptAt(obj interface{}, ignoreIndex []int, values *[]interface{}) {
	o := reflect.ValueOf(obj)

	for i := 0; i < o.NumField(); i++ {
		if existInArray(i, ignoreIndex) {
			continue
		}
		val := o.Field(i).Interface()
		*values = append(*values, &val)
	}
}

func existInArray(elem int, arr []int) bool {
	if len(arr) > 0 {
		for _, item := range arr {
			if elem == item {
				return true
			}
		}
	}

	return false
}

func genCacheKey(obj interface{}) string {
	t := reflect.TypeOf(obj)

	return t.PkgPath() + "/" + t.Name()
}

func generateDolar(length int) string {
	result := ""

	if length > 0 {
		for i := 0; i < length; i++ {
			result = result + `$` + strconv.Itoa((i + 1))
			if i+1 < length {
				result = result + ", "
			}
		}
	}

	return result
}

func AppendColumnNames(columnNames []string) string {
	result := ""

	length := len(columnNames)
	for i, item := range columnNames {
		result = result + item
		if i+1 < length {
			result = result + ", "
		}
	}

	return result
}

func FindPkColumn(obj interface{}, prefix string) string {
	var result string
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if checkForTags(field, TAG_VALUE_PK) {
			return getColumnName(field, prefix)
		}
	}

	return result
}

func GetColumnNames(obj interface{}, prefix string, ignoreTagValues ...string) []string {
	logrus.Debug("Get column names ", obj)
	usePtr := false

	if obj == nil {
		return make([]string, 0)
	}

	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		usePtr = true
	}

	logrus.Debug("Type of object ", t)

	o := reflect.ValueOf(obj)
	if usePtr {
		o = o.Elem()
	}

	logrus.Debug("Value of object ", o)

	logrus.Debug("Does not struct? ", t.Kind())
	if t.Kind() != reflect.Struct {
		return make([]string, 0)
	}

	result := make([]string, 0)

	logrus.Debug("Begin loop ")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		logrus.Debug("field ", field)
		valueItem := o.Field(i).Interface()
		logrus.Debug("valueItem ", valueItem)

		if checkForTags(field, ignoreTagValues...) {
			continue
		}

		if isStructField(valueItem) {

			subValue := o.Field(i).Interface()
			result = append(result, GetColumnNames(subValue, prefix, ignoreTagValues...)...)

		} else {
			columnName := getColumnName(field, prefix)
			result = append(result, columnName)
			logrus.Debug("Append value : ", o.Field(i).Interface())
		}
	}

	logrus.Debug("DONE Get column name")

	return result
}

func isStructField(valueItem interface{}) bool {
	switch valueItem.(type) {
	case time.Time, sql.NullTime, sql.NullInt64, sql.NullString, sql.NullBool, sql.NullFloat64,
		sql.NullInt32:
		return false
	}

	return reflect.TypeOf(valueItem).Kind() == reflect.Struct
}

func getColumnNamesWithValues(obj interface{}, prefix string, values *[]interface{}, ignoreTagValues ...string) []string {
	if obj == nil {
		return make([]string, 0)
	}

	t := reflect.TypeOf(obj)
	o := reflect.ValueOf(obj)
	logrus.Debug("Get column name from : ", t.Name())

	cacheKey := genCacheKey(obj)

	if t.Kind() != reflect.Struct {
		return make([]string, 0)
	}

	result := make([]string, 0)

	ignoreKey := make([]int, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		valueItem := o.Field(i).Interface()

		logrus.Debug("Proses field ", field.Name, " --> ", valueItem)

		if checkForTags(field, ignoreTagValues...) {
			ignoreKey = append(ignoreKey, i)
			continue
		}

		typeTime := false
		switch valueItem.(type) {
		case time.Time, sql.NullTime:
			typeTime = true
		}

		if typeTime || field.Type.Kind() != reflect.Struct {
			columnName := getColumnName(field, prefix)
			result = append(result, columnName)

			logrus.Debug("Len before : ", len(*values))

			*values = append(*values, o.Field(i).Interface())
			logrus.Debug("Append value : ", o.Field(i).Interface())

			logrus.Debug("Len after : ", len(*values))

		} else {
			logrus.Debug("Masuk sini kan? ", t.Field(i).Name)
			logrus.Debug("Masuk sini kan? ", len(*values))
			subValue := o.Field(i).Interface()
			result = append(result, getColumnNamesWithValues(subValue, prefix, values, ignoreTagValues...)...)
		}

	}

	cacheIndexI[cacheKey] = ignoreKey

	return result
}

func checkForTags(field reflect.StructField, tagToIgnores ...string) bool {
	if len(tagToIgnores) > 0 {
		val, exist := field.Tag.Lookup(TAG_GLEAF)
		if exist {
			for _, tag := range tagToIgnores {
				logrus.Debug("Check tag ", val, " contains ", tag)
				if strings.Contains(val, tag) {
					return true
				}
			}

		}
	}
	return false
}

func getColumnName(field reflect.StructField, prefix string) string {
	fromTag, dbFilled := field.Tag.Lookup("db")
	if !dbFilled {
		fromTag = glutil.ToUnderedScore(field.Name)
	}

	if len(prefix) > 0 {
		return prefix + "." + fromTag
	} else {
		return fromTag
	}
}
