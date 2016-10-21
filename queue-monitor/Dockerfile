# Start from Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
MAINTAINER Devin Nance <dnance@vmware.com>

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/vmtrain/queue-monitor

# Get dependencies.
run go get github.com/Shopify/sarama

# Build the microservice inside the container.
RUN go install github.com/vmtrain/queue-monitor/cmd/monitor

# Run the microservice command by default when the container starts
#ENTRYPOINT /go/bin/queue-monitor -t .

