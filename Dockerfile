FROM golang:1.20.1-alpine3.17
ADD ./app /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
