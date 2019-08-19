package heap

import "jvmgo/ch03/classfile"

type InterfaceMethodRef struct {
	MethodRef
	method *Method
}

func newInterfaceMethodRef(cp *RunConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
