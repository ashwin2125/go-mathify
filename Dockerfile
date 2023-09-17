# Official Build Image as build environment.
FROM golang:1.21 AS builder

# Arguments that can be passed at build time.
ARG VERSION="development"
ARG BUILD_DATE=""
ARG VCS_URL=""

# Set environment variables.
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Set up the build directory.
WORKDIR /src/app

# Install the project dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy all files.
COPY . .

# Build the binary, and add the metadata as labels.
RUN go build -ldflags "-X main.Version=${VERSION} -X main.BuildDate=${BUILD_DATE} -X main.VCS_URL=${VCS_URL}" -o /go-mathify ./cmd/go-mathify

# This is a library, so no ports will be exposed.

# Final stage: Run the binary
FROM scratch AS final

LABEL org.label-schema.version=$VERSION
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.vcs-url=$VCS_URL

COPY --from=builder /go-mathify /go-mathify

ENTRYPOINT ["/go-mathify"]