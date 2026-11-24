FROM golang:1.26-alpine AS builder

WORKDIR /src/backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend/ecohuerta-server ./cmd/main.go

FROM alpine:3.20

WORKDIR /app/backend

COPY --from=builder /app/backend/ecohuerta-server ./ecohuerta-server
COPY front/ /app/front/

EXPOSE 8080

CMD ["./ecohuerta-server"]