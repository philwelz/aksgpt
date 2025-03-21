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

func OpenAiChat(UserMessage string, SystemInstructions string) {
	// Check if the environment variable is set
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		color.Red("Environment variable OPENAI_API_KEY is not set.")
		os.Exit(1)
	}

	// Create a new OpenAI client
	client := openai.NewClient()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Set Chat parameter
	params := openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(UserMessage),
			openai.SystemMessage(SystemInstructions),
		}),
		Model: openai.F(openai.ChatModelGPT4oMini),
		// Seed is used to generate deterministic results
		Seed: openai.Int(1),
		// Temperature is used to control the randomness of the output
		Temperature: openai.F(0.5),
		// MaxTokens is used to limit the length of the output
		MaxTokens: openai.Int(10000),
		// TopP is used to control the diversity of the output
		TopP: openai.F(0.7),
	}

	// Create a new chat completion
	chatCompletion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		panic(err.Error())
	}

	// Print a newline for better readability
	println()
	println(chatCompletion.Choices[0].Message.Content)
}
