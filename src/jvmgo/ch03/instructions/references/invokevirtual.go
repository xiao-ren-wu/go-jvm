package references

import (
	"fmt"
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MemberRef)

	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(C)V":
			fmt.Printf("%c\n", stack.PopInt())
		case "(B)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(S)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(I)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(J)V":
			fmt.Printf("%v\n", stack.PopLong())
		case "(F)V":
			fmt.Printf("%v\n", stack.PopFloat())
		case "(D)V":
			fmt.Printf("%v\n", stack.PopDouble())
		default:
			panic("println:" + methodRef.Descriptor())
		}
		stack.PopRef()
	}
}
