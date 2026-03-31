# learn-eino

- Entry: `main.go`

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
go run .
```

# Reference:
- https://www.cloudwego.io/docs/eino/quick_start/chapter_01_chatmodel_and_message/
- https://github.com/cloudwego/eino-examples/blob/main/quickstart/chatwitheino/docs/ch02_chatmodel_agent_runner_console.md
