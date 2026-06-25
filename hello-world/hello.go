package main

import "fmt"

const (
	portuguese = "Portuguese"
	spanish    = "Spanish"
	french     = "French"

	englishGreetingPrefix    = "Hello, "
	portugueseGreetingPrefix = "Olá, "
	spanishGreetingPrefix    = "Hola, "
	frenchGreetingPrefix     = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
