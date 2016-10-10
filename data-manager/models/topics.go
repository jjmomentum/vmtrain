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

// Topics is a JSON transfer object for providing the list of topics
// currently available in the journal (where available just means a queue
// exists for the topic, though may not have an available Message).
type Topics struct {
	Topics []string `json:"topics"`
}

// FromJSON populates Topics from JSON data.
func (t *Topics) FromJSON(bytes []byte) {
	if err := json.Unmarshal(bytes, t); err != nil {
		log.Println(err)
		t.Topics = make([]string, 0)
	}
}

// ToJSON returns a JSON formatted string representation of the Topics.
func (t *Topics) ToJSON() string {
	b, err := json.Marshal(t)
	if err != nil {
		b = make([]byte, 0)
	}
	return string(b)
}
