FROM golang:1.21 AS build
WORKDIR /app
ADD . /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o transmission-telegram-bot .

FROM alpine
RUN apk update && apk add sane sane-utils
COPY --from=build /app /app

# Run
ENTRYPOINT ["/app/transmission-telegram-bot"]