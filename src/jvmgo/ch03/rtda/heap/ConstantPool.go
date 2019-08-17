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
func newConstantPool(class *Class,cfCp classfile.ConstantPool) *RunConstantPool {
	cpCount:=len(cfCp)
	consts:=make([]RunConstant,cpCount)
	rtCp:=&RunConstantPool{class: class,consts:consts}
	for i:=1;i<cpCount;i++{
		cpInfo:=cfCp[i]
		switch cpInfo.(type) {

		}
	}
	return rtCp
}
func (self *RunConstantPool) GetConstant(index uint) RunConstant {
	if c:=self.consts[index];c!=nil{
		return c
	}
	panic(fmt.Sprintf("No constant at index %d",index))
}
