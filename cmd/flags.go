package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

var baseUrl string
var token string
var user string
var pass string

func init() {
	flag.Usage = func() {
		log.Print(usage)
		os.Exit(1)
	}

	flag.StringVar(&baseUrl, "url", "https://repo1.maven.org/maven2", "")
	flag.StringVar(&token, "token", "", "")
	flag.StringVar(&user, "user", "", "")
	flag.StringVar(&pass, "pass", "", "")
}

func parseArgs() (string, []string) {
	help := flag.Bool("help", false, "")

	flag.Parse()

	if *help {
		fmt.Print(usage)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if _, err := url.Parse(baseUrl); err != nil {
		log.Fatalf("invalid url format; %s\n", err.Error())
	}

	cmd := flag.Args()[0]
	args := flag.Args()[1:]
	return cmd, args
}
