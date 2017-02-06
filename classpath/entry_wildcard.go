package classpath

import (
	"os"
	"path/filepath"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	ce := []Entry{}
	fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return filepath.SkipDir
		}

		entry := newEntry(path)
		ce = append(ce, entry)

		return nil
	}
	filepath.Walk(baseDir, fn)

	return ce
}
