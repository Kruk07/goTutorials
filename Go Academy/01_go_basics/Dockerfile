FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/go
RUN go build -o /app/main .

EXPOSE 8080

CMD ["/app/main"]