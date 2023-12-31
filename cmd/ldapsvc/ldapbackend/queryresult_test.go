package ldapbackend

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryResult_GetPropertyValueOK(t *testing.T) {
	d := QueryResult{
		CN:        "cn-val",
		UID:       "uid-val",
		ClientId:  "123",
		Country:   "IT",
		Groups:    nil,
		Roles:     nil,
		FirstName: "testfirst",
		LastName:  "testlast",
	}
	v := d.GetPropertyValue("cn")
	assert.Equal(t, d.CN, v)

	v = d.GetPropertyValue("invalid")
	assert.Equal(t, InvalidFieldName, v)

	var p *QueryResult
	assert.Equal(t, "", p.GetPropertyValue("cn"))
}

func TestQueryResult_IsValidFieldName(t *testing.T) {
	d := QueryResult{
		CN:        "cn-val",
		UID:       "uid-val",
		ClientId:  "123",
		Country:   "IT",
		Groups:    nil,
		Roles:     nil,
		FirstName: "testfirst",
		LastName:  "testlast",
	}
	assert.True(t, d.IsValidFieldName("cn"))

	assert.False(t, d.IsValidFieldName("invalid"))

	var p *QueryResult
	assert.False(t, p.IsValidFieldName("cn"))
}
