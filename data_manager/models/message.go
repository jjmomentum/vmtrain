// Package models has the structs that are used by the application.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package models

import (
	"encoding/json"
	"log"
)

// Message for storing in journal queues.
type Message struct {
	// Unique id for this Message.
	ID int `json:"id"`

	// The message, which in JSON must be Base64 encoded.
	Base64 []byte `json:"message"`
}

// Initialize a Message back to 'empty' (effectively, zeroed out).
func (m *Message) init() {
	m.ID = 0
	m.Base64 = make([]byte, 0)
}

// FromJSON populates a message from JSON data.
func (m *Message) FromJSON(bytes []byte) {
	m.init()
	if err := json.Unmarshal(bytes, m); err != nil {
		log.Println(err)
		m.init()
	}
}

// ToJSON returns a JSON formatted string representation of the Message.
func (m *Message) ToJSON() string {
	b, _ := json.Marshal(m)
	return string(b)
}
