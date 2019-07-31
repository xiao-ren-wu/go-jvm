package classfile

/**
字段或方法
 */
type MemberInfo struct {
	cp 					ConstantPool //常量池
	accessFlags			uint16	//访问标志
	nameIndex			uint16	//字段或者方法名的索引
	descriptorIndex		uint16	//字段或者方法的描述
	attributes 			[]AttributeInfo	//属性表
}

/**
	读取字段表或者方法表
 */
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount:=reader.readUnit16()
	members:=make([]*MemberInfo,memberCount)
	for i:=range members{
		members[i]=readMember(reader,cp)
	}
	return members
}

/**
表示读取字段或者方法数据
 */
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:					cp,
		accessFlags:		reader.readUnit16(),
		nameIndex: 			reader.readUnit16(),
		descriptorIndex: 	reader.readUnit16(),
		attributes:			readAttributes(reader,cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
/**
从常量池中查找方法名或者字段名
 */
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

/**
从常量池中查找方法或者字段描述符
 */
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

