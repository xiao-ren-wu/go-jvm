package references

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

type INVOKE_SPECIAL struct {
	base.Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PopRef()
}

