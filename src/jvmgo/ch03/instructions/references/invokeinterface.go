package references

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

/*
invokeinterface指令的操作码后面跟着4字节而非2字节
前两字节含义和其他指令相同，是个uint16运行时常量池索引
而第三个字节是给方法传递参数需要的slot数，其含义是给Method结构体定义的argCount字段相同
第四个字节的值必须是0，为了兼容某些Java虚拟机
*/
type INVOKE_INTERFACE struct {
	index uint
	//count uint8
	//zero uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8() //count
	reader.ReadUint8() //zero
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	/*
		从运行时常量池中拿到解析接口方法符号引用，如果解析后的方法是静态方法或者私有方法，则抛出异常
	 */
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//从操作数栈中弹出this引用，如果引用是null，则抛出NullPointerException异常
	//如果引用所指对象的类没有实现解析出来的接口，则抛出IncompatibleClassChangeError异常
	ref:=frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount()-1)
	if ref==nil{
		panic("java.lang.NullPointException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()){
		panic("java.lang.IncompatibleClassChangeError")
	}

	//查找最终调用方法，如果找不到，或者找到的方法是抽象的，抛出异常，如果找到的方法的访问权限不是public，抛出异常
	methodToBeInvoked:=heap.LookupMethodInClass(ref.Class(),methodRef.Name(),methodRef.Descriptor())
	if methodToBeInvoked==nil ||methodToBeInvoked.IsAbstract(){
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic(){
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame,methodToBeInvoked)
}
