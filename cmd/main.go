package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
}

func main() {

	var cmd string
	var args []string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
		args = os.Args[2:]
	}

	switch cmd {
	case "list":
		list(args[0])
	case "download":
		download(args[0])
	case "upload":
		upload(args[0])
	case "remove":
		remove(args[0])
	case "help":
		fmt.Print(usage)
	default:
		log.Print(usage)
		os.Exit(1)
	}

}

func readGav(args []string) (gid string, aid string, ver string) {
	if len(args) > 0 {
		gid = args[0]
	}
	if len(args) > 1 {
		aid = args[1]
	}
	if len(args) > 2 {
		ver = args[2]
	}
	return gid, aid, ver
}

func list(arg string) {
	/*
		{
			Name:      "list",
			Aliases:   []string{"ls"},
			Usage:     "Show artifact ids, versions or artifacts",
			ArgsUsage: "<gid> [<aid>] [<ver>]",
			Action: func(ctx *cli.Context) error {
				gid, aid, ver := readGav(ctx.Args().Slice())
				if strings.TrimSpace(gid) == "" {
					// TODO: print usage
				}
				log.Printf("%q %q %q\n", gid, aid, ver)
				// TODO
				return nil
			},
		},
	*/
}

func download(arg string) {
	/*
		{
			Name:      "download",
			Aliases:   []string{"dl"},
			Usage:     "Fetch artifact(s) from remote server",
			ArgsUsage: "<gid> <aid> <ver> [<file>]",
			Action: func(ctx *cli.Context) error {
				gid, aid, ver := readGav(ctx.Args().Slice())
				var file string
				if ctx.NArg() > 3 {
					file = ctx.Args().Get(3)
				}
				log.Printf("%q %q %q %q\n", gid, aid, ver, file)
				// TODO
				return nil
			},
		},
	*/
}

func upload(arg string) {
	/*
		{
			Name:      "upload",
			Aliases:   []string{"up"},
			Usage:     "Upload artifact to remote server",
			ArgsUsage: "<gid> <aid> <ver> <file>",
			Action: func(ctx *cli.Context) error {
				gid, aid, ver := readGav(ctx.Args().Slice())
				var file string
				if ctx.NArg() > 3 {
					file = ctx.Args().Get(3)
				}
				log.Printf("%q %q %q %q\n", gid, aid, ver, file)
				// TODO
				return nil
			},
		},
	*/
}

func remove(arg string) {
	/*
		{
			Name:      "remove",
			Aliases:   []string{"rm"},
			Usage:     "Remove artifact from remote server",
			ArgsUsage: "<gid> <aid> <ver> <file>",
			Action: func(ctx *cli.Context) error {
				gid, aid, ver := readGav(ctx.Args().Slice())
				var file string
				if ctx.NArg() > 3 {
					file = ctx.Args().Get(3)
				}
				log.Printf("%q %q %q %q\n", gid, aid, ver, file)
				// TODO
				return nil
			},
		},
	*/
}
