package loads

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

//Load int from variable
//加载指令：从局部变量表中获取变量，然后推入栈操作数栈顶
type ILOAD struct {
	base.Index8Instruction
}

//后面的数字代表指令的索引
type ILOAD_0 struct {
	base.NoOperandsInstruction
}
type ILOAD_1 struct {
	base.NoOperandsInstruction
}
type ILOAD_2 struct {
	base.NoOperandsInstruction
}
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}
func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}
func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}
func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}
func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
