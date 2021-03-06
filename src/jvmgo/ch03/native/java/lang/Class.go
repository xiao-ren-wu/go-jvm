package lang

import (
	"jvmgo/ch03/native"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()java/lang/String;", getName0)
	native.Register("java/lang/Class", "desireAssertionStatus0", "(Ljava/lang/Class;)Z", desireAssertionStatus0)
}

//static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}
func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

//private static native boolean desiredAssertionStatus0(Class<?> clazz);
//do nothing
func desireAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}
