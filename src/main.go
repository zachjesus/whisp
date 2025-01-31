package main

import (
	"os"
	"log"

	"github.com/urfave/cli/v2" 
)

func main() {
    app := &cli.App{
        Name:  "whisp",
        Usage: "IRC chat",
		Commands: []*cli.Command{
			{
				Name:    "connect",
				Usage:   "Connect to an IRC server",
				Aliases: []string{"c"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "address",
						Aliases:  []string{"a"},
						Usage:    "IRC server address",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "port",
						Aliases:  []string{"p"},
						Usage:    "IRC server port",
						Value:    "6667",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "nickname",
						Aliases:  []string{"n"},
						Usage:    "Your IRC nickname",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					address := c.String("address")
					port := c.String("port")
					nickname := c.String("nickname")
					
					log.Printf("Connecting to IRC server at %s:%s with nickname %s...", address, port, nickname)
					return nil
				},
			},
		},
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
