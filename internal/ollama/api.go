package ollama

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

func SendChatRequest(ctx context.Context, model string, think bool, oneline bool, data string) error {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}
	prompt := userPrompt
	if oneline {
		prompt = oneLineUserPrompt
	}
	req := &api.ChatRequest{
		Model: model,
		Messages: []api.Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: fmt.Sprintf(prompt, data)},
		},
		Think: &think,
	}
	return client.Chat(ctx, req, func(resp api.ChatResponse) error {
		chunk := resp.Message.Content
		fmt.Print(chunk)
		// 始终以换行符结尾
		if resp.Done {
			fmt.Println()
		}
		return nil
	})
}
