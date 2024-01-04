package main

import (
	"embed"
	"log"
)

//go:embed USAGE.txt
var fs embed.FS

var usage string

func init() {
	data, err := fs.ReadFile("USAGE.txt")
	if err != nil {
		log.Fatalln(err)
	}
	usage = string(data)
}
