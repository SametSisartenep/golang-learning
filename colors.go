package main

import "fmt"

var x = map[string]string{
	"red":   "\x1b[0;31m",
	"green": "\x1b[0;32m",
	"blue":  "\x1b[0;34m",
	"reset": "\x1b[0m",
}

func colorme(t, text string) string {
	switch t {
	case "red":
		return x["red"] + text + x["reset"]
	case "green":
		return x["green"] + text + x["reset"]
	case "blue":
		return x["blue"] + text + x["reset"]
	default:
		return "There was a error."
	}
}

func main() {
	fmt.Println(colorme("blue", "This text should be red."))
}
