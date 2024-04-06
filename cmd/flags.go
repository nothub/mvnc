package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

var repo string
var token string
var user string
var pass string
var verbose bool
var printHelp bool
var printVersion bool

func strFlag(p *string, short string, long string, value string) {
	flag.StringVar(p, short, value, "")
	flag.StringVar(p, long, value, "")
}

func boolFlag(p *bool, short string, long string, value bool) {
	flag.BoolVar(p, short, value, "")
	flag.BoolVar(p, long, value, "")
}

func parseArgs() (cmd string, args []string) {

	strFlag(&repo, "r", "repo", "https://repo1.maven.org/maven2")
	strFlag(&token, "t", "token", "")
	strFlag(&user, "u", "user", "")
	strFlag(&pass, "p", "pass", "")
	boolFlag(&verbose, "v", "verbose", false)
	boolFlag(&printHelp, "h", "help", false)
	boolFlag(&printVersion, "V", "version", false)

	flag.Usage = func() {
		logI.Print(usage)
		os.Exit(1)
	}

	flag.Parse()

	if printHelp {
		fmt.Print(usage)
		os.Exit(0)
	}

	if printVersion {
		fmt.Print(version())
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		flag.Usage()
	}

	if _, err := url.Parse(repo); err != nil {
		logE.Fatalf("invalid url format; %s\n", err.Error())
	}

	return flag.Args()[0], flag.Args()[1:]
}
