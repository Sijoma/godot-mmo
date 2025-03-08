# Stage 1: Build the application
FROM golang:1.23 AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY cmd/main.go cmd/main.go
COPY internal internal
COPY pkg pkg

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -o godot-mmo-server cmd/main.go

# Stage 2: Create a lightweight image
FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/godot-mmo-server .

EXPOSE 8080

ENTRYPOINT["/godot-mmo-server"]
