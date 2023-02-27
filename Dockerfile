FROM golang:1.20.1-alpine3.17
ADD ./app /app
WORKDIR /app
RUN go mod init test/go-app && \
    go mod tidy && \
    go build -o main .
CMD ["/app/main"]