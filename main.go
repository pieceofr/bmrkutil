package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var filepath string

const prefix = "01"

var version = "zero"

func main() {
	app := cli.NewApp()
	app.Name = "bitmark util"
	app.Version = version
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr

	app.Commands = []cli.Command{
		{
			Name:      "fingerprint",
			Usage:     "generate a fingerprint of a given file",
			ArgsUsage: "\n   (* = required)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Value: "",
					Usage: "*set the `filepath` of a file",
				},
			},
			Action: runFingerPrint,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
