package commands

import (
	"html/template"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

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
		log.Printf("[ERROR]: could not open the template, reason: %v", err)
		return err
	}

	u, err := user.Current()
	if err != nil {
		log.Printf("[ERROR]: could not get current user, reason: %v", err)
		return err
	}

	openuemDir := filepath.Join(u.HomeDir, ".openuem")
	if err := os.MkdirAll(openuemDir, 0770); err != nil {
		log.Printf("[ERROR]: could not create openuem dir for current user, reason: %v", err)
		return err
	}

	dstPath := filepath.Join(openuemDir, "pin.html")
	file, err := os.Create(dstPath)
	if err != nil {
		log.Printf("[ERROR]: could not create the message file: %v", err)
		return err
	}

	err = tmpl.Execute(file, m)
	if err != nil {
		log.Printf("[ERROR]:could not generate the message file from the template: %v", err.Error())
		return err
	}

	xdgOpenPath := locateXdgOpen()

	if xdgOpenPath != "" {
		if err := exec.Command(xdgOpenPath, file.Name()).Run(); err != nil {
			log.Printf("[ERROR]:could not open the browser window to show the message file from the template using xdg-open: %v", err.Error())
		}
	} else {
		if err := browser.OpenFile(file.Name()); err != nil {
			log.Printf("[ERROR]:could not open the browser window to show the message file from the template using browser utils: %v", err.Error())
			return err
		}
	}
	return nil
}

func locateXdgOpen() string {
	xdgOpenPath := ""
	if err := filepath.WalkDir("/var/lib/flatpak", func(path string, d fs.DirEntry, err error) error {
		if err == nil && d.Name() == "xdg-open" {
			xdgOpenPath = path
		}
		return nil
	}); err != nil {
		log.Printf("[ERROR]: could not locate xdg-open, reason: %s\n", err)
	}

	return xdgOpenPath
}
