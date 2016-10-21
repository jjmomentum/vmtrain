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
)

// Blob is a struct used to store data into the blob service.
type Blob struct {
	ID      int    `json:"id"`
	Version string `json:"version"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

// FromJSON populates Reservation from JSON data.
func (b *Blob) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, b)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns a JSON formatted string representation of the Reservation.
func (b *Blob) ToJSON() (string, error) {
	blobJson, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(blobJson), nil
}
