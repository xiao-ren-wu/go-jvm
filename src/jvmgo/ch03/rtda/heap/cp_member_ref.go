package heap

import "jvmgo/ch03/classfile"

type MemberRef struct {
	SymRef
	name        string
	description string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.description = refInfo.NameAndDescriptor()
}

