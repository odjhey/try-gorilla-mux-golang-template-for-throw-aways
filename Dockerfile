FROM golang:1.14.6-alpine3.12 as builder
RUN apk add --no-cache git cmake build-base

RUN mkdir -p /tmp/ && cd /tmp/ \
    && git clone https://github.com/ninja-build/ninja.git \
    && cd /tmp/ninja/ \
    && cmake -Bbuild-cmake -H. \
    && cmake --build build-cmake

COPY go.mod go.sum /go/src/myapis/main/
WORKDIR /go/src/myapis/main
RUN go mod download
COPY . /go/src/myapis/main

# RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o build/main myapis/main
RUN /tmp/ninja/build-cmake/ninja

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/myapis/main/build/main /usr/bin/main
EXPOSE 8080
ENTRYPOINT ["/usr/bin/main"]
