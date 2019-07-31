package classfile

/*
	cp_info{
		u1	tag;
		u1	info[];
	}
*/
//tag 常量定义
const (
	CONSTANT_Class					= 7
	CONSTANT_Fieldref				= 9
	CONSTANT_Methodref				= 10
	CONSTANT_InterfaceMethodref		= 11
	CONSTANT_String					= 8
	CONSTANT_Integer				= 3
	CONSTANT_Float					= 4
	CONSTANT_Long					= 5
	CONSTANT_Double					= 6
	CONSTANT_NameAndType			= 12
	CONSTANT_Utf8					= 1
	CONSTANT_MethodHandle			= 15
	CONSTANT_MethodType				= 16
	CONSTANT_InvokeDynamic			= 18
)
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

//读取常量池信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag:=reader.readUnit8()
	c:=newConstantInfo(tag,cp)
	c.readInfo(reader)
	return c
}

//根据tag创建具体的常量
func newConstantInfo(tag utf8, cp ConstantPool) ConstantInfo {
	switch tag{
	case CONSTANT_Integer:return &ConstantIntegerInfo{}
	case
	}
}
















