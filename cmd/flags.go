package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

var baseUrl string
var token string
var user string
var pass string
var verbose bool
var help bool

func parseArgs() (cmd string, args []string) {
	flag.StringVar(&baseUrl, "url", "https://repo1.maven.org/maven2", "")
	flag.StringVar(&token, "token", "", "")
	flag.StringVar(&user, "user", "", "")
	flag.StringVar(&pass, "pass", "", "")
	flag.BoolVar(&verbose, "verbose", false, "")
	flag.BoolVar(&help, "help", false, "")

	flag.Parse()

	if help {
		fmt.Print(usage)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		logI.Print(usage)
		os.Exit(1)
	}

	if _, err := url.Parse(baseUrl); err != nil {
		logE.Fatalf("invalid url format; %s\n", err.Error())
	}

	return flag.Args()[0], flag.Args()[1:]
}
