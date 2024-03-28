package api

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

var openAIKey = envOPENAIKEY()

type _Response struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// func to handle the logic for fetching the responce from the openAI end point
func (app *Config) GetTextCompletion(prompt string) (string, error) {
	//create the request data

	//parse the data with json

	//check err

	//call the openAi url
	c := openai.NewClient(openAIKey)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     "davinci-002",
		MaxTokens: 5,
		Prompt:    prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	//check for err
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return "", err
	}

	fmt.Println(resp)
	return "", nil

}
