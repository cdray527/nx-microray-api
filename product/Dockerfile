FROM golang:1.23-alpine AS dev

# Install GO
RUN apk add --no-cache git

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 5200

ENV GO_ENV=development
