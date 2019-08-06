package rtda
//栈帧
type Frame struct {
	lower        *Frame        //用于实现链表数据结构
	localVars    LocalVars     //保存局部变量表指针
	operandStack *OperandStack //操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
