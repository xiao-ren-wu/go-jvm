package classfile

import "fmt"

type ClassFile struct {
	//magic	uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

/**
将[]byte解析成ClassFile结构体
*/
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", err)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool()
	self.accessFlags = reader.readUnit16()
	self.thisClass = reader.readUnit16()
	self.superClass = reader.readUnit16()
	self.interfaces = reader.readUnit16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

/**
魔数检查
*/
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUnit32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

/*
	M.m
	java 在1.2之前用过次版本号，主版本号为45
	1.2之后，没用过次版本号每次发布都是主版本号+1
*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUnit16()
	self.majorVersion = reader.readUnit16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
func (self *ClassFile) MinorVersion() uint16 {

}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {

}

//访问标志：返回该class文件是接口还是类，访问级别是public还是private
func (self *ClassFile) AccessFlags() uint16 {

}
func (self *ClassFile) Fields() []*MemberInfo {

}
func (self *ClassFile) Methods() []*MemberInfo {

}
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	//只有Java.lang.Object没有超类
	return ""
}
func (self *ClassFile) InterfacesNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
