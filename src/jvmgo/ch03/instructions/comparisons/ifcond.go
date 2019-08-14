package comparisons

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)
/*
	if<cond>
	if<cond>指令把操作数栈顶的int变量弹出，然后0跟进比较，满足则跳转
		假设从栈顶弹出的变量是X，则跳转如下：
			ifeq: x==0
			ifne: x!=0
			iflt: x<0
			ifgt: x>0
			ifge: x>=0
			ifle: x<=0
 */

type IFEQ struct {
	base.NoOperandsInstruction
}
type IFNE struct {
	base.NoOperandsInstruction
}
type IFLT struct {
	base.NoOperandsInstruction
}
type IFLE struct {
	base.NoOperandsInstruction
}
type IFGT struct {
	base.NoOperandsInstruction
}
type IFGE struct {
	base.NoOperandsInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val:=frame.OperandStack().PopInt()
	if val==0{
		base.Branch(frame,self.Offset)
	}
}
