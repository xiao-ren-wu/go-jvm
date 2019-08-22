package references

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)
/*
先从操作数栈中弹出对象引用，再推回去，这样不会更改操作数栈状态
如果引用是null,则指令执行结束，null可以转换成任意类型
不为空，解析类符号引用，判断对象是否是类的实例，如果是，则指令执行结束，否则抛出ClassCastException
 */
type CHECK_CAST struct {
	base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	ref:=stack.PopRef()
	stack.PushRef(ref)
	if ref==nil{
		return
	}
	cp:=frame.Method().Class().ConstantPool()
	classRef:=cp.GetConstant(self.Index).(*heap.ClassRef)
	class:=classRef.ResolvedClass()
	if !ref.IsInstanceOf(class){
		panic("java.lang.ClassCastException")
	}
}

