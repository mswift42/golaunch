package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/qml.v0"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	qml.Init(nil)
	engine := qml.NewEngine()
	var ctrl Control
	engine.Context().SetVar("ctrl", &ctrl)
	component, err := engine.LoadFile("golaunch.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
	window.Show()
	window.Wait()
	return nil
}

type Control struct {
	Searchresult Searchresult
}

type Searchresult struct {
	Len     int
	results []string
}

func (sr *Searchresult) Text(i int) string {
	return sr.results[i]
}

func (*Control) Quit() {
	os.Exit(0)
}
func (c *Control) Search(s string) {
	res := NewSearch(s).results
	c.Searchresult.Len = 0
	qml.Changed(&c.Searchresult, &c.Searchresult.Len)
	c.Searchresult.Len = len(res)
	c.Searchresult.results = res
	qml.Changed(&c.Searchresult, &c.Searchresult.Len)
}

func NewSearch(s string) Searchresult {
	cmd := exec.Command("locate", s)
	out, _ := cmd.Output()
	var sr Searchresult
	sr.Len = len(out)
	sr.results = strings.Split(string(out), "\n")
	return sr
}

// func (c *Control) Len(s string) *Searchresult {
