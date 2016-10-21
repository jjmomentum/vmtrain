// Package models has the structs that are used by the application.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package models

import (
	"encoding/base64"
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

// FromBase64 populates from base64 string data.
func (c *Content) FromBase64(b64 string) error {
	sDec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return err
	}

	err = json.Unmarshal(sDec, c)
	if err != nil {
		return err
	}
	return nil
}

// ToBase64 returns a base64 string of the struct.
func (c *Content) ToBase64() (string, error) {
	blobJson, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(blobJson), nil
}
