package control

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

/*
return void
*/
type RETURN struct {
	base.NoOperandsInstruction
}

/**
return reference
*/
type ARETURN struct {
	base.NoOperandsInstruction
}

/**
return double
*/
type DRETURN struct {
	base.NoOperandsInstruction
}

/**
return float
*/
type FRETURN struct {
	base.NoOperandsInstruction
}

/**
return int
*/
type IRETURN struct {
	base.NoOperandsInstruction
}

/**
return long
*/
type LRETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokeFrame.OperandStack().PushInt(retVal)
}
func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}
func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}
func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}























