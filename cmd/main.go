package main

import (
	"flag"
	"hub.lol/mvnc"
	"os"
)

func main() {
	cmd, args := parseArgs()

	var gid, aid, ver, file string
	if len(args) > 0 {
		gid = args[0]
	}
	if len(args) > 1 {
		aid = args[1]
	}
	if len(args) > 2 {
		ver = args[2]
	}
	if len(args) > 3 {
		file = args[3]
	}

	var err error
	switch cmd {
	case "ls":
		err = mvnc.List(gid, aid, ver)
	case "dl":
		err = mvnc.Download(gid, aid, ver, file)
	case "up":
		err = mvnc.Upload(gid, aid, ver, file)
	case "rm":
		err = mvnc.Remove(gid, aid, ver, file)
	default:
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		logE.Fatalln(err)
	}

}
