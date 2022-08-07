# app builder
FROM golang:1.18.4-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# app runner
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY scripts/* /app/
COPY config.yaml .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
