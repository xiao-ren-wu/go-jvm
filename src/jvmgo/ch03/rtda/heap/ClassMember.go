package heap

import "jvmgo/ch03/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	description string
	class       *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.description = memberInfo.Descriptor()
}
