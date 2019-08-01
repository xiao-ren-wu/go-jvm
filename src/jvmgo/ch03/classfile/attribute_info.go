package classfile

/*
attribute_info{
	u2	attribute_name_index;
	u4 	attribute_length;
	u1	info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUnit16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

/*
读取单个属性
先读取属性名索引，根据它从常量池中找到属性名，
然后读取属性长度，接着调用newAttributeInfo函数创建具体的属性实例
*/
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUnit16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUnit32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, pool ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &ExceptionsAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyncheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
