package gldb

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/frediansah/goleafcore/glutil"
)

type QBuilder struct {
	query       string
	paramValues map[string]interface{}
	params      []string
}

func (q *QBuilder) Add(query string) *QBuilder {
	q.query = q.query + query

	return q
}

func (q *QBuilder) AddIfNotEmpty(variableValue, query string) *QBuilder {
	if len(variableValue) > 0 {
		q.query = q.query + query
	}

	return q
}
func (q *QBuilder) AddIfEmpty(variableValue, query string) *QBuilder {
	if len(variableValue) == 0 {
		q.query = q.query + query
	}

	return q
}

func (q *QBuilder) AddIfNotEquals(variableValue1, variableValue2 interface{}, query string) *QBuilder {
	if glutil.ToString(variableValue1) != glutil.ToString(variableValue2) {
		q.query = q.query + query
	}

	return q
}

func (q *QBuilder) AddIfEquals(variableValue1, variableValue2 interface{}, query string) *QBuilder {
	if glutil.ToString(variableValue1) == glutil.ToString(variableValue2) {
		q.query = q.query + query
	}

	return q
}

func (q *QBuilder) Query() string {
	if q.params == nil {
		q.params = fetchAllParams(q.query)
	}

	return replaceAllParamsWithDollarNumber(q.query, q.params)
}

func (q *QBuilder) SetParam(name string, value interface{}) error {
	log.Println("set param :", name, "  ->", value)
	if q.params == nil {
		q.params = fetchAllParams(q.query)
	}
	if q.paramValues == nil {
		q.paramValues = map[string]interface{}{}
	}

	if !paramExists(q.params, name) {
		return errors.New("param not found: " + name)
	}

	q.paramValues[name] = value

	log.Println("params : ", q.params)
	log.Println("isi param values : ", q.paramValues)

	return nil
}

func (q *QBuilder) GetParamValues() ([]interface{}, error) {
	log.Println("param values ", q.paramValues)
	if q.params == nil {
		q.params = fetchAllParams(q.query)
	}
	if q.paramValues == nil {
		q.paramValues = map[string]interface{}{}
	}

	if len(q.params) == 0 {
		return nil, nil
	}

	var resultList []interface{}
	var errorList []string
	for _, param := range q.params {
		val, exists := q.paramValues[param]
		if !exists {
			errorList = append(errorList, param)
		}

		resultList = append(resultList, val)
	}

	if len(errorList) > 0 {
		return nil, errors.New("params.not.set: " + glutil.AppendSliceString(errorList, ", "))
	}

	return resultList, nil
}

func fetchAllParams(query string) []string {
	patternParams := "[^:]:([a-zA-Z0-9]+)"
	r, _ := regexp.Compile(patternParams)
	mapMath := r.FindAllStringSubmatch(query, -1)

	params := []string{}
	for _, group := range mapMath {
		params = append(params, group[1])
	}

	return params
}

func replaceAllParamsWithDollarNumber(query string, params []string) string {
	var dolar int = 0
	for _, param := range params {
		dolar++
		query = strings.ReplaceAll(query, ":"+param, "$"+glutil.ToString(dolar))
	}

	return query
}

func paramExists(params []string, name string) bool {
	for _, param := range params {
		if param == name {
			return true
		}
	}

	return false
}
