FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/smartcontractkit/bea-adapter
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

# Run the executable
CMD ["bea-adapter"]
