package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "maminactl",
		Usage:   "helper cli for mamina development. It is usefull for generating new routes with pages.",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "page",
				Aliases: []string{"p"},
				Usage:   "Generate a New Page code",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Specify the name of the page",
					},
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "url Path to the page",
					},
					&cli.StringFlag{
						Name:    "title",
						Aliases: []string{"t"},
						Usage:   "title of the page, usually in slovak language",
					},
					&cli.BoolFlag{
						Name:  "dry-run",
						Value: false,
						Usage: "command will output result into stdout instead of real files",
					},
				},
				Action: GeneratePage,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
