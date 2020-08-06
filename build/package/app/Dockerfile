# Build stage
FROM golang:alpine AS builder

RUN apk add --no-cache git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./ ./

# Build the application
RUN go build -o /go/bin/app ./cmd/server

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY ./web /web

ARG DB_PWD
ARG DB_NAME
ENV POSTGRES_PASSWORD=$DB_PWD
ENV POSTGRES_DB=$DB_NAME

ENTRYPOINT ./app
#LABEL Name=${serviceName} Version=${version}
ARG APP_PORT
EXPOSE $APP_PORT
