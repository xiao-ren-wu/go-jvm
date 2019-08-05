package rtda

/**
Java虚拟机栈
 */
type Thread struct {
	pc		int
	stack 	*Stack
}
/*
参数表示Stack最多可以容纳多少帧
 */
func NewThread() *Thread{
	return &Thread{
		stack:newThread(1024),
	}
}
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPc(pc int) {
	self.pc = pc
}
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()_
}
/**
返回当前帧
 */
func (self *Thread) CurrentFrame() *Frame{
	return self.stack.top()
}