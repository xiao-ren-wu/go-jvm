package heap

import "jvmgo/ch03/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *RunConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

