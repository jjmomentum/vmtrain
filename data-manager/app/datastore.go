// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import "github.com/vmtrain/data-manager/models"

// Constants used to map the queries to tables and columns found in the
// database schema
const (
	blobId = 04
)

// Datastore is an interface to abstract the reading and writing of data.
type Datastore interface {
	Write(b models.Content) error
	Read() (models.Content, error)
}
