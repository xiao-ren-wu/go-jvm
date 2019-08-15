package rtda

//栈帧
type Frame struct {
	lower        *Frame        //用于实现链表数据结构
	localVars    LocalVars     //保存局部变量表指针
	operandStack *OperandStack //操作数栈指针
	thread       *Thread       //栈帧所属线程
	nextPC       int           //程序计数器
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
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
// getters & setters
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}
