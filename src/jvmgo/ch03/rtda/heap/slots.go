package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}
type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount == 0 {
		return nil
	}
	return make([]Slot, slotCount)
}

/*
存储int
对于boolean,byte,short,char类型，可以直接使用int值来处理
*/
func (self Slots) SetInt(index uint, val int32) {
	self[index].num = val
}
func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

/*
float类型可以先转换成int类型，然后按照int变量来处理
*/
func (self Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

/*
long 类型需要拆成两个int变量
*/
func (self Slots) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

/*
double变量可以先转换成long,然后按照long变量来处理
*/
func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

/*
引用类型直接存取
*/
func (self Slots) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}
func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
}
