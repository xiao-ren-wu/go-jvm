package heap

import "jvmgo/ch03/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *RunConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
