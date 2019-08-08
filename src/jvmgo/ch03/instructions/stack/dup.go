package stack

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)
//复制栈顶变量
type DUP struct {
	base.NoOperandsInstruction
}
type DUP_X1 struct {
	base.NoOperandsInstruction
}
type DUP_X2 struct {
	base.NoOperandsInstruction
}
type DUP2 struct {
	base.NoOperandsInstruction
}
type DUP2_X1 struct {
	base.NoOperandsInstruction
}
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot:=stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	val1:=stack.PopSlot()
	val2:=stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}
func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}
func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	val4 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val4)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}
