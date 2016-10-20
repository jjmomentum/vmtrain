// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import "github.com/vmtrain/approval-monitor/models"

// Datastore is an interface to abstract the reading and writing of data.
type Datastore interface {
	Write(b *models.Blob) error
	Read(id int) (*models.Blob, error)
}
