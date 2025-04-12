/*
Copyright © 2025 Philip Welz
*/
package ai

import "github.com/openai/openai-go"

// DefaultChatParams defines the default parameters for chat completion
func DefaultChatCompletionParams(
	userMessage, systemInstructions string,
) openai.ChatCompletionNewParams {
	return openai.ChatCompletionNewParams{
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
}
