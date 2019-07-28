package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUnit8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUnit16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUnit32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUnit64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/**
readUnit16s读取uint16表，表的大小由开头的uint16数据指出
*/
func (self *ClassReader) readUnit16s() []uint16 {
	n := self.readUnit16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUnit16()
	}
	return s
}

/**
用于读取指定数量的字节
*/
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[:n]
	return bytes
}
