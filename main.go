package main

import (
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
		if c.Args().First() == " " {
			nm, err := os.Getwd()
			name = nm
			if err != nil {
				log.Fatal(err)
			}
		}
		lang := getlang(c)
		color.Magenta("🛠 Creating new WASM project %s for language: %s...", name, lang)
		var newp = NewProject{lang: lang, name: name}
		newp.Index()
	}

	return nil
}

func getlang(c *cli.Context) string {
	lang := c.String("language")
	if lang != "" {
		return lang
	}
	return ""
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
					Required: false,
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
