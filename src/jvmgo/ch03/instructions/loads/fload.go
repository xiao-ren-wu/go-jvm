package loads

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

type FLOAD struct {
	base.Index8Instruction
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame,uint(self.Index))
}

type FLOAT_0 struct {
	base.NoOperandsInstruction
}

func (self *FLOAT_0) Execute(frame *rtda.Frame) {
	_fload(frame,0)
}

type FLOAT_1 struct {
	base.NoOperandsInstruction
}

func (self *FLOAT_1) Execute(frame *rtda.Frame) {
	_fload(frame,1)
}

type FLOAT_2 struct {
	base.NoOperandsInstruction
}

func (self *FLOAT_2) Execute(frame *rtda.Frame) {
	_fload(frame,2)
}

type FLOAT_3 struct {
	base.NoOperandsInstruction
}

func (self *FLOAT_3) Execute(frame *rtda.Frame) {
	_fload(frame,3)
}
func _fload(frame *rtda.Frame,index uint)  {
	val:=frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
