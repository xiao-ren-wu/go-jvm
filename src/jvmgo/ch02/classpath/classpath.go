package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

/**
	-Xjre 选项解析启动类路径和拓展类路径，
	使用-classpath/-cp选项解析用户类路径
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp:=&Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
/**
	优先使用用户的-Xjre选项作为jre目录，如果没有输入该选项，则在当前目录下寻找jre目录
	如果找不到，尝试使用JAVA_HOME环境变量
 */
func(self *Classpath) parseBootAndExtClasspath(jreOption string)  {
	jreDir:=getJreDir(jreOption)
	//jre/lib/*
	jreLibPath:=filepath.Join(jreDir,"lib","*")
	self.bootClasspath=newWildcardEntry(jreLibPath)
	//jre/lib/ext/*
	jreExtPath:=filepath.Join(jreDir,"lib","ext","*")
	self.extClasspath=newWildcardEntry(jreExtPath)
}


func (self *Classpath)parseUserClasspath(cpOption string)  {
	if cpOption == ""{
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption){
		return jreOption
	}
	if exists("./jre"){
		return "./jre"
	}
	if jh:=os.Getenv("JAVA_HOME");jh!=""{
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder!")
}

/**
	用于判断目录是否存在
 */
func exists(path string) bool {
	if _,err :=os.Stat(path);err!=nil{
		if os.IsExist(err){
			return false
		}
	}
	return true
}

/**
	如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径。
	ReadClass方法依次从启动类路径，拓展类路径和用户搜索class文件
	传递给ReadClass方法的类名不包含".class"后缀。
 */
func (self *Classpath)ReadClass (className string) ([]byte,Entry,error){
	className = className +".class"
	if data,entry,err:=self.bootClasspath.readClass(className);err==nil{
		return data,entry,err
	}
	if data,entry,err:=self.extClasspath.readClass(className);err==nil{
		return data,entry,err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string{
	return self.userClasspath.String()
}

