package memory

import "fmt"

type BlobStorage struct {
	memory map[string][]byte
}

func NewBlobStorage() *BlobStorage {
	storage := new(BlobStorage)
	storage.memory = make(map[string][]byte)
	return storage
}

func (storage *BlobStorage) Create(key string, data []byte) error {
	if _, ok := storage.memory[key]; ok {
		return fmt.Errorf("cannot add key %q to storage, because already exists", key)
	}
	storage.memory[key] = data
	return nil
}

func (storage *BlobStorage) Read(key string) ([]byte, error) {
	if _, ok := storage.memory[key]; ok {
		return storage.memory[key], nil
	}
	return nil, fmt.Errorf("cannot read key %q from storage, because doesn't exist", key)
}

func (storage *BlobStorage) Update(key string, data []byte) error {
	if _, ok := storage.memory[key]; ok {
		storage.memory[key] = data
		return nil
	}
	return fmt.Errorf("cannot update key %q in storage, because doesn't exist", key)
}

func (storage *BlobStorage) Delete(key string) error {
	if _, ok := storage.memory[key]; ok {
		delete(storage.memory, key)
		return nil
	}
	return fmt.Errorf("cannot delete key %q from storage, because doesn't exist", key)
}
