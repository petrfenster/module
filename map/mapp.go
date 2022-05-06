package mapp

import "fmt"

func mapp() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#f5f5f5",
		"blue":  "#4ba345",
	}

	colours["white"] = "#ffffff"

	PrintMap(colours)
}

func PrintMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}

func PrintTest() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#f5f5f5",
		"blue":  "#4ba345",
		"jdsk":  "henene",
	}

	colours["white"] = "#ffffff"

	PrintMap(colours)
}
