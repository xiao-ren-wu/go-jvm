package lang

import (
	"jvmgo/ch03/native"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

func init() {
	native.Register("java/lang/String","intern","()Ljava/lang/String;",intern)
}

func intern(frame *rtda.Frame) {
	this:=frame.LocalVars().GetThis()
	interned:=heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
