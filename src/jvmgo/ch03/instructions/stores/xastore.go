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

/**
<t>store指令的三个操作数分别是：
1.赋给数组元素的值
2.数组索引
3.数组引用
依次从操作数栈中弹出
如果引用是null,则抛出NullPointerException异常
如果数组索引小于0或者大于等于数组的长度，则抛出ArrayIndexOutOfBoundsException
如果一切正常，按索引给数组赋值
*/

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
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	bytes[index] = int8(val)
}
func (self *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	chars[index] = int16(val)
}
func (self *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	doubles[index] = float64(val)
}
func (self *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	floats[index] = float32(val)
}
func (self *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	longs[index] = int64(val)
}
func (self *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	shorts[index] = int16(val)
}

func (self *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = val
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
