FROM golang:1.22.0

WORKDIR /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN ls -la

RUN go build -o main .

RUN ls -la

EXPOSE 8080

CMD ["go", "run", "main.go"]

# デバックモード
# CMD ["dlv", "--listen=:9229", "--headless=true", "--api-version=2", "--accept-multiclient", "--continue", "--allow-non-terminal-interactive", "debug", "."]