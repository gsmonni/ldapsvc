package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileExists(t *testing.T) {
	assert.False(t, FileExists(""))
	assert.False(t, FileExists("."))
}

func TestIsDir(t *testing.T) {
	assert.False(t, IsDir(""))
	assert.True(t, IsDir("."))
}

func TestSaveJson(t *testing.T) {
	assert.Error(t, SaveJson("", nil))

	var d map[string]int = map[string]int{"test": 1}
	assert.Error(t, SaveJson("", d))
}

func TestReadJson(t *testing.T) {
	assert.Error(t, ReadJson("", nil))

	var d map[string]int = map[string]int{"test": 1}
	assert.Error(t, ReadJson("", d))
}

func TestGetRootPath(t *testing.T) {
	p := GetRootPath()
	assert.NotEmpty(t, "/", p)
}
