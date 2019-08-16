package rtda

import (
	"jvmgo/ch03/rtda/heap"
	"math"
)

/*
局部变量表：
	执行方法所需的局部变量表大小和操作数栈深度是由编译器提前计算好的
存储在class文件中method_info结构的Code属性中
	Java虚拟机规范中，
局部变量表的每个元素至少可以容纳一个int类型或者引用值，
			两个连续的元素至少可以容纳一个long或者double
*/
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

/*
存储int
对于boolean,byte,short,char类型，可以直接使用int值来处理
*/
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

/*
float类型可以先转换成int类型，然后按照int变量来处理
*/
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

/*
long 类型需要拆成两个int变量
*/
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

/*
double变量可以先转换成long,然后按照long变量来处理
*/
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

/*
引用类型直接存取
*/
func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}
