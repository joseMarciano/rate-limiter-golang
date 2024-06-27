FROM golang:1.22.4 as builder
WORKDIR /app
COPY . .
RUN env GOOS=linux GOARCH=amd64 go mod tidy && go build -o golang-app ./cmd/main.go
EXPOSE 8080
ENTRYPOINT ["./golang-app"]
