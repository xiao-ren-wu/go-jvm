package references

import (
	"fmt"
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass:=frame.Method().Class()
	cp:=currentClass.ConstantPool()
	methodRef:=cp.GetConstant(self.Index).(*heap.MethodRef)
	resolveMethod:=methodRef.ResolvedMethod()
	if resolveMethod.IsStatic(){
		panic("java.lang.incompatibleClassError")
	}
	ref:=frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount()-1)
	if ref==nil{
		panic("java.lang.NullPointerException")
	}
	if resolveMethod.IsProtected()&&resolveMethod.Class().IsSuperClassOf(currentClass)&&resolveMethod.Class().GetPackageName()&&!ref.Class().IsSubClassOf(currentClass){
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked:=heap.LookupMethodInClass(ref.Class(),methodRef.Name(),methodRef.Descriptor())

	if methodToBeInvoked==nil||methodToBeInvoked.IsAbstract(){
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame,methodToBeInvoked)

	ref:=frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount()-1)
	if ref!=nil{
		//hack!
		if methodRef.Name()=="println"{
			_println(frame.OperandStack(),methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointException")
	}
	//...
}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(B)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(I)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	default:
		panic("println:"+descriptor)
	}
	stack.PopRef()
}
