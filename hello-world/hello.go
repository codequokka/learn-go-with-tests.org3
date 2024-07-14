package main

import "fmt"

const (
	spanish               = "Spanish"
	french                = "French"
	englishGreetingPrefix = "Hello, "
	spanishGreetingPrefix = "Hola, "
	frenchGreetingPrefix  = "Bonjour, "
)

func Hello(name string, langugae string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(langugae) + name
}

func greetingPrefix(langugae string) (prefix string) {
	switch langugae {
	case french:
		prefix = frenchGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
