package references

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

/*
new指令的操作数是一个uint16索引，来自字节码，
通过这个索引，可以从当前类的运行时常量池中找到一个类符号引用
拿到类数据然后创建对象，并把对象引用引入栈顶
*/
type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted(){
		frame.RevertNextPC()
		base.InitClass(frame.Thread(),class)
		return
	}
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
