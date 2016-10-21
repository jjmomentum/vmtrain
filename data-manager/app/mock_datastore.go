package app

import (
	"fmt"

	"github.com/vmtrain/data-manager/models"
)

// MockDatastore is a struct to simulate the datastore of the API server during test.
type MockDatastore struct {
	blobMap map[int]*models.Blob
}

// NewMockDatastore is a constructor that returns a new instance of the MockBackend
// struct.
func NewMockDatastore(blobs map[int]*models.Blob) MockDatastore {
	return MockDatastore{
		blobMap: blobs,
	}
}

// Read is a function that lookps up data in memory.
func (m MockDatastore) Read() (models.Content, error) {
	content := models.Content{}
	blob, ok := m.blobMap[blobId]
	if !ok {
		return content, fmt.Errorf("Blob with ID: %d not found", blobId)
	}

	err := content.FromJSON([]byte(blob.Content))
	if err != nil {
		return content, err
	}

	return content, nil

}

// Write is a function that stores data in memory.
func (m MockDatastore) Write(content models.Content) error {
	contentString, err := content.ToJSON()
	if err != nil {
		return err
	}
	m.blobMap[blobId].Content = contentString
	return nil
}
