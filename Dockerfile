FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /build
ADD go.mod .
ADD go.sum .
ADD config.yml .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/client cmd/client/main.go

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ Europe/Moscow
WORKDIR /app
COPY --from=builder /app/client /app/client
COPY --from=builder /build/config.yml /app/config.yml
CMD ["./client"]
