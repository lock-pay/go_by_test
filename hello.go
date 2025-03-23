package main

import "fmt"

func main() {
	fmt.Println(Hello("World"))
}

const Hi = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return Hi + name
}
