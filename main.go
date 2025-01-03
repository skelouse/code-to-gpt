package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/skelouse/code-to-gpt/parser"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "clipboard",
				Usage:   "does not work with split-files",
				Aliases: []string{"c"},
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Usage:   "silences success and tree printout",
				Aliases: []string{"q"},
			},
			&cli.StringSliceFlag{
				Name:    "include",
				Usage:   "Include files matching glob patterns",
				Aliases: []string{"i"},
			},
			&cli.StringSliceFlag{
				Name:    "exclude",
				Usage:   "Exclude files matching glob patterns",
				Aliases: []string{"e"},
			},
			&cli.IntFlag{
				Name:  "max-size",
				Usage: "Maximum total size of files to process (in bytes)",
				Value: 2 * 1024 * 1024, // Default to 2MB
			},
			&cli.BoolFlag{
				Name:    "split-files",
				Aliases: []string{"s"},
			},
		},
		Action: func(ctx *cli.Context) error {
			// Check for unexpected positional arguments
			if ctx.Args().Len() > 0 {
				if len(ctx.StringSlice("include")) > 0 || len(ctx.StringSlice("exclude")) > 0 {
					return fmt.Errorf(
						"unexpected arguments: %v\nThis may be due to shell glob expansion.\nPlease quote your glob patterns to prevent shell expansion.\nExample: --include \"**/*.js\"",
						ctx.Args().Slice())
				}

				return fmt.Errorf("arguments not supported: %s", ctx.Args().Slice())
			}

			return parser.Run(parser.Options{
				SplitFiles: ctx.Bool("split-files"),
				Clipboard:  ctx.Bool("clipboard"),
				Include:    ctx.StringSlice("include"),
				Exclude:    ctx.StringSlice("exclude"),
				Quiet:      ctx.Bool("quiet"),
				MaxSize:    ctx.Int("max-size"),
			})
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
