FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go 

FROM alpine:3.24 AS runner

COPY --from=builder /app/main /usr/local/bin

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

CMD ["main"]