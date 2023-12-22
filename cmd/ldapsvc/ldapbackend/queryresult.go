package ldapbackend

import (
	"reflect"
	"strings"
)

func (q *QueryResult) IsValidFieldName(fieldName string) bool {
	if q == nil {
		return false
	}
	st := reflect.TypeOf(*q)
	f := strings.ToLower(fieldName)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		t := strings.ToLower(strings.Split(field.Tag.Get("json"), ",")[0])
		if t == f {
			return true
		}
	}
	return false
}

func (q *QueryResult) GetPropertyValue(fieldName string) (r string) {
	if q == nil {
		return ""
	}
	st := reflect.TypeOf(*q)
	r = InvalidFieldName
	f := strings.ToLower(fieldName)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		t := strings.ToLower(strings.Split(field.Tag.Get("json"), ",")[0])
		if t == f {
			return reflect.ValueOf(*q).Field(i).String()
		}
	}
	return r
}
