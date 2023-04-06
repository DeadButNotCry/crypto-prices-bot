FROM golang:1.20

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o ./build/main ./cmd/crypto-prices-bot/main.go
CMD ["./main"]