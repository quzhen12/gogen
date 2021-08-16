package common

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func CopyFiles(dst string, src string) error {

	file, err := os.Stat(src)
	if err != nil {
		return err
	}
	if strings.Contains(src, ".git") {
		return nil
	}
	dst = path.Join(dst, file.Name())
	err = copy(dst, src, file.IsDir())
	if err != nil {
		return err
	}
	if !file.IsDir() {
		return nil
	}
	rd, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, f := range rd {
		err := CopyFiles(dst, path.Join(src, f.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

func copy(dst, src string, isDir bool) error {
	if isDir {
		_, err := os.Stat(dst)
		if os.IsNotExist(err) {
			return os.Mkdir(dst, 0700)
		}
		return nil
	}
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, data, 0700)

}
