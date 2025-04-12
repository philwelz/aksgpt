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
	"github.com/openai/openai-go/option"
)

const githubModelsAPIURL = "https://models.inference.ai.azure.com"

func GitHubModelsChat(userMessage, systemInstructions string) {
	// Check if the environment variable is set
	apiKey := os.Getenv("GITHUB_TOKEN")
	if apiKey == "" {
		color.Red("Environment variable GITHUB_TOKEN is not set.")
		os.Exit(1)
	}
	// Create a new OpenAI client
	client := openai.NewClient(
		option.WithBaseURL(githubModelsAPIURL),
		option.WithAPIKey(apiKey),
	)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// Set Chat parameter
	params := DefaultChatCompletionParams(userMessage, systemInstructions)

	// Create a new chat completion
	chatCompletion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	// Check if we got a valid response with at least one choice
	if len(chatCompletion.Choices) == 0 {
		color.Red("Error: No choices returned in the response")
		return
	}

	// Print a newline for better readability
	println()
	println(chatCompletion.Choices[0].Message.Content)
}
