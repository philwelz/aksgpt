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
	"github.com/openai/openai-go/azure"
)

// The latest API versions, including previews, can be found here:
// https://learn.microsoft.com/en-us/azure/ai-services/openai/reference#rest-api-versionng
const azureOpenAIAPIVersion = "2024-12-01-preview"

func AzureOpenAiChat(userMessage, systemInstructions string) {
	// Check if the environment variable is set
	azureOpenAIAPIKey := os.Getenv("AZURE_OPENAI_API_KEY")
	if azureOpenAIAPIKey == "" {
		color.Red("Environment variable AZURE_OPENAI_API_KEY is not set.")
		os.Exit(1)
	}

	azureOpenAIEndpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	if azureOpenAIEndpoint == "" {
		color.Red("Environment variable AZURE_OPENAI_ENDPOINT is not set.")
		os.Exit(1)
	}

	// Create a new OpenAI client
	client := openai.NewClient(
		azure.WithEndpoint(azureOpenAIEndpoint, azureOpenAIAPIVersion),
		azure.WithAPIKey(azureOpenAIAPIKey),
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
