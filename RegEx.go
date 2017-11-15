package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ElizaResponse(input string) string {

	rand.Seed(42)
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	return answers[rand.Intn(len(answers))]
}
func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Input: " + "People say I look like both my mother and father.")
	fmt.Println("Output: " + ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()

	fmt.Println("Input: " + "Father was a teacher")
	fmt.Println("Output: " + ElizaResponse("Father was a teacher"))
	fmt.Println()

	fmt.Println("Input: " + "I was my father’s favourite.")
	fmt.Println("Output: " + ElizaResponse("I was my father’s favourite."))
	fmt.Println()

	fmt.Println("Input: " + "I’m looking forward to the weekend.")
	fmt.Println("Output: " + ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()

	fmt.Println("Input: " + "My grandfather was French!")
	fmt.Println("Output: " + ElizaResponse("My grandfather was French!"))
	fmt.Println()
}
