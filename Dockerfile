FROM golang:1.12-alpine as builder

RUN apk add --no-cache make curl git gcc musl-dev linux-headers

ADD . /go/src/github.com/smartcontractkit/bea-adapter
RUN cd /go/src/github.com/smartcontractkit/bea-adapter && GO111MODULE=on go build -o cl-bea

# Copy into a second stage container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/smartcontractkit/bea-adapter/cl-bea /usr/local/bin/

EXPOSE 8080
CMD ["cl-bea"]
