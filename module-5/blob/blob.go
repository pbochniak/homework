package blob

type BlobStorage interface {
	Create(string, []byte) error
	Read(string) ([]byte, error)
	Update(string, []byte) error
	Delete(string) error
}
