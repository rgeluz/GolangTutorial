package main

import "fmt"

func main() {

	//declare a map method 1
	var colors1 map[string]string
	fmt.Println(colors1)

	//declare a map method 2
	colors2 := make(map[string]string)
	fmt.Println(colors2)

	//declare a map method 3
	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors3)

	//add additonal key value pair
	colors3["white"] = "#ffffff"
	fmt.Println(colors3)

	//delete keys from existing map
	//delete(colors3, "white")
	//fmt.Println(colors3)

	//iterate through a map
	printMap(colors3)
}

func printMap(c map[string]string) {
	for color, hex := range c { //color and hex is key and value
		fmt.Println("Hex code for", color, "is", hex)
	}
}
