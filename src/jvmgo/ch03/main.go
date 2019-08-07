package main

import (
	"fmt"
	"jvmgo/ch03/classfile"
	"jvmgo/ch03/classpath"
	"jvmgo/ch03/rtda"
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

func startJVM() {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 299724580)
	vars.SetLong(4, -299724580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.7182812845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}
func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(299724580)
	ops.PushLong(-299724580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.7182812845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}
