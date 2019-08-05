package rtda

type Frame struct {
	lower        *Frame        //用于实现链表数据结构
	localVars    LocalVars     //保存局部变量表指针
	operandStack *OperandStack //操作数栈指针
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
