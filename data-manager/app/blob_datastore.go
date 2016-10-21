package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
func (m BlobDatastore) Read() (models.Content, error) {
	var (
		blob    models.Blob
		content models.Content
	)
	err := MakeRequest(
		fmt.Sprintf("%s/%d", m.blobServiceURL, blobId),
		http.MethodGet,
		&blob,
		nil,
	)
	if err != nil {
		return content, err
	}

	err = content.FromJSON([]byte(blob.Content))
	if err != nil {
		return content, err
	}
	return content, nil
}

// Write is a function that stores data.
func (m BlobDatastore) Write(c models.Content) error {
	var (
		payloadBlob  models.Blob
		returnedBlob models.Blob
	)

	// Get the blob from the service
	err := MakeRequest(
		fmt.Sprintf("%s/%d", m.blobServiceURL, blobId),
		http.MethodGet,
		&payloadBlob,
		nil,
	)
	if err != nil {
		return err
	}

	// Prepare the content for the blob payload
	contentString, err := c.ToJSON()
	if err != nil {
		return err
	}
	payloadBlob.Content = contentString
	b, err := json.Marshal(payloadBlob)
	if err != nil {
		return err
	}

	// Update the blob
	err = MakeRequest(
		fmt.Sprintf("%s/%d", m.blobServiceURL, blobId),
		http.MethodPost,
		&returnedBlob,
		bytes.NewReader(b),
	)
	if err != nil {
		return err
	}

	return nil
}
