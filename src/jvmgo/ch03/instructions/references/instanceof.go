package references

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

/*
第一个指令时uint16索引，从方法的字节码中获取，通过这个索引可以从当前类的运行时常量池中找到类符号引用
第二个操作数是对象引用，从操作数栈中弹出

先判断对象引用，如果是null,则把0推入栈中，
如果引用obj是null的话，不管ClassYYY是什么类型，都是false
*/
type INSTANCE_OF struct {
	base.Index16Instruction
}

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
