# Create a Docker image for main.go
# Build the image with:
#   docker build -t main .
# Run the image with:
#   docker run -p 8080:8080 main

# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/infamousjoeg/nhl-game-today
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get -d -v ./...
RUN go install -v ./...

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM gcr.io/distroless/base

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/bin/nhl-game-today /nhl-game-today

# Expose the port the app is listening on.
EXPOSE 8080

# Run the web service on container startup.
CMD ["/nhl-game-today"]

# [END dockerfile]