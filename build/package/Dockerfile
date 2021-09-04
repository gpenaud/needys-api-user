FROM golang:alpine AS build

# Add Maintainer Info
LABEL maintainer="guillaume.penaud@gmail.com"

RUN \
  apk add --no-cache git &&\
  mkdir /application

WORKDIR /application
ADD . /application

# Download all the dependencies
RUN go mod download

# Set build argument variables for build
ARG PROJECT
ARG RELEASE
ARG COMMIT
ARG BUILD_TIME

# Set environment variable
ENV CGO_ENABLED=0
ENV GOOS=linux

# Build the binary
RUN go build \
  -a \
  -installsuffix cgo \
  -ldflags "-s -w \
    -X main.Release=${RELEASE} \
    -X main.Commit=${COMMIT} \
    -X main.BuildTime=${BUILD_TIME}" \
    -o /needys-api-user \
  /application/cmd/needys-api-user-server/main.go

# ---------------------------------------------------------------------------- #

FROM alpine:latest

RUN adduser --system --disabled-password --home /needys-api-user needys-api-user

WORKDIR /needys-api-user
USER needys-api-user

COPY --from=build /needys-api-user .

EXPOSE 8010

ENTRYPOINT ["./needys-api-user"]
