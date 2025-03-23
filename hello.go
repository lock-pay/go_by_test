package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("World", ""))
}

const spanish = "spanish"
const french = "french"
const englishHello = "Hello, "
const spanishHello = "Hola, "
const frenchHello = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return (getPrefix(language) + name)
}

func getPrefix(
	language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello
	}
	return
}
