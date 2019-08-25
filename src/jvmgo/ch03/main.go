package main

import (
	"fmt"
	"jvmgo/ch03/classfile"
	"jvmgo/ch03/classpath"
	"jvmgo/ch03/rtda/heap"
	"strings"
)

func main() {
	//cmd := parseCmd()
	//	//if cmd.versionFlag {
	//	//	fmt.Println("version 0.0.1")
	//	//}else if cmd.helpFlag || cmd.class == "" {
	//	//	printUsage()
	//	//}else{
	//	//	startJVM(cmd)
	//	//}
	//cmd := &Cmd{XjreOption: "C:\\Program Files\\Java\\jre1.8.0_191", class: "java.lang.String"}
	startJVM()
}

func startJVM1(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}

func startJVM2(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Printf(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count:%v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags:0x%x\n", cf.AccessFlags())
	fmt.Printf("this.class:%v\n", cf.ClassName())
	fmt.Printf("super.class:%v\n", cf.SuperClassName())
	fmt.Printf("interfaces:%v\n", cf.InterfacesNames())
	fmt.Printf("fields count:%v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count:%v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}

//func startJVM12() {
//	cmd := &Cmd{XjreOption: "H:\\go-jvm", class: "src.jvmgo.ch03.GaussTest"}
//	className := strings.Replace(cmd.class, ".", "/", -1)
//	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
//	cf := loadClass(className, cp)
//	mainMethod := getMainMethod(cf)
//	if mainMethod != nil {
//		interpreter(mainMethod)
//	} else {
//		fmt.Printf("No Main method")
//	}
//}

//func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
//	for _, m := range cf.Methods() {
//		if m.Name() == "main" {
//			return m
//		}
//	}
//	return nil
//}
func startJVM() {
	cmd := &Cmd{XjreOption: "C:\\Program Files\\Java\\jre1.8.0_191", class: "src.jvmgo.ch03.MyObject"}
	cp:=classpath.Parse(cmd.XjreOption,cmd.cpOption)
	classloader:=heap.NewClassLoader(cp)
	className:=strings.Replace(cmd.class,".","/",-1)
	mainClass:=classloader.LoadClass(className)
	mainMethod:=mainClass.GetMainMethod()
	if mainMethod !=nil{
		interpreter(mainMethod)
	}else{
		fmt.Printf("Main method not find in class %s\n",cmd.class)
	}
}

