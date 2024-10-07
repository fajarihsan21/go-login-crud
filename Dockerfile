FROM golang:1.20-alpine3.17 AS builder
LABEL stage=builder
# RUN apk add --no-cache gcc libc-dev tzdata

RUN apk update && apk add --no-cache gcc libc-dev tcptraceroute busybox-extras curl
ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ /etc/timezone

# Set the Current Working Directory inside the container
WORKDIR /go-login-crud/

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
# COPY go.sum .

RUN go mod download

COPY . .

RUN rm -rf pkg
RUN rm -rf .git

COPY main.go .

USER 0:0
# Build the Go app
WORKDIR /go-login-crud
RUN go build -ldflags '-linkmode=external' main.go

RUN go build -o main .
# RUN chmod -R 777 /

#second stage

FROM alpine:3.17 AS final

# Install ca-certificates and libc6-compat for Go programs to work properly
RUN apk add --no-cache ca-certificates libc6-compat
RUN apk add --update tzdata
ENV TZ=Asia/Jakarta

COPY --from=builder /go-login-crud .

# RUN chmod -R 777 /

# Run the binary program produced by `go install`
CMD ["./main"]
EXPOSE 8081