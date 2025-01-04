package main

import (
	"log"
	"os"

	"github.com/doncicuto/openuem_message/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "openuem-message",
		Commands:  getCommands(),
		Usage:     "Show OpenUEM messages to Windows users",
		Authors:   []*cli.Author{{Name: "Miguel Angel Alvarez Cabrerizo", Email: "mcabrerizo@openuem.eu"}},
		Copyright: "2024 - Miguel Angel Alvarez Cabrerizo <https://github.com/doncicuto>",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCommands() []*cli.Command {
	return []*cli.Command{
		commands.InfoMessage(),
	}
}
