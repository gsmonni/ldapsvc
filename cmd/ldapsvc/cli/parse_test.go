package cli

import (
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/websvc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	assert.Error(t, Parse(nil))

	p := websvc.Parameters{}
	assert.NoError(t, Parse(&p))

	p.SaveLastConfig = true
	assert.NoError(t, Parse(&p))
}
