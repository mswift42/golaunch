package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileFromPath(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(getFileFromPath("/home/severin/Documents/sicp.pdf"), "sicp")
}
