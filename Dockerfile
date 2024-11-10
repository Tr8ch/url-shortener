FROM golang:1.22-bookworm AS builder
ARG CMD=api

# Set non-root user and group to avoid root privileges
RUN addgroup --system appgroup && adduser --system --ingroup appgroup appuser

WORKDIR /app

ENV GOCACHE=/tmp/.cache/go-build

RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

WORKDIR /app/cmd/$CMD
# Disable VCS stamping with -buildvcs=false
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -buildvcs=false -o /app/main .

# Final stage: use a minimal, secure base image
FROM gcr.io/distroless/base-debian11

# Set non-root user in the runtime stage as well
USER nobody:nogroup

COPY --from=builder /app/main /app/main

WORKDIR /app

CMD ["/app/main"]
