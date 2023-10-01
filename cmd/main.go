package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)

	app := cli.NewApp()

	app.Name = "mvnc"
	app.Usage = "A client for Apache Maven HTTP repositories."

	app.UseShortOptionHandling = true
	app.Suggest = true

	app.Flags = []cli.Flag{}

	app.Commands = []*cli.Command{
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
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err.Error())
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
