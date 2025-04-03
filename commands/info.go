package commands

import (
	"github.com/urfave/cli/v2"
)

type Message struct {
	PIN string
}

func InfoMessage() *cli.Command {
	return &cli.Command{
		Name:   "info",
		Usage:  "Send an OpenUEM's info message",
		Flags:  InfoFlags(),
		Action: showInfoMessage,
	}
}

func showInfoMessage(cCtx *cli.Context) error {
	messageType := cCtx.String("type")
	message := Message{}

	switch messageType {
	case "pin":
		message.PIN = cCtx.String("message")
		return showPINMessage(&message)
	}

	return nil
}

func InfoFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "title",
			Usage: "the title you want to show",
		},
		&cli.StringFlag{
			Name:     "message",
			Usage:    "the message you want to show",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "type",
			Usage:    "the type of message you want to show",
			Required: true,
		},
	}
}
