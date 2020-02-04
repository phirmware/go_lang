package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	print(colors)
}

func print(c map[string]string) {
	for key, value := range c {
		fmt.Println(value, "This is the value for ", key)
	}
}
