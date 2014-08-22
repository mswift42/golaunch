package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
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

type Result struct {
	FullPath string
	Name     string
}

type Searchresult struct {
	Len     int
	results []Result
}

func (sr *Searchresult) Path(i int) string {
	return sr.results[i].FullPath
}
func (sr *Searchresult) Name(i int) string {
	return sr.results[i].Name
}

func (*Control) Quit() {
	os.Exit(0)
}
func (c *Control) Search(s string) {
	c.Searchresult.Len = 0
	search := NewSearch(s)
	results := search.results
	qml.Changed(&c.Searchresult, &c.Searchresult.Len)
	Len := search.Len
	c.Searchresult.Len = Len
	c.Searchresult.results = results
	qml.Changed(&c.Searchresult, &c.Searchresult.Len)
}

func NewSearch(s string) Searchresult {
	cmd := exec.Command("locate", "-l", "10", "-i", s)
	out, _ := cmd.Output()
	var sr Searchresult
	split := strings.Split(string(out), "\n")
	sr.results = NewResults(split)
	sr.Len = len(sr.results)
	return sr
}
func NewResults(s []string) []Result {
	length := len(s)
	results := make([]Result, length)
	for i, j := range s {
		results[i].FullPath = j
		results[i].Name = getFileFromPath(j)
	}
	return results
}
func (*Control) Select(s string) {
	err := exec.Command("xdg-open", s).Run()
	if err != nil {
		panic(err)
	}

}
func getFileFromPath(s string) string {
	_, file := path.Split(s)
	return strings.Split(file, ".")[0]
}

// func (c *Control) Len(s string) *Searchresult {
