# learn-eino

Minimal project-specific notes for this repo (Chapter 1: ChatModel & Message).

- Entry (example): `main.go`
- Purpose: demonstrate a single-turn call to a `ChatModel` and streaming response handling.

## Setup:
```
go mod init learn-eino
go mod tidy
touch .env
```

### Required env variables:
```
GITHUB_TOKEN=...
MODEL=gpt-4.1-mini
BASE_URL=openai
```

## Run (example):
```
go run . "Who are you?"
```

# Reference:
https://www.cloudwego.io/docs/eino/quick_start/chapter_01_chatmodel_and_message/
