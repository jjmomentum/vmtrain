# q3-training-journal
Sample microservice to manage multiple topic queues (a "cheap man's journal).
The queues are FIFO.

# Getting Started
Until such time as we release this repository to the public, clone it in
your GOPATH at src/github.com/tdhite/q3-training-journal, then build it:

For example:

    export GOPATH=${HOME}/go
    cd ${HOME}/go
    mkdir -p $GOPATH/src/github.com/tdhite
    cd $GOPATH/src/github.com/tdhite
    git clone http://gerrit.eng.vmware.com:8080/q3-training-journal
    cd q3-training-journal
    make

Note for Docker for Mac users:

If you get an error like: `standard_init_linux.go:175: exec user process caused "exec format error"` do the build like this:
(See issue: https://github.com/docker/docker/issues/23865)

    docker run -it -v "$PWD":/go/src/github.com/tdhite/q3-training-journal -w /go/src/github.com/tdhite/q3-training-journal golang:1.6 make q3-training-journal
    docker build -t q3-training-journal --rm=true .

## The API
Run the service, e.g.:

    docker run -d -p 8080:8080 q3-training-journal /q3-training-journal -l 8080 -t . -apiHost localhost

You should then see the container running when you do a ```docker ps```.  If not, run a ```docker ps -a``` and ```docker logs <container-id>``` to troubleshoot.

Assuming that provided you a Docker generated container address as
172.17.0.1, the REST API exists at http://172.17.0.1:8080/api/topics and paths
further thereuafter pursuant to the pattern:

- GET /api/topics:
Returns all topics currently held by the microservice.

- GET /api/topic/:topic :
Returns the next message off the named queue, where :topic is the name.
Note the message should be base64 encoded for any nontrivial payload.

- POST /api/topic/:topic :
Given a JSON body similar to:
Note the message should be base64 encoded for any nontrivial payload.

```
    {
      "id": 0,
      "message": "dGVzdGluZw0K"
    }
```

creates a new queue message on the queue name by the topic: path parameter.

## The HTML Interface
To reach the HTML interface (given the same sample as above), browse to:
http://172.17.0.1/html/tmpl/index and the bulk  of the HTML paths are
available from that page or others as appropriate given traversal of the 'site.'

Another HTML page not linked by the inde page is /html/tmpl/hits, which provides
a view of hit counts on the various URLs involved in the service (API and HTML).

# Dependencies
This service requires a valid Go language environment and gnu make.

# License and Author
Copyright: Copyright (c) 2015 VMware, Inc. All Rights Reserved

Author: Tom Hite, VMware, Inc.

License: MIT

For details of the license, see the LICENSE file.
