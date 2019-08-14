package stores

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

/*
存储指令：
	把变量从操作数栈顶弹出，然后存入局部变量表
*/
type LSTORE struct {
	base.Index8Instruction
}
type LSTORE_0 struct {
	base.NoOperandsInstruction
}
type LSTORE_1 struct {
	base.NoOperandsInstruction
}
type LSTORE_2 struct {
	base.NoOperandsInstruction
}
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}
func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}
func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}
func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}
func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
