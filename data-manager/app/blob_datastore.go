package app

import (
	"fmt"

	"github.com/vmtrain/data-manager/models"
)

// MockDatastore is a struct to simulate the datastore of the API server during test.
type BlobDatastore struct {
	blobServiceURL string
}

// NewBlobDatastore is a constructor that returns a new instance of the BlobDatastore
// struct.
func NewBlobDatastore(url string) BlobDatastore {
	return BlobDatastore{
		blobServiceURL: url,
	}
}

// Read is a function that lookps up data.
func (m BlobDatastore) Read(id int) (*models.Blob, error) {

	return nil, nil

}

// Write is a function that stores data.
func (m BlobDatastore) Write(b *models.Blob) error {

	return nil
}
