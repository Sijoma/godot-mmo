# Stage 1: Build the application
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN go build -o godot-mmo-server cmd/main.go

# Stage 2: Create a lightweight image
FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/godot-mmo-server .

CMD ["./godot-mmo-server"]
