package main

import "fmt"

func Hello(name, lang string) string {
	if lang != "en" {
		return "Only English is supported"
	}
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world", "en"))
}
