package heap

import "jvmgo/ch03/classfile"

type ClassMember struct {
	AccessFlags
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}
func (self *ClassMember) Class() *Class  {
	return self.class
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
func (self *ClassMember)Name() string  {
	return self.name
}
