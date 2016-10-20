# Start from Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
MAINTAINER Luis M. Valerio <lvaleriocasti@vmware.com>

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/vmtrain/approval-monitor

# Get dependencies.
RUN go get github.com/ant0ine/go-json-rest/rest
RUN go get github.com/stretchr/graceful
RUN go get github.com/pborman/uuid

# Build the microservice inside the container.
RUN go install github.com/vmtrain/approval-monitor

# Run the microservice command by default when the container starts
ENTRYPOINT /go/bin/approval-monitor -l 6001 -t .

# Document that the service listens on port 6001
EXPOSE 6001
