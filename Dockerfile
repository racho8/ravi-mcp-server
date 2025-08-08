FROM golang:1.23.4

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o mcp-server

EXPOSE 8080

CMD ["./mcp-server"]