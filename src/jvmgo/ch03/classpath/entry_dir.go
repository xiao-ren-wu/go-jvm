package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//只有一个字段，用于存放目录的绝对路径
type DirEntry struct {
	absDir string
}
//首先把参数转换成绝对路径，如果转换出错，调用panic()函数终止程序运行，否则创建DirEntry实例并返回
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err !=nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
//先把目录和文件夹拼接城一个完整的路径，然后调用ioUtil包提供的ReadFile()函数读取class文件内容，最后返回
func (self *DirEntry) readClass(className string) ([]byte,Entry,error){
	fileName := filepath.Join(self.absDir,className)
	data, err :=ioutil.ReadFile(fileName)
	return data, self, err
}
//直接返回目录
func (self *DirEntry) String() string {
	return self.absDir
}

