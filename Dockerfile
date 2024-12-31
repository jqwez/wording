FROM golang:1.23.4-alpine as builder

RUN apk add --no-cache \
    make \
    curl \
    build-base

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-arm64 && chmod +x tailwindcss-linux-arm64 && mv tailwindcss-linux-arm64 /bin/tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY . .

RUN make release

CMD ["make", "run"]

EXPOSE 8080