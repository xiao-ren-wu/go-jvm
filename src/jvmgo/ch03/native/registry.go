package native

import "jvmgo/ch03/rtda"

/*
用来注册和查找本地方法

把本地方法定义成一个函数，参数是Frame结构体指针，没有返回值
这个参数就是本地方法的工作空间，也就是连接Java虚拟机和Java类库的桥梁


本地方法在class文件中没有Code属性，所以要给maxStack和maxLocals字段赋值
本地方法帧的操作数栈至少能容纳返回值，为了简化代码，暂时给maxStack字段赋值为
因为本地方法帧在局部变量表中只用来存放参数值，所以把argSlotCount赋值给maxLocals字段刚好
至于code字段，也就是本地方法的字节码，所以第一条指令都是0XFE，
第二条指令则是根据函数的返回值选择相应的返回指令
*/
type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNative" {
		return emptyNativeMethod
	}
	return nil
}
func emptyNativeMethod(frame *rtda.Frame) {
	//do nothing
}
