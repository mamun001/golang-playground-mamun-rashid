package main

import (
    "context"
    "fmt"
    "log"
    "os"

    openai "github.com/sashabaranov/go-openai"
)

func main() {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        log.Fatal("Please set your OpenAI API key in the environment variable OPENAI_API_KEY")
    }

    client := openai.NewClient(apiKey)

    request := openai.CompletionRequest{
        Model:     "text-davinci-003",
        Prompt:    "Hello, how are you?",
        MaxTokens: 5,
    }

    response, err := client.CreateCompletion(context.Background(), request)
    if err != nil {
        log.Fatalf("Completion error: %v", err)
    }

    fmt.Println("Response:", response.Choices[0].Text)
    fmt.Printf("Tokens Used: %+v\n", response.Usage)
}

