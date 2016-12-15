package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/yudppp/gosplit"
)

var lineCount int
var headerLineCount int

func main() {
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "To show help for the tool",
	}
	cli.AppHelpTemplate = HelpTemplate()
	app := cli.NewApp()
	app.Name = "gosplit"
	app.Usage = "file split tool"
	app.ArgsUsage = "[infile [outfile-prefix]]"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "lines, l",
			Value:       100,
			Usage:       "line count",
			Destination: &lineCount,
		},
		cli.IntFlag{
			Name:        "header, h",
			Value:       0,
			Usage:       "header line count",
			Destination: &headerLineCount,
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			return fmt.Errorf("infile not set\nplease show gosplit help")
		}
		args := c.Args()
		opts := gosplit.Options{
			Infile:          args.Get(0),
			OutPrefix:       args.Get(1),
			LineCount:       lineCount,
			HeaderLineCount: headerLineCount,
		}
		return gosplit.Split(opts)
	}

	app.Run(os.Args)
}

func HelpTemplate() string {
	return `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [options] {{.ArgsUsage}}
   {{if len .Authors}}
AUTHOR(S):
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
`
}
