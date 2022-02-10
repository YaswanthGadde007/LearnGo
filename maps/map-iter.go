package main

import "fmt"

func main2() {
	// First way of declaring and intializing values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#33cc33",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for ", color, "is", hex)
	}
}
