package commands

import (
	"os"
	"path/filepath"

	"github.com/ncruces/zenity"
	"github.com/urfave/cli/v2"
)

func InfoMessage() *cli.Command {
	return &cli.Command{
		Name:   "info",
		Usage:  "Send an OpenUEM's info message",
		Flags:  InfoFlags(),
		Action: showInfoMessage,
	}
}

func showInfoMessage(cCtx *cli.Context) error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}

	return zenity.Info(cCtx.String("message"),
		zenity.Title(cCtx.String("title")),
		zenity.InfoIcon,
		zenity.WindowIcon(filepath.Join(filepath.Dir(ex), "assets", "icon.png")))
}

func InfoFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "title",
			Usage:    "the title you want to show",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "message",
			Usage:    "the message you want to show",
			Required: true,
		},
	}
}
