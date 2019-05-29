package main

import "fmt"

func createMapFirstApproach() map[string]string {
	colors := map[string]string {
		"red": "#ff0000",
		"white": "#fff",
		"black": "#000",
	}

	return colors
}

func createMapSecondApproach() map[string]string {
	return make(map[string]string)
}

func printMap(colourMap map[string]string) {
	for key, value := range colourMap {
		fmt.Printf("%v is : %v\n", key, value)
	}
}

func main() {
	ourMap := createMapFirstApproach()
	printMap(ourMap)
}

