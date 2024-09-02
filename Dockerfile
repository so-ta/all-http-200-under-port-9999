FROM golang:1.23-alpine

WORKDIR /app

COPY main.go .
COPY go.mod .

RUN go build -o main .

CMD ["./main"]
