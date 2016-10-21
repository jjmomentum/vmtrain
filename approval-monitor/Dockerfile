# Start from Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
MAINTAINER Luis M. Valerio <lvaleriocasti@vmware.com>

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/vmtrain/approval-monitor

# Build the microservice inside the container.
RUN go install github.com/vmtrain/approval-monitor

# Run the microservice command by default when the container starts
ENTRYPOINT /go/bin/approval-monitor -data-manager=http://0.0.0.0:6001 -poll-interval=10
