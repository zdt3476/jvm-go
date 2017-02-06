package classpath

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := new(Classpath)
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)

	return cp
}

func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className += ".class"

	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := cp.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return cp.userClasspath.readClass(className)
}

func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jrePath, err := getJre(jreOption)
	if err != nil {
		log.Fatalln(err)
	}

	// jre/lib/*
	bootPath := filepath.Join(jrePath, "lib", "*")
	cp.bootClasspath = newWildcardEntry(bootPath)

	// jre/lib/ext/*
	extPath := filepath.Join(jrePath, "lib", "ext", "*")
	cp.extClasspath = newWildcardEntry(extPath)
}

func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	cp.userClasspath = newEntry(cpOption)
}

func getJre(jreOption string) (string, error) {
	if jreOption != "" && exists(jreOption) {
		return jreOption, nil
	}

	if exists("./jre") {
		return "./jre", nil
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" && exists(jh) {
		return filepath.Join(jh, "jre"), nil
	}

	return "", errors.New("Can not found the jre folder!")
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsExist(err)
}
