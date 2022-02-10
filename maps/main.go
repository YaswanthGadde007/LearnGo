package main

import (
	"fmt"
)

func main() {
	// First way of declaring and intializing values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#33cc33",
	}
	fmt.Println(colors)
	// delete a key:value
	delete(colors, "red")

	fmt.Println(colors)

	// second way
	colors2 := make(map[string]string)
	fmt.Println(colors2)

	// adding values
	colors2["blue"] = "#cy467"
	fmt.Println(colors2)

	// third way
	var numbers map[int]string
	fmt.Println(numbers)

	main2()
}
