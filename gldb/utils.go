package gldb

import (
	"github.com/frediansah/goleafcore/glconstant"
	"github.com/frediansah/goleafcore/glutil"
)

func QAddIfNotEmpty(added *bool, variableValue, query string) string {
	*added = false
	if len(variableValue) == 0 {
		return glconstant.EMPTY_VALUE
	}
	*added = true

	return ` ` + query
}

func QAddIfNotEquals(added *bool, variableValue1, variableValue2 interface{}, query string) string {
	*added = false
	if glutil.ToString(variableValue1) == glutil.ToString(variableValue2) {
		return glconstant.EMPTY_VALUE
	}
	*added = true

	return ` ` + query
}

func QAddIfEquals(added *bool, variableValue1, variableValue2 interface{}, query string) string {
	*added = false
	if glutil.ToString(variableValue1) != glutil.ToString(variableValue2) {
		return glconstant.EMPTY_VALUE
	}
	*added = true

	return ` ` + query
}

func QCriteriaHelperLikeBoth(variableValue string, queryColumns ...string) string {
	param := QWrapLikeBoth(variableValue)
	return QCriteriaHelperLikeBase(param, queryColumns...)
}

func QCriteriaHelperLikeLeft(variableValue string, queryColumns ...string) string {
	param := QWrapLikeLeft(variableValue)
	return QCriteriaHelperLikeBase(param, queryColumns...)
}

func QCriteriaHelperLikeRight(variableValue string, queryColumns ...string) string {
	param := QWrapLikeRight(variableValue)
	return QCriteriaHelperLikeBase(param, queryColumns...)
}

func QCriteriaHelperLikeBase(param string, queryColumns ...string) string {
	result := ""
	for _, queryColumn := range queryColumns {
		if len(result) > 0 {
			result = result + " OR "
		}

		result = result + queryColumn + " ILIKE " + param
	}

	return "(" + result + ")"
}

func QWrapLikeBoth(variableValue string) string {
	return "%" + variableValue + "%"
}

func QWrapLikeLeft(variableValue string) string {
	return "%" + variableValue
}
func QWrapLikeRight(variableValue string) string {
	return variableValue + "%"
}
