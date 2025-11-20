package storage

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

type ZipStorage struct {
	*Storage
}

func NewZipStorage(filename string) *ZipStorage {
	return &ZipStorage{
		&Storage{FileName: filename},
	}
}

func (s *ZipStorage) Save(data []byte) error {
	f, err := os.Create(s.GetFileName())
	if err != nil {
		return err
	}

	zw := zip.NewWriter(f)

	var w io.Writer
	w, err = zw.Create("calendar.zip")
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	err = zw.Close()
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *ZipStorage) Load() ([]byte, error) {
	r, err := zip.OpenReader(s.GetFileName())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if len(r.File) == 0 {
		return nil, errors.New("zip is empty")
	}

	file := r.File[0]

	var rc io.ReadCloser
	rc, err = file.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return io.ReadAll(rc)
}
