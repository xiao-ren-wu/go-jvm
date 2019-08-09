package math

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

/*
位移指令
*/

//int 左位移
type ISHL struct {
	base.NoOperandsInstruction
}

//int 右位移
type ISHR struct {
	base.NoOperandsInstruction
}

//int 逻辑右位移
type IUSHR struct {
	base.NoOperandsInstruction
}

//long 左位移
type LSHL struct {
	base.NoOperandsInstruction
}

// long 右位移
type LSHR struct {
	base.NoOperandsInstruction
}

//long 逻辑右位移
type LUSHR struct {
	base.NoOperandsInstruction
}

/*
从操作数栈中弹出两个变量
v1:要位移的变量
v2:指出要位移多少比特
位移之后把结果推入操作数栈中
[注意]
	1.int变量只有32位，所以，只要取出v2的前5个比特足够表示位移数了
	2.Go语言的位移操作符右侧必须是无符号整数，所以要对v2进行类型转换
*/
func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}
func (self *ISHR) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	s:=uint32(v2)&0x1f
	result:=v1>>s
	stack.PushInt(result)
}
func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}
func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopLong()
	s:=uint32(v2)&0x3f
	result:=v1>>s
	stack.PushLong(result)
}
func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopLong()
	s:=uint32(v2)&0x3f
	result:=int64(uint64(v1)>>s)
	stack.PushLong(result)
}






















