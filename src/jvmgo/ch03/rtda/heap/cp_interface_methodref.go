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
func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}
func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassError")
	}
	method := lookupInterfaceMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError.")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}
func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return LookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
