package stores

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

type AASTORE struct {
	base.NoOperandsInstruction
}
type BASTORE struct {
	base.NoOperandsInstruction
}
type CASTORE struct {
	base.NoOperandsInstruction
}
type DASTORE struct {
	base.NoOperandsInstruction
}
type FASTORE struct {
	base.NoOperandsInstruction
}
type IASTORE struct {
	base.NoOperandsInstruction
}
type LASTORE struct {
	base.NoOperandsInstruction
}
type SASTORE struct {
	base.NoOperandsInstruction
}

func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}
func (self *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Bytes()
	checkIndex(len(ints), index)                       //todo 变量名称
	ints[index] = int8(val)
}
func (self *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Chars()
	checkIndex(len(ints), index)
	ints[index] = int16(val)
}
func (self *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Chars()
	checkIndex(len(ints), index)
	ints[index] = int16(val)
}
func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
