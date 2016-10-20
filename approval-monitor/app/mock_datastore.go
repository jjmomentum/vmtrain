package app

import (
	"fmt"

	"github.com/vmtrain/approval-monitor/models"
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
func (m MockDatastore) Read(id int) (*models.Blob, error) {
	blob, ok := m.blobMap[id]
	if !ok {
		return nil, fmt.Errorf("Blob with ID: %d not found", id)
	}

	return blob, nil

}

// Write is a function that stores data in memory.
func (m MockDatastore) Write(b *models.Blob) error {
	m.blobMap[b.ID] = b
	return nil
}
