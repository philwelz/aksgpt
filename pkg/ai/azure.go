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

	// The latest API versions, including previews, can be found here:
	// https://learn.microsoft.com/en-us/azure/ai-services/openai/reference#rest-api-versionng
	const azureOpenAIAPIVersion = "2025-01-01-preview"
	// const azureOpenAIAPIVersion = "2024-10-21"
	// const azureOpenAIAPIVersion = "2024-06-01"

	// Create a new OpenAI client
	client := openai.NewClient(
		azure.WithEndpoint(azureOpenAIEndpoint, azureOpenAIAPIVersion),

		// Choose between authenticating using a TokenCredential or an API Key
		//azure.WithTokenCredential(tokenCredential),
		azure.WithAPIKey(azureOpenAIAPIKey),
	)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// Set Chat parameter
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
			openai.SystemMessage(systemInstructions),
		},
		Model: openai.ChatModelGPT4o,
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
		panic(err.Error())
	}

	// Print a newline for better readability
	println()
	println(chatCompletion.Choices[0].Message.Content)
}
