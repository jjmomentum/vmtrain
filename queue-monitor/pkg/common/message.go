// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package common

import (
	"encoding/json"
	"log"
)

// Messsage for storing in journal queues.
type Message struct {
	// Unique id for this Message.
	Id int `json:"id"`

	// The message, which in JSON must be Base64 encoded.
	Base64 []byte `json:"message"`
}

// Initialize a Message back to 'empty' (effectively, zeroed out).
func (m *Message) init() {
	m.Id = 0
	m.Base64 = make([]byte, 0)
}

// Populate a message from JSON data.
func (m *Message) FromJson(bytes []byte) {
	m.init()
	if err := json.Unmarshal(bytes, m); err != nil {
		log.Println(err)
		m.init()
	}
}

// Return a JSON formatted string representation of the Message.
func (m *Message) ToJson() string {
	b, _ := json.Marshal(m)
	return string(b)
}
