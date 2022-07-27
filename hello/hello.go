package main

import "fmt"

const englishHello = "Hello, "
const spanishHello = "Hola, "

func Hello(recipient, lang string) string {
	if recipient == "" {
		recipient = "World"
	}
	if lang == "es" {
		return spanishHello + recipient
	}
	return englishHello + recipient
}

func main() {
	fmt.Println(Hello("world", ""))
}
