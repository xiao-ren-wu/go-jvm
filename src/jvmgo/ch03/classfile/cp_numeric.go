package classfile

import "math"

/*
	CONSTANT_Integer_Info{
		u1	tag;
		u4	bytes;
	}
*/
type ConstantIntegerInfo struct {
	val int32
}

/*
	CONSTANT_Float_Info{
		u1	tag;
		u4	bytes;
	}
*/
type ConstantFloatInfo struct {
	val float32
}

/*
	CONSTANT_Long_Info{
		u1	tag;
		u4	high_bytes;
		u4	low_bytes;
	}
*/
type ConstantLongInfo struct {
	val int64
}

/*
	CONSTANT_Double_Info{
		u1	tag;
		u4	high_bytes;
		u4	low_bytes;
	}
*/
type ConstantDoubleInfo struct {
	val float64
}

/*
	u1	tag;
	u2	length;
	u1	bytes[length]
*/
type ConstantUtf8Info struct {
	str string
}

/*
	CONSTANT_String_Info{
		u1	tag;
		u2	string_index;
	}
	本身不存放字符串数据，只存了常量池索引，这个索引指向一个CONSTANT_Utf8_Info常量
*/
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

/*
	CONSTANT_Class_info{
		u1	tag;
		u2	name_index;
	}
*/
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

/*
给出字段和方法的描述符
	CONSTANT_NameAndType_info{
		u1	tag;
		u2	name_index;
		u2	descriptor_index;
	}
CONSTANT_Class_info和CONSTANT_NameAndType_info加在一起可以确定一个字段或者方法
字段或者方法名由name_index给出，
字段或方法描述符由descriptor_index给出
Java虚拟机定义了一种简单的语法来确定描述字段和方法：
1. 类型描述符
	1.1基本类型byte,short,char,int,long,float,double的描述符是单个字母，分别对应B,S,C,I,J,F和D
	1.2引用类型描述符是L+类的完全限定名+分号
	1.3数组类型描述符是[+数组元素类型描述符
2.  字段描述符就是字段类型的描述符
3.	方法描述符是(分号分隔的参数类型描述符)+返回值类型描述符，其中void返回值由单个字母V表示
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}
func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}
func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}
func (self *ConstantLongInfo) Value() int64 {
	return self.val
}
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMTF8(bytes)
}
func decodeMTF8(bytes []byte) string {
	return string(bytes)
}
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
