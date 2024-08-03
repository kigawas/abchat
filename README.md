# abchat

Realtime chat application. See [design](./design) folder for more details.

## Features

- One-to-one or group chat, unified by "conversation"
- Send messages by HTTP request, receive messages by WebSocket
- Redis as message queue for horizontal scalability
- Clean architecture following [clean-fiber](https://github.com/kigawas/clean-fiber)

## Run

### Start server

```bash
cp .env.example .env
go run main.go

# or with air for auto reload
# go install github.com/air-verse/air@latest
air

# or build for production
go build
./abchat
```

### Run tests

```bash
go test -v ./tests/*
```

### Test chat

Edit and run scripts in [scripts](./scripts) folder to test chat functionality.

1. Create users
2. Create conversations
3. Connect via WebSocket
4. Send message
5. Get messages
6. Delete message
