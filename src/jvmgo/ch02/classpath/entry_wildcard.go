package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
	首先把路径末尾的*去掉，得到baseDir,然后调用filepath包中的Walk函数遍历baseDir创建ZipEntry.
	Walk()函数的第二个参数也是一个函数，
 */
func newWildcardEntry(path string) CompositeEntry {
	baseDir:=path[:len(path)-1] //remove *
	compositeEntry:=[]Entry{}
	/**
		根据后缀名选出JAR文件，
		并且返回SkipDir跳过子目录（通配符类路径不能递归匹配子目录下的JAR文件）
	 */
	walkFn:= func(path string,info os.FileInfo,err error) error{
		if err!=nil{
			return err
		}
		if info.IsDir()&&path!=baseDir{
			return filepath.SkipDir
		}
		if strings.HasSuffix(path,".jar")||strings.HasSuffix(path,".JAR"){
			jarEntry:=newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir,walkFn)
	return compositeEntry
}
