package classfile

/**
	CONSTANT_Fieldref_info----------------字段符号引用
	CONSTANT_Methodref_info---------------普通（非接口）方法引用
	CONSTANT_InterfaceMethodref_info------接口方法引用

	eg:
	CONSTANT_Fieldref_info{
		u1		tag;
		u2		class_index;
		u2		name_and_type_index;
	}
 */
type ConstantMemberrefInfo struct {
	cp 					ConstantPool
	classIndex			uint16
	nameAndTypeIndex	uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.nameAndTypeIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}























