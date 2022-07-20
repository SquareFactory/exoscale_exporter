FROM golang:1.18.4-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

# ---
FROM docker.io/library/busybox:1.35.0-musl

ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static-muslc-amd64 /tini
RUN chmod +x /tini

RUN mkdir /app
RUN addgroup -S app && adduser -S -G app app
WORKDIR /app

COPY --from=builder /build/app .

RUN chown -R app:app .
USER app

EXPOSE 9116

ENTRYPOINT [ "/tini", "--", "./app" ]
