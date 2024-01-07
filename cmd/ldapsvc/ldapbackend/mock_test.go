package ldapbackend

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateMockData(t *testing.T) {
	assert.Nil(t, GenerateMockData(0))

	r := GenerateMockData(1)
	assert.NotNil(t, r)
	assert.Equal(t, 1, len(*r))
}

func TestQueryMockData(t *testing.T) {
	//
	_, err := QueryMockData("cn", nil)
	assert.Error(t, err)

	r := GenerateMockData(2)
	_, err = QueryMockData("", r)
	assert.NoError(t, err)

	q, err := QueryMockData(fmt.Sprintf("cn=%s", (*r)[0].CN), r)
	assert.NoError(t, err)
	assert.NotEmpty(t, q)
}

func TestSaveResult(t *testing.T) {
	assert.Error(t, SaveResult(nil, ""))
	r := GenerateMockData(2)
	assert.Error(t, SaveResult(r, ""))

}
