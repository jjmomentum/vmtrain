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
	"sync"
)

// Blob is a struct used to store data into the blob service.
type Content struct {
	Servers      map[string]Server      `json:"servers"`
	Reservations map[string]Reservation `json:"reservations"`
	Users        map[string]User        `json:"users"`
	sync.RWMutex
}

// FromJSON populates from JSON data.
func (c *Content) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns a JSON formatted of the struct.
func (c *Content) ToJSON() (string, error) {
	blobJson, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(blobJson), nil
}
