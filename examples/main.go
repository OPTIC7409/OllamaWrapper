package main

import (
	"fmt"
	"log"

	ollama "github.com/OPTIC7409/OllamaWrapper"
)

func main() {
	// Create a new Ollama client
	client := ollama.NewClient("llama3.2", "http://localhost:11434/api/generate", true)

	// Define a custom prompt
	userPrompt := `
	You are a fast-food assistant at a drive-thru. Keep responses brief and to the point. 
	Confirm the items clearly and maintain an efficient tone.`

	// Get input from the customer
	input := "I'd like a Big Mac with fries and a Coke."

	// Process the AI response
	response := client.ProcessAIResponse(input, userPrompt)
	if response != "" {
		fmt.Println(response)
	} else {
		log.Println("Error processing AI response")
	}
}
