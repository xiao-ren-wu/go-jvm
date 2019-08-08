package stack

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}
type POP2 struct {
	base.NoOperandsInstruction
}

//用于弹出int，float等占用一个操作数栈位置的变量
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

//用于弹出Long，Double变量在操作数栈中占据两个位置
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
