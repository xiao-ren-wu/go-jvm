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

func (self *MethodRef) resolveMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError.")
	}
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError.")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}
func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = LookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}

func (self *MethodRef) ResolvedMethod() *Method {
	return self.method
}
