package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

func main() {
	var instruction string
	flag.StringVar(&instruction, "instruction", "You are a helpful assistant.", "")
	flag.Parse()

	query := strings.TrimSpace(strings.Join(flag.Args(), " "))
	if query == "" {
		_, _ = fmt.Fprintln(os.Stderr, "usage: go run . -- \"your question\"")
		os.Exit(2)
	}

	ctx := context.Background()
	cm, err := newChatModel(ctx)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	messages := []*schema.Message{
		schema.SystemMessage(instruction),
		schema.UserMessage(query),
	}

	_, _ = fmt.Fprint(os.Stdout, "[assistant] ")
	stream, err := cm.Stream(ctx, messages)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer stream.Close()

	for {
		frame, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if frame != nil {
			_, _ = fmt.Fprint(os.Stdout, frame.Content)
		}
	}
	_, _ = fmt.Fprintln(os.Stdout)
}

func newChatModel(ctx context.Context) (model.ToolCallingChatModel, error) {
	return openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  os.Getenv("GITHUB_TOKEN"),
		Model:   os.Getenv("MODEL"),
		BaseURL: os.Getenv("BASE_URL"),
	})
}
