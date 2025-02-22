FROM golang:1.23-alpine
WORKDIR /app

RUN apk add --no-cache make gcc g++ curl

RUN go install github.com/air-verse/air@latest

COPY dev/scripts/entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]
