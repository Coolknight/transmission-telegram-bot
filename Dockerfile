FROM golang:1.21 AS build
WORKDIR /app
ADD . /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o transmission-telegram-bot .

FROM ubuntu
RUN apt update && apt install -y \
    sane \
    sane-utils \
 && rm -rf /var/lib/apt/lists/*

COPY --from=build /app /app

# Run
ENTRYPOINT ["/app/transmission-telegram-bot"]