FROM golang:1.22.0

WORKDIR /app

COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN ls -la

RUN go build -o main .

RUN ls -la

EXPOSE 8080

CMD ["go", "run", "main.go"]