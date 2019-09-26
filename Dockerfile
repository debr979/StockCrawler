FROM golang:1.12.0-alpine3.9
RUN apk add git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o stockcrawler .
CMD ["./stockcrawler"]