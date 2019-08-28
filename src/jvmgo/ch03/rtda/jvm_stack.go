package rtda

type Stack struct {
	maxSize uint   //最多可以容纳所少帧
	size    uint   //当前有多少帧
	_top    *Frame //栈顶指针
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("Java.lang.StackOverFlowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	self.size--

	return top
}
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
func (self *Stack) isEmpty()bool {
	return self._top==nil
}
