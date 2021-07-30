package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red": "#FF0000",
		"Yellow": "#FFFF00",
		"Olive": "#808000",
	}

	//colors["white"] = "#FFFFFF"

	//delete(colors, "white")

	printMap(colors)
	//fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color," is ", hex)
	}
}
