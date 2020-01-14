package main

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	home := os.Getenv("HOME")
	templatesDir := fmt.Sprintf("%s/.cbp-templates/", home)
	app := &cli.App{
		Name:  "cbp",
		Usage: "copy templates to clipboard",
		Commands: []*cli.Command{
			copyFunction(templatesDir),
			listFunction(templatesDir),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func copyFunction(templateDir string) *cli.Command {
	return &cli.Command{
		Name:    "cp",
		Aliases: []string{"copy", "c"},
		Usage:   "copies to clipboard",
		Action: func(context *cli.Context) error {
			template := context.Args().First()
			if template == "" {
				return errors.New("template name required")
			}
			fileName := fmt.Sprintf("%s/%s.txt", templateDir, template)
			file, err := os.Open(fileName)
			if err != nil {
				log.Println(err.Error())
				return fmt.Errorf("template not found file: %s", template)
			}
			defer func() { _ = file.Close() }()
			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println(err.Error())
				return fmt.Errorf("failed to load file: %s", template)
			}
			err = clipboard.WriteAll(string(bytes))
			if err != nil {
				log.Println(err.Error())
				return errors.New("failed to write to clipboard")
			}
			return nil
		},
	}
}

func listFunction(templateDir string) *cli.Command {
	return &cli.Command{
		Name:    "ls",
		Aliases: []string{"list", "l"},
		Usage:   "list all templates",
		Action: func(context *cli.Context) error {
			dir, err := os.Open(templateDir)
			if err != nil {
				log.Println(err.Error())
				return errors.New("not found template directory")
			}
			defer func() { _ = dir.Close() }()
			files, err := dir.Readdirnames(10000)
			if err != nil {
				log.Println(err.Error())
				return errors.New("failed to list files")
			}
			for _, file := range files {
				name := strings.Replace(file, ".txt", "", 1)
				fmt.Println(name)
			}
			return nil
		},
	}
}
