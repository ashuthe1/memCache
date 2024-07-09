package persistence

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"

	"github.com/ashuthe1/kuki-memcache/cache"
)

// FilePersistence handles saving and loading data to/from a file.
type FilePersistence struct {
	filePath string
	mu       sync.Mutex
}

// NewFilePersistence creates a new FilePersistence instance with the given filePath.
func NewFilePersistence(filePath string) *FilePersistence {
	return &FilePersistence{
		filePath: filePath,
	}
}

// SaveToFile saves data to a file.
func (fp *FilePersistence) SaveToFile(data interface{}) error {
	fp.mu.Lock()
	defer fp.mu.Unlock()

	file, err := os.Create(fp.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// LoadFromFile loads data from a file.
func (fp *FilePersistence) LoadFromFile() (map[string]cache.CacheItem, error) {
	fp.mu.Lock()
	defer fp.mu.Unlock()

	file, err := os.Open(fp.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data map[string]cache.CacheItem
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	// Adjust expiration times based on current time
	for key, item := range data {
		item.ExpiresAt = time.Now().Add(item.TTL)
		data[key] = item
	}

	return data, nil
}
