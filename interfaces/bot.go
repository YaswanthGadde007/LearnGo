package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreetings() string
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)

}

func (englishBot) getGreetings() string {
	return "Hello"
}

func (spanishBot) getGreetings() string {
	return "Hola"
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}
