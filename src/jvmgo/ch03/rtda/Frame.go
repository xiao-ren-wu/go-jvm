package rtda

import "jvmgo/ch03/rtda/heap"

//栈帧
type Frame struct {
	lower        *Frame        //用于实现链表数据结构
	localVars    LocalVars     //保存局部变量表指针
	operandStack *OperandStack //操作数栈指针
	thread       *Thread       //栈帧所属线程
	method		 *heap.Method  //栈帧所属方法
	nextPC       int           //程序计数器
	maxLocals    uint
	maxStack     uint
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		maxLocals:    method.MaxLocals(),
		maxStack:     method.MaxStack(),
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
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

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}
