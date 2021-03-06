package control

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame,self.Offset)
}

