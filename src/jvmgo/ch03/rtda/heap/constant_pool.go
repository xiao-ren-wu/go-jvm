package heap

import (
	"fmt"
	"jvmgo/ch03/classfile"
)

/**
运行时常量池：
	字面量：
		整数、浮点数、字符串字面量
	符号引用：
		类符号引用、字段符号引用、方法符号引用、接口符号引用
*/

type RunConstant interface {
}

type RunConstantPool struct {
	class  *Class
	consts []RunConstant
}

/*
将class文件中的常量池转换成运行时常量池
*/
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *RunConstantPool {
	cpCount := len(cfCp)
	consts := make([]RunConstant, cpCount)
	rtCp := &RunConstantPool{class: class, consts: consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldRefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldRefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodRefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodRefInfo)
		}
	}
	return rtCp
}

/*
根据索引返回常量
*/
func (self *RunConstantPool) GetConstant(index uint) RunConstant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constant at index %d", index))
}
