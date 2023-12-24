package ldapbackend

import (
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLDAPParameters_Validate(t *testing.T) {
	var p *LDAPParameters
	assert.Error(t, p.Validate())

	p = &LDAPParameters{}
	assert.NoError(t, p.Validate())

	p.Mock = true
	assert.Error(t, p.Validate())
	p.MockItemsNum = 1
	assert.Error(t, p.Validate())
}

func TestNew(t *testing.T) {
	_, err := New(nil)
	assert.Error(t, err)

	par := LDAPParameters{}
	_, err = New(&par)
	assert.Error(t, err)

	par.Mock = true
	par.MockItemsNum = uint16(1)
	_, err = New(&par)
	assert.Error(t, err)

	d := filepath.Join(common.GetRootPath(), "data", "test", "unit")
	common.Datapath = d
	if !common.IsDir(d) {
		_ = os.MkdirAll(d, 0666)
	}
	par.MockDataFile = "test.json"
	_ = os.Remove(par.MockDataFile)
	_, err = New(&par)
	assert.NoError(t, err)
	_ = os.RemoveAll(par.MockDataFile)
}
