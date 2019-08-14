package comparisons

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

/*
if_acmpeq和if_acmpne指令把栈顶的两个引用弹出，根据引用是否相同进行跳转
 */
type IF_ACMPEQ struct {
	base.NoOperandsInstruction
}
type IF_ACMPNE struct {
	base.NoOperandsInstruction
}

func (self *IF_ACMPEQ)Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	ref2:=stack.PopRef()
	ref1:=stack.PopRef()
	if ref1==ref2{
		base.Branch(frame,self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack();
	ref1:=stack.PopRef()
	ref2:=stack.PopRef()
	if ref1!=ref2{
		base.Branch(frame,self.Offset)
	}
}

