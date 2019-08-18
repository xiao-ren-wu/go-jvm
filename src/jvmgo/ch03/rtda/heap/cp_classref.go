package heap

import "jvmgo/ch03/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *RunConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
