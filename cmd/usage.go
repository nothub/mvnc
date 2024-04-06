package main

import (
	"embed"
)

//go:embed USAGE.txt
var fs embed.FS

var usage string

func init() {
	data, err := fs.ReadFile("USAGE.txt")
	if err != nil {
		logE.Fatalln(err)
	}
	usage = string(data)
}
