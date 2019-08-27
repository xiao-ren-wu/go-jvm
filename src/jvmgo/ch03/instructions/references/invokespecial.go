package references

import (
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}
/**
1.获取当前类，当前常量池，方法符号引用，然后解析符号引用，拿到解析后的类和方法
2.假定从当前方法符号引用中解析出来的是类C，方法是M，如果M是构造方法，则声明M的类必须是C，否则抛出NoSuchMethodError异常，
	如果M是静态方法，则抛出IncompatibleClassChangeError
3.从操作数栈中弹出this引用，如果该引用是null,抛出NullPointerException异常，注意，在传递参数之前不能破坏操作数栈的状态。

 */
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame)  {
	currentClass:=frame.Method().Class()
	cp:=currentClass.ConstantPool()
	methodRef:=cp.GetConstant(self.Index).(*heap.MethodRef)
	resolveClass:=methodRef.ResolvedClass()
	resolveMethod:=methodRef.ResolvedMethod()

	if resolveMethod.Name()=="<init>"&&resolveMethod.Class()!=resolveClass{
		panic("java.lang.NoSuchMethodError")
	}
	if resolveMethod.IsStatic(){
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref:=frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount())
	if ref==nil{
		panic("java.lang.NullPointer.Exception")
	}

	if resolveMethod.IsProtected() &&resolveMethod.Class().IsSuperClassOf(currentClass) &&resolveMethod.Class().GetPackageName()!=currentClass.GetPackageName() &&ref.Class()!=currentClass &&!ref.Class().IsSubClassOf(currentClass){
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked:=resolveMethod
	if currentClass.IsSuper()&&resolveClass.IsSuperClassOf(currentClass)&&resolveMethod.Name()!="<init>"{
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),methodRef.Name(),methodRef.Descriptor())
	}

	if methodToBeInvoked==nil||methodToBeInvoked.IsAbstract(){
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame,methodToBeInvoked)

}

