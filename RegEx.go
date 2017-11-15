package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func ElizaResponse(input string) string {

	if matched, _ := regexp.MatchString(`(?!).*\bfather\b.*`, input); matched {
		return "why don't you tell me more about your father?"
	}

	re := regexp.MustCompile("I am ([^.!?]*)[.!?]?")
	if re.MatchString(input) {
		return re.ReplaceAllString(input, "How do you know you are $1?")
	}
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

	fmt.Println("Input: " + "I am happy ")
	fmt.Println("Input: " + "I am not happy with your responses ")
	fmt.Println("Input: " + "I am not sure that you understand the effect that your questions are having on me ")
	fmt.Println("Input: " + "I am supposed to just take what you’re saying at face value? ")
}
