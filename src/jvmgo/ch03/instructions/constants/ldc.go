package constants

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

/*
用于从运行时常量池中加载常量值，并将值推入操作数栈中
ldc和ldc_w
	指令用于加载int，float，和字符串常量，java.lang.Class实例或者MethodType和MethodHandler实例
ldc2_w
	用于加载long或者double常量
ldc和ldc_w指令区别仅在于操作数的宽度
*/
type LDC struct {
	base.Index8Instruction
}
type LDC_w struct {
	base.Index16Instruction
}
type LDC2_w struct {
	base.Index16Instruction
}

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}
func (self *LDC_w) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}
func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	class := frame.Method().Class()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classRef:=c.(*heap.ClassRef)
		classObj:=classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	default:
		panic("todo ldc")
	}
}
func (self *LDC2_w) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
