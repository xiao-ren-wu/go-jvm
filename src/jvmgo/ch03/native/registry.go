package native

import "jvmgo/ch03/rtda"

/*
用来注册和查找本地方法

把本地方法定义成一个函数，参数是Frame结构体指针，没有返回值
这个参数就是本地方法的工作空间，也就是连接Java虚拟机和Java类库的桥梁
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
