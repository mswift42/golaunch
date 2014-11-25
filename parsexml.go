package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type Xbel struct {
	Bookmark []Bookmark `xml:"bookmark"`
}

type Bookmark struct {
	Href string `xml:"href,attr"`
	Info []Info `xml:"info"`
}

type Info struct {
	BookApps []BookApps
}

type BookApps struct {
	BookApp []BookApp `xml:"bookmark:applications"`
}

type BookApp struct {
	Name string `xml:"name,attr"`
}

type Result struct {
	Xbel
	Bookmark []Bookmark
	BookApps []BookApps
}

func main() {
	usr, _ := user.Current()
	home := usr.HomeDir
	recently := home + "/.local/share/recently-used.xbel"
	file, err := os.Open(recently)
	if err != nil {
		fmt.Println("Error opening File: ", err)
	}
	defer file.Close()
	f, _ := ioutil.ReadAll(file)
	var b Result
	if err := xml.Unmarshal(f, &b); err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println(b.BookApps)
}
