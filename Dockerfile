# --- Base ---
FROM golang:1.15.3-alpine as base

WORKDIR /app

# Configure build
ENV CGO_ENABLED=0 GOOS=linux

# Warm-up go mod cache
COPY go.mod go.sum ./
RUN go mod download

# Copy code into arbitrary build path
COPY . .

# Run straight from source
ENTRYPOINT [ "go", "run", "cmd/teams/main.go" ]