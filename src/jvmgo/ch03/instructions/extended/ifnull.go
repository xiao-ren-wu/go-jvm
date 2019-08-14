package extended

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}
type IFNOTNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFNOTNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
