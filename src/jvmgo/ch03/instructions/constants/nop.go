package constants

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// just no nothing
}

