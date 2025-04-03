package commands

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/browser"
)

func showPINMessage(m *Message) error {
	wd, err := GetWd()
	if err != nil {
		return err
	}

	templatesPath := filepath.Join(wd, "templates", "pin.templ")
	tmpl, err := template.New("pin.templ").ParseFiles(templatesPath)
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
