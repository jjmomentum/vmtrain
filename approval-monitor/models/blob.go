// Package models has the structs that are used by the application.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package models

// Blob is a struct used to store data into the blob service.
type Blob struct {
	ID      int     `json:"id"`
	Version string  `json:"version"`
	Name    string  `json:"name"`
	Content Content `json:"content"`
	Tag     string  `json:"tag"`
}
