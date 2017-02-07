package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(absPath)
	}
	return &ZipEntry{absPath}
}

func (ze *ZipEntry) String() string {
	return ze.absPath
}

func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(ze.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.Name == className {
			content, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer content.Close()

			data, err := ioutil.ReadAll(content)
			if err != nil {
				return nil, nil, err
			}
			return data, ze, nil
		}
	}

	return nil, nil, errors.New("class not found:" + className)
}
