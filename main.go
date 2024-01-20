package main

import (
	"context"
	_ "embed"
	"log"
	"os"

	"github.com/skelouse/code-to-gpt/parser"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "split-files",
				Aliases: []string{"s"},
			},
			&cli.BoolFlag{
				Name:    "clipboard",
				Usage:   "does not work with split-files",
				Aliases: []string{"c"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return parser.Run(parser.Options{
				SplitFiles: ctx.Bool("split-files"),
				Clipboard:  ctx.Bool("clipboard"),
			})
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
