package loads

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

type AALOAD struct {
	base.NoOperandsInstruction
}
type BALOAD struct {
	base.NoOperandsInstruction
}
type CALOAD struct {
	base.NoOperandsInstruction
}
type DALOAD struct {
	base.NoOperandsInstruction
}
type FALOAD struct {
	base.NoOperandsInstruction
}
type LALOAD struct {
	base.NoOperandsInstruction
}
type IALOAD struct {
	base.NoOperandsInstruction
}
type SALOAD struct {
	base.NoOperandsInstruction
}

func (self *AALOAD) Execute(frame *rtda.Frame) {
	_xaload(frame)
}
func (self *BALOAD) Execute(frame *rtda.Frame) {
	_xaload(frame)
}
func (self *CALOAD) Execute(frame *rtda.Frame) {
	_xaload(frame)
}
func (self *DALOAD) Execute(frame *rtda.Frame) {
	_xaload(frame)
}
func (self *FALOAD) Execute(frame *rtda.Frame) {
	_xaload(frame)
}
func (self *IALOAD) Execute(frame *rtda.Frame) {
	_xaload(frame)
}
func (self *SALOAD) Execute(frame *rtda.Frame) {
	_xaload(frame)
}
func (self *LALOAD) Execute(frame *rtda.Frame) {
	_xaload(frame)
}

func _xaload(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}
func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lanng.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
