package references

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	methodRef:=cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod :=methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic(){
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame, resolvedMethod)
}
