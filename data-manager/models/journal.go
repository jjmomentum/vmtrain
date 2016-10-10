// Package models has the structs that are used by the application.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package models

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
)

// Journal is a multi FIFO queue manager. A map of topics is kept in wich
// each topic implements a FIFO queue for Messages.
type Journal interface {
	// Get one message from the topic. Normally, peek is false, in which case
	// the FIFO removes the next entry off the queue. If peek is true, the
	// next Message is returned, but the Message is not removed from the queue
	// and subsequent calls to GetTopic will return the same message.
	GetTopic(topic string, peek bool) *Message

	// Get all available topics currently in the journal.
	GetTopics() *Topics

	// Append a Message onto the queue. By Append, this means add the message
	// to the head, as the queues are always FIFO.
	Append(topic string, m *Message)

	// REST wrapper for GetTopic. Peek is passed as a query parameter on the
	// HTTP GET, which defaults to false.
	RestGetTopic(w rest.ResponseWriter, r *rest.Request)

	// REST wrapper for GetTopics.
	RestGetTopics(w rest.ResponseWriter, r *rest.Request)

	// REST wrapper for Append. The body of the POST shall contain a Message.
	RestPostTopic(w rest.ResponseWriter, r *rest.Request)
}

func init() {
	log.Println("Initialized journal.")
}
