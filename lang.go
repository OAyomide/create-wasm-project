package main

import (
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

	default:
		fmt.Println("No language passed...")
	}
}
