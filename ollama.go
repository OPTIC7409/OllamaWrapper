package ollama

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	Model                   string
	APIURL                  string
	KeepConversationHistory bool
	conversationHistory     string
}

func NewClient(model, apiURL string, keepConversationHistory bool) *Client {
	return &Client{
		Model:                   model,
		APIURL:                  apiURL,
		KeepConversationHistory: keepConversationHistory,
	}
}

func (c *Client) ProcessAIResponse(input, userPrompt string) string {
	if c.KeepConversationHistory {
		c.conversationHistory += fmt.Sprintf("Input: %s\n", input)
	}

	prompt := fmt.Sprintf(`
	%s

	Conversation so far:
	%s
	Input: %s

	Respond briefly and clearly.`, userPrompt, c.conversationHistory, input)

	requestBody, err := json.Marshal(map[string]string{
		"model":  c.Model,
		"prompt": prompt,
	})
	if err != nil {
		log.Printf("Error creating AI request: %v", err)
		return ""
	}

	resp, err := http.Post(c.APIURL, "application/json", strings.NewReader(string(requestBody)))
	if err != nil {
		log.Printf("Error sending request to AI service: %v", err)
		return ""
	}
	defer resp.Body.Close()

	var responseParts []string
	decoder := json.NewDecoder(resp.Body)

	for {
		var responseChunk map[string]interface{}
		if err := decoder.Decode(&responseChunk); err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Printf("Error decoding response: %v", err)
			return ""
		}

		if response, ok := responseChunk["response"].(string); ok {
			responseParts = append(responseParts, response)
		}
	}

	finalResponse := strings.Join(responseParts, "")
	if c.KeepConversationHistory {
		c.conversationHistory += fmt.Sprintf("AI: %s\n", finalResponse)
	}

	return finalResponse
}
