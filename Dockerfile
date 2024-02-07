FROM golang:1.22-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o go-tiger .
RUN go build -ldflags="-s -w" -o runner ./pkg/runner
RUN ls

FROM gcr.io/distroless/base

# Copy binary and config files from /build to root folder of distroless container.
COPY --from=builder ["/build/go-tiger", "/build/runner", "./"]

# Command to run when starting the container.
ENTRYPOINT ["./go-tiger"]
