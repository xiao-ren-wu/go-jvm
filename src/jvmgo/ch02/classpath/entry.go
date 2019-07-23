package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator string = string(os.PathSeparator)

type Entry interface {
	//负责寻找和加载class文件
	/**
	 * 参数是class文件的相对路径，文件之间用斜线（/）分隔，文件名有.class后缀
		返回值是读取到的字节数组，最终定位到class文件的Entry，以及错误信息
	 */
	readClass(className string) ([]byte, Entry, error)
	//相当于Java中的toString方法
	String() string
}

/**
	根据参数创建不同的Entry实例
 */
func newEntry(path string) Entry {
	if strings.Contains(path,pathListSeparator){
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR")||
		strings.HasSuffix(path,".zip")||strings.HasSuffix(path,".ZIP"){
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
