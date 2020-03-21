package main

import (
	"errors"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func about() {
	app.Name = "Festus - Create new WASM project"
	app.Author = "Ayomide Onigbinde"
	app.Version = "1.0.0"
}

func createForLang(c *cli.Context) error {
	if c.NArg() > 0 {
		name := c.Args().First()
		lang, err := getlang(c)
		if err != nil {
			log.Fatal(err)
		}
		color.Magenta("🛠 Creating new WASM project %s for language: %s...", name, lang)
		var newp = NewProject{lang: lang, name: name}
		newp.Index()
	}

	return nil
}

func getlang(c *cli.Context) (string, error) {
	lang := c.String("language")
	if lang != "" {
		return lang, nil
	}
	return "", errors.New("Please pass a language")
}

func commands() {
	app.Commands = []cli.Command{
		cli.Command{
			Name:    "new",
			Aliases: []string{"n"},
			Action:  createForLang,

			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "language, l",
					Required: true,
				},
			},
		},
	}
}

func main() {
	// fmt.Println("Hello there, new version of go. its been a while")
	about()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
