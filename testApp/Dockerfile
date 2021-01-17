FROM golang:1.15.6-alpine

RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .


COPY . .

RUN go build

EXPOSE 8080

CMD ["go run main.go"]

