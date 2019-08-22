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
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.ResolveFieldRef()
	}
	return self.field
}
func (self *FieldRef) ResolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookUpField(c, self.name, self.descriptor)
	if field == nil {
		panic("Java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("Java.lang.IllegalAccessError")
	}
	self.field = field
}
func lookUpField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookUpField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookUpField(c.superClass, name, descriptor)
	}
	return nil
}
