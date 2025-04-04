package commands

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/browser"
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

func showPINMessage(m *Message) error {
	wd, err := GetWd()
	if err != nil {
		return err
	}

	templatesPath := filepath.Join(wd, "templates", "pin.tmpl")
	tmpl, err := template.New("pin.tmpl").ParseFiles(templatesPath)
	if err != nil {
		log.Printf("[ERROR]: could not open the template: %v", err)
		return err
	}

	dstPath := filepath.Join(os.TempDir(), "pin.html")
	file, err := os.Create(dstPath)
	if err != nil {
		log.Printf("[ERROR]: could not create the message file: %v", err)
		return err
	}
	defer os.Remove(dstPath)

	err = tmpl.Execute(file, m)
	if err != nil {
		log.Printf("[ERROR]:could not generate the message file from the template: %v", err.Error())
		return err
	}

	if err := browser.OpenFile(file.Name()); err != nil {
		log.Printf("[ERROR]:could not open the browser window to show the message file from the template: %v", err.Error())
		return err
	}

	// Give the browser some time to open and show the message (race condition)
	time.Sleep(10 * time.Second)
	return nil
}
