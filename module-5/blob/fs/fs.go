package fs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Row struct {
	Key  string
	Data []byte
}

func (row Row) ToBytes() []byte {
	return bytes.Join([][]byte{[]byte(row.Key), []byte(strconv.Itoa(len(row.Data))), row.Data}, []byte{0})
}

func (row *Row) FromBytes(data []byte) int {
	pos := 0
	index := bytes.IndexByte(data[pos:], 0)
	row.Key = string(data[pos : pos+index])
	pos += index + 1
	index = bytes.IndexByte(data[pos:], 0)
	length, _ := strconv.Atoi(string(data[pos : pos+index]))
	pos += index + 1
	row.Data = data[pos : pos+length]
	return pos + length
}

type BlobStorage struct {
	path string
}

func NewBlobStorage(path string) *BlobStorage {
	storage := new(BlobStorage)
	storage.path = path
	return storage
}

func (storage *BlobStorage) Create(key string, data []byte) error {
	content, err := ioutil.ReadFile(storage.path)
	if err != nil {
		return err
	}
	row := new(Row)
	for position := 0; position < len(content); {
		position += row.FromBytes(content[position:])
		if row.Key == key {
			return fmt.Errorf("cannot add key %q to storage, because already exists", key)
		}
	}
	row.Key = key
	row.Data = data
	content = append(content, row.ToBytes()...)
	return ioutil.WriteFile(storage.path, content, 0644)
}

func (storage *BlobStorage) Read(key string) ([]byte, error) {
	content, err := ioutil.ReadFile(storage.path)
	if err != nil {
		return nil, err
	}
	row := new(Row)
	for position := 0; position < len(content); {
		position += row.FromBytes(content[position:])
		if row.Key == key {
			return row.Data, nil
		}
	}
	return nil, fmt.Errorf("cannot read key %q from storage, because doesn't exist", key)
}

func (storage *BlobStorage) Update(key string, data []byte) error {
	content, err := ioutil.ReadFile(storage.path)
	if err != nil {
		return err
	}
	row := new(Row)
	for position := 0; position < len(content); {
		rowLength := row.FromBytes(content[position:])
		if row.Key == key {
			row.Data = data
			newContent := []byte{}
			newContent = append(newContent, content[:position]...)
			newContent = append(newContent, row.ToBytes()...)
			newContent = append(newContent, content[position+rowLength:]...)
			return ioutil.WriteFile(storage.path, newContent, 0644)
		}
		position += rowLength
	}
	return fmt.Errorf("cannot update key %q in storage, because doesn't exist", key)
}

func (storage *BlobStorage) Delete(key string) error {
	content, err := ioutil.ReadFile(storage.path)
	if err != nil {
		return err
	}
	row := new(Row)
	for position := 0; position < len(content); {
		rowLength := row.FromBytes(content[position:])
		if row.Key == key {
			newContent := []byte{}
			newContent = append(newContent, content[:position]...)
			newContent = append(newContent, content[position+rowLength:]...)
			return ioutil.WriteFile(storage.path, newContent, 0644)
		}
		position += rowLength
	}
	return fmt.Errorf("cannot delete key %q from storage, because doesn't exist", key)
}
