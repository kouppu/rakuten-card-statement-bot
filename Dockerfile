# Builder Container
FROM golang:1.17.7-alpine3.15 AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN addgroup -g 10001 -S myuser \
  && adduser -u 10001 -G myuser -S myuser

RUN apk update && apk add --no-cache \
  alpine-sdk \
  git

COPY . .

RUN go build -o /bin/rakuten-card-statement-bot

# Runtime Container
FROM alpine

WORKDIR /

RUN apk update && apk add --no-cache \
  chromium \
  chromium-chromedriver

COPY --from=builder /bin/rakuten-card-statement-bot /rakuten-card-statement-bot
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd

USER myuser

ENTRYPOINT ["/rakuten-card-statement-bot"]
