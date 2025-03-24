/*
Copyright © 2025 Philip Welz
*/
package ai

import (
	"context"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/openai/openai-go"
)

func OpenAiChat(userMessage, systemInstructions string) {
	// Check if the environment variable is set
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		color.Red("Environment variable OPENAI_API_KEY is not set.")
		os.Exit(1)
	}

	// Create a new OpenAI client
	client := openai.NewClient()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// Set Chat parameter
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
			openai.SystemMessage(systemInstructions),
		},
		Model: openai.ChatModelGPT4oMini,
		// Seed is used to generate deterministic results
		Seed: openai.Int(1),
		// Temperature is used to control the randomness of the output
		Temperature: openai.Opt(0.5),
		// MaxTokens is used to limit the length of the output
		MaxTokens: openai.Int(10000),
		// TopP is used to control the diversity of the output
		TopP: openai.Opt(0.7),
	}

	// Create a new chat completion
	chatCompletion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		color.Red("Error: %v", err)
	}

	// Print a newline for better readability
	println()
	println(chatCompletion.Choices[0].Message.Content)
}
