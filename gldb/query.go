package gldb

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"gitlab.com/stsgoleaf/goleafcore/glinit"
	"gitlab.com/stsgoleaf/goleafcore/glutil"
)

const TAG_GLEAF string = "gleaf"
const TAG_VALUE_PK string = "pk"
const TAG_VALUE_SEQUENCE string = "seq"

var cacheQueryI map[string]string = make(map[string]string)
var cacheIndexI map[string][]int = make(map[string][]int)

func Insert(obj interface{}, tableName string) error {

	values := make([]interface{}, 0)
	query := queryInsertWithValues(obj, tableName, &values)
	db := glinit.GetDB()

	logrus.Debug("Query insert : ", query)
	result, err := db.Exec(glinit.DB_CTX, query, values...)
	logrus.Debug("Result insert : ", result)
	if err != nil {
		logrus.Error("Error on insert : ", err)
	}

	return err
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
		*values = append(*values, o.Field(i).Interface())
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

func getColumnNames(obj interface{}, prefix string, ignoreTagValues ...string) []string {
	if obj == nil {
		return make([]string, 0)
	}

	t := reflect.TypeOf(obj)
	o := reflect.ValueOf(obj)

	if t.Kind() != reflect.Struct {
		return make([]string, 0)
	}

	result := make([]string, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if checkForTags(field, ignoreTagValues...) {
			continue
		}

		if field.Type.Kind() != reflect.Struct {
			columnName := getColumnName(field, prefix)
			result = append(result, columnName)

		} else {

			subValue := reflect.Indirect(o).FieldByName(field.Name).Interface()
			result = append(result, getColumnNames(subValue, prefix)...)
		}

	}

	return result
}

func getColumnNamesWithValues(obj interface{}, prefix string, values *[]interface{}, ignoreTagValues ...string) []string {
	if obj == nil {
		return make([]string, 0)
	}

	t := reflect.TypeOf(obj)
	o := reflect.ValueOf(obj)

	cacheKey := genCacheKey(obj)

	if t.Kind() != reflect.Struct {
		return make([]string, 0)
	}

	result := make([]string, 0)

	ignoreKey := make([]int, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if checkForTags(field, ignoreTagValues...) {
			ignoreKey = append(ignoreKey, i)

			continue
		}

		if field.Type.Kind() != reflect.Struct {
			columnName := getColumnName(field, prefix)
			result = append(result, columnName)
			*values = append(*values, o.Field(i).Interface())
		} else {
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
