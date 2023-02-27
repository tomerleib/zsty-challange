FROM golang:1.20.1-alpine3.17
RUN mkdir /app
WORKDIR /app
ADD ./app/main.go /app
RUN go mod init test/go-app && \
    go mod tidy && \
    go build -o main .
CMD ["/app/main"]