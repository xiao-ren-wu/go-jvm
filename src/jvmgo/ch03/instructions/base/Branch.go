package base

import "jvmgo/ch03/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc:=frame.Thread().PC()
	nextPC:=pc+offset
	frame.SetNextPC(nextPC)
}
