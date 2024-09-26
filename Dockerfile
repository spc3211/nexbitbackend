
# Start from the latest golang base image
FROM golang:1.22.6-alpine3.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

ARG GH_PAT

ENV GO111MODULE=on
RUN apk update && apk add --no-cache git openssh-client ca-certificates tzdata && update-ca-certificates
RUN apk add build-base
RUN apk add curl
RUN apk add libxml2-dev libxslt-dev xz-dev zlib-dev


RUN git config --global url."https://${GH_PAT}@github.com/".insteadOf "https://github.com/"

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o go-api ./cmd

# Command to run the executable
ENTRYPOINT ["sh", "-c", "./go-api"]
