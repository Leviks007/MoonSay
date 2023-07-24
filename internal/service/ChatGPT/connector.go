package ChatGPT

import (
	"context"
	"errors"
	"fmt"
	gpt3 "github.com/sashabaranov/go-openai"
	"io"
	"log"
	"strings"
)

func GetResponse(client *gpt3.Client, ctx context.Context, quesiton string) string {
	req := gpt3.CompletionRequest{
		Model:     gpt3.GPT3TextDavinci001,
		MaxTokens: 300,
		Prompt:    quesiton,
		Stream:    true,
	}

	resp, err := client.CreateCompletionStream(ctx, req)
	if err != nil {
		fmt.Errorf("CreateCompletionStream returned error: %v", err)
	}
	defer resp.Close()
	answerGpt := ""
	counter := 0
	for {
		data, err := resp.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Errorf("Stream error: %v", err)
		} else {
			counter++
			answerGpt += data.Choices[0].Text
			fmt.Print(data.Choices[0].Text)

		}
	}
	if counter == 0 {
		fmt.Errorf("Stream did not return any responses")
	}

	fmt.Println("")
	return answerGpt
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func GetanswerGpt(question string) string {
	log.SetOutput(new(NullWriter))
	apiKey := "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"

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	answerGPT := GetResponse(client, ctx, question)
	return answerGPT
}

func validateQuestion(question string) string {
	quest := strings.Trim(question, " ")
	keywords := []string{"", "loop", "break", "continue", "cls", "exit", "block"}
	for _, x := range keywords {
		if quest == x {
			return ""
		}
	}
	return quest
}
