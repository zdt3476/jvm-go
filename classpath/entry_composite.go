package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	paths := strings.Split(pathList, pathListSeparator)
	ce := make(CompositeEntry, len(paths))
	for i := range ce {
		entry := newEntry(paths[i])
		ce[i] = entry
	}
	return ce
}

func (ce CompositeEntry) String() string {
	arr := []string{}
	for _, entry := range ce {
		arr = append(arr, entry.String())
	}

	return strings.Join(arr, pathListSeparator)
}

func (ce CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range ce {
		data, e, err := entry.readClass(className)
		if err == nil {
			return data, e, nil
		}
	}

	return nil, nil, errors.New("class not found:" + className)
}
