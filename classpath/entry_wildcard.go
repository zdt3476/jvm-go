package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	ce := []Entry{}
	fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && baseDir != path {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") ||
			strings.HasSuffix(path, ".JAR") {
			entry := newZipEntry(path)
			ce = append(ce, entry)
		}

		return nil
	}
	filepath.Walk(baseDir, fn)

	return ce
}
