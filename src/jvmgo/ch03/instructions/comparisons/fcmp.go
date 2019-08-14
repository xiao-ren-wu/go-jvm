package comparisons

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

type FCMPG struct {
	base.NoOperandsInstruction
}
type FCMPL struct {
	base.NoOperandsInstruction
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

/*
当两个两个变量中至少有一个是NaN时，用fcmpg指令比较结果是1，用fcmpl指令比较是-1
 */
func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame,true)
}
func (self *FCMPL) Execute(frame *rtda.Frame)  {
	_fcmp(frame,false)
}
