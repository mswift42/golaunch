package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileFromPath(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(getFileFromPath("/home/severin/Documents/sicp.pdf"), "sicp")
	assert.Equal(getFileFromPath("/home/severin/Documents/Dive Into Python.pdf"),
		"Dive Into Python")
}

func TestNewResults(t *testing.T) {
	assert := assert.New(t)
	s := []string{"/home/severin/Documents/sicp.pdf", "/home/severin/Documents/Dive Into Python.pdf"}
	nr := NewResults(s)
	assert.Equal(nr[0].Name, "sicp")
	assert.Equal(nr[1].Name, "Dive Into Python")
	assert.Equal(nr[0].FullPath, "/home/severin/Documents/sicp.pdf")
}
func TestNewSearchFind(t *testing.T) {
	find, _ := NewSearchFind("Lisp")
	for _, i := range find.results {
		if !strings.Contains(i.Name, "Sportschau") {
			t.Errorf("expected lisp got nil")
		}

	}
}
