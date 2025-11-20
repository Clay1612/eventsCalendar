package storage

type Store interface {
	Save(data []byte) error
	Load() ([]byte, error)
	GetFileName() string
}
type Storage struct {
	FileName string
}

func (s *Storage) GetFileName() string {
	return s.FileName
}
