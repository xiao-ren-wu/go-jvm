package heap

/**
boolean类型的数组也是用int8来表示
*/
func (self *Object) Bytes() []int8 {
	return self.data.([]int8)
}
func (self *Object) Shorts() []int16 {
	return self.data.([]int16)
}
func (self *Object) Ints() []int32 {
	return self.data.([]int32)
}
func (self *Object) Longs() []int64 {
	return self.data.([]int64)
}
func (self *Object) Chars() []int16 {
	return self.data.([]int16)
}
func (self *Object) Floats() []float32 {
	return self.data.([]float32)
}
func (self *Object) Doubles() []float64 {
	return self.data.([]float64)
}
func (self *Object) Refs() []*Object {
	return self.data.([]*Object)
}
func (self *Object) ArrayLength() int32 {
	switch self.data.(type) {
	case []int8:
		return int32(len(self.data.([]int8)))
	case []int16:
		return int32(len(self.data.([]int16)))
	case []int32:
		return int32(len(self.data.([]int32)))
	case []int64:
		return int32(len(self.data.([]int64)))
	case []uint16:
		return int32(len(self.data.([]uint16)))
	case []float32:
		return int32(len(self.data.([]float32)))
	case []float64:
		return int32(len(self.data.([]float64)))
	case []*Object:
		return int32(len(self.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos+length]
		_dst := dst.data.([]int32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos+length]
		_dst := dst.data.([]*Object)[dstPos : dstPos+length]
		copy(_dst, _src)
		//todo others ...

	}
}
