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
func (m MockDatastore) Read(b *models.Blob) error {
	b, ok := m.blobMap[b.ID]
	if !ok {
		return fmt.Errorf("Blob with ID: %d not found", b.ID)
	}

	return nil

}

// Write is a function that stores data in memory.
func (m MockDatastore) Write(b *models.Blob) error {
	m.blobMap[b.ID] = b
	return nil
}
