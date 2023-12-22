package ldapbackend

import (
	"reflect"
	"strings"
)

func (q *QueryResult) GetPropertyValue(fieldName string) (r string) {
	if q == nil {
		return ""
	}
	st := reflect.TypeOf(*q)
	r = InvalidFieldName
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		t := strings.Split(field.Tag.Get("json"), ",")[0]
		if t == fieldName {
			return reflect.ValueOf(*q).Field(i).String()
		}
	}
	return r
}
