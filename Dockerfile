FROM golang:1.24-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install swag for documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/api/main.go

RUN go build -o main cmd/api/main.go

EXPOSE 8084

CMD ["./main"]
