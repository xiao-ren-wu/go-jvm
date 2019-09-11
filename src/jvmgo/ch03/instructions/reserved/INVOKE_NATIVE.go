package reserved

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/native"
	"jvmgo/ch03/rtda"
	_ "jvmgo/ch03/native/java/lang"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError:" + methodInfo)
	}
	nativeMethod(frame)
}
