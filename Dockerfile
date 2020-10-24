# --- Base ---
FROM golang:1.15.3-alpine as base

# Ca-certificates is required to call HTTPS endpoints.
RUN apk add --no-cache ca-certificates git coreutils

# Set build directory
ENV BUILD_DIR=/build
WORKDIR ${BUILD_DIR}

# Configure build
ENV CGO_ENABLED=0 GOOS=linux

# Warm-up go mod cache
COPY go.mod go.sum ./
RUN go mod download


# --- Source ---
FROM base as source
# Copy code into arbitrary build path
COPY . .


# --- Service ---
FROM source as service
ARG GO_MAIN_PACKAGE

# Service specific config
WORKDIR /build/${GO_MAIN_PACKAGE}


# --- Dev ---
FROM service as dev

# Run straight from source
ENTRYPOINT [ "go", "run", "." ]


# --- Build ---
FROM service as builder

# Build the Linux binary
RUN go build -o /go/bin/app


# --- Release ---
FROM scratch as release

# unprivileged user:group
USER 1001:1001

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/app /app

# Run the service binary.
ENTRYPOINT [ "/app" ]
