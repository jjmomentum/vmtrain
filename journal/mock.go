// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package journal

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/tdhite/q3-training-journal/app"
)

// A Topic is a single FIFO queue.
type Topic struct {
	// The cursor points to the Message that the queue will return on a
	// GetTopic call.
	cursor int

	// The array of Messages in the queue. Messages are never destroyed,
	// Rather remain in the queue as in persistent journals not having a
	// stale-out time. When backed by a system such as kafka or bookkeeper,
	// this a Topic would (should) manage a sliding window on the queue.
	messages []*Message
}

// A Mock journal, which implements a purely in-memory FIFO topic journal.
type Mock struct {
	// A lock channel for implementing semaphore blocked access to Mocks.
	lock chan int /* not used yet */

	// The map of Topics. The string index is the topic name.
	topics map[string]Topic
}

// Create and return a new Mock journal.
func NewMock() *Mock {
	mock := &Mock{
		make(chan int),
		make(map[string]Topic),
	}

	return mock
}

// [Private] Retrieve the next message on the queue and update the cursor
// in the Topic retrieved to obtain the message. Not that if the receiver
// (t) is a pointer to a temporary object, obviously the journal will not
// see the change to the cursor unless that which is pointed to by t gets
// placed into the journal, replacing an 'old' Topic. This allows peek
// operations to use this function without potentially harming the journal.
func (t *Topic) getMessage() *Message {
	msgs := len(t.messages)
	log.Printf("There are %d messages\n", msgs)
	if t.cursor < msgs {
		msg := t.messages[t.cursor]
		t.cursor += 1
		return msg
	} else {
		return nil
	}
}

// Retrieve the next message on the queue. If peek is false, the queue
// cursor updates to point at the next (FIFO) message. If peek is true the
// cursor remains unchanged.
func (m *Mock) GetTopic(topic string, peek bool) *Message {
	t, ok := m.topics[topic]
	if ok {
		msg := t.getMessage()
		/* the topic will be modified from its original state */
		if !peek {
			m.topics[topic] = t
		}
		return msg
	}
	log.Println("Topic not found for: " + topic)
	return nil
}

// Retrieve all Message entries in a queue. This is purely a convenience
// function for viewing messages, for example, in a UI for debugging
// purposes or to potentially replay a portion of a journal topic. The
// journal itself remains unmodified.
func (m *Mock) PeekTopicMessages(topic string) []Message {
	var msgs []Message

	t, ok := m.topics[topic]
	if ok {
		msgs = make([]Message, len(t.messages))
		for m := range t.messages {
			msgs[m] = *t.messages[m]
		}
		return msgs
	}
	log.Println("Topic not found for: " + topic)
	return nil
}

// Return all topics in the journal.
func (m *Mock) GetTopics() *Topics {
	keys := make([]string, len((m.topics)))
	i := 0
	for k := range m.topics {
		keys[i] = k
		i++
	}

	topics := &Topics{
		Topics: keys,
	}

	return topics
}

// Add a message to the Topic.
func (m *Mock) Append(topic string, msg *Message) {
	t, ok := m.topics[topic]
	if !ok {
		t = Topic{
			0,
			make([]*Message, 0),
		}
	}
	t.messages = append(t.messages, msg)
	m.topics[topic] = t
}

// REST handler to retrieve the current journal topic message.
func (m *Mock) RestGetTopic(w rest.ResponseWriter, r *rest.Request) {
	app.Context.Stats.AddHit(r.RequestURI)

	peekall := r.URL.Query().Get("peekall")
	topic := r.PathParam("topic")
	log.Println("checking for topic: " + topic)

	if len(peekall) > 0 {
		msgs := m.PeekTopicMessages(topic)
		if msgs == nil {
			rest.NotFound(w, r)
		} else {
			w.WriteJson(msgs)
		}
	} else {
		msg := m.GetTopic(topic, false)
		if msg == nil {
			rest.NotFound(w, r)
		} else {
			w.WriteJson(msg)
		}
	}
}

// REST handler to retrieve the current journal topic message.
func (m *Mock) RestGetTopics(w rest.ResponseWriter, r *rest.Request) {
	app.Context.Stats.AddHit(r.RequestURI)
	topics := m.GetTopics()
	w.WriteJson(topics)
}

// REST handler to retrieve the current journal topic message.
func (m *Mock) RestPostTopic(w rest.ResponseWriter, r *rest.Request) {
	app.Context.Stats.AddHit(r.RequestURI)

	topic := r.PathParam("topic")
	var msg Message
	if err := r.DecodeJsonPayload(&msg); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("appending: %s:%p to topic %s in %v, which is %d in length.\n", msg.ToJson(), &msg, topic, *m, len(m.topics))
	m.Append(topic, &msg)
	w.WriteJson(&msg)
}
