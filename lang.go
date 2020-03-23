package main

import (
	langs "festus/langs"
	"fmt"
)

type NewProject struct {
	lang string
	name string
}

func (proj *NewProject) Index() {
	switch proj.lang {
	case "go":
		fmt.Println("Should handle go stuff here...")
		// CreateGo(proj.lang, proj.name)
		langs.CreateGo(proj.lang, proj.name)
	case "rust":
		langs.Install_deps()
	default:
		fmt.Println("No language passed...")
	}
}
