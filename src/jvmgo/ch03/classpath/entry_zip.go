package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

/**
首先打开ZIP文件，如果这一步出错直接返回
然后遍历ZIP压缩包里的文件，看看能否找到class文件，如果能找到，则打开class文件，把内容读出来并返回
如果找不到，或者出现其他错误，则返回错误信息，有两处使用了defer语句来确保打开的文件得以关闭
*/
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}
func (self *ZipEntry) String() string {
	return self.absPath
}
