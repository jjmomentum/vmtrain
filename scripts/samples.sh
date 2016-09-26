#!/bin/bash
#
# Copyright (c) 2015 VMware
# Author: Tom Hite (thite@vmware.com)
#
# License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
#
HOST=localhost
PORT=8080
# Post some messages to various topics
curl -X POST http://$HOST:$PORT/api/topic/Tom -H 'Content-Type: application/json' -d "{ \"id\": 0, \"message\": \"$(echo -n 'This is a message.' | base64)\"}"
curl -X POST http://$HOST:$PORT/api/topic/Tom2 -H 'Content-Type: application/json' -d "{ \"id\": 1, \"message\": \"$(echo -n 'This is another message.' | base64)\"}"
curl -X POST http://$HOST:$PORT/api/topic/Topic2 -H 'Content-Type: application/json' -d "{ \"id\": 2, \"message\": \"$(echo -n 'This is yet another message.' | base64)\"}"
curl -X POST http://$HOST:$PORT/api/topic/Topic2 -H 'Content-Type: application/json' -d "{ \"id\": 3, \"message\": \"$(echo -n 'Are we sick of messages yet?' | base64)\"}"
curl -X POST http://$HOST:$PORT/api/topic/Tom -H 'Content-Type: application/json' -d "{ \"id\": 4, \"message\": \"$(echo -n 'This is the last darned message!' | base64)\"}"

# Get all the topics created by the posts above
curl -X GET http://$HOST:$PORT/api/topics

# Get the next message off the queues
curl -X GET http://$HOST:$PORT/api/topic/Tom
curl -X GET http://$HOST:$PORT/api/topic/Tom2
curl -X GET http://$HOST:$PORT/api/topic/Topic2
