FROM golang:1.24 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app/app /app/
COPY --from=builder /app/migrations /migrations/
CMD [ "/app/app" ]