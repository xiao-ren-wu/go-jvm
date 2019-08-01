package classfile
/*
用于表示常量表达式的值
ConstantValue_attribute{
	u2	attribute_name_index;
	u4	attribute_length;
	u2	constantvalue_index;
}
 */
type ConstantValueAttribute struct {
	constantValueIndex	uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUnit16()
}
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}

