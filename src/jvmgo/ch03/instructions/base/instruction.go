package base

import "jvmgo/ch03/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

//无操作数指令
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//just do nothing ~
}
//跳转指令
type BranchInstruction struct {
	Offset 	int //跳转偏移量
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}
/**
存储和加载指令需要根据索引存取局部变量表，索引由单字节操作数给出
Index表示局部变量表索引
 */
type Index8Instruction struct {
	Index 	uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

/**
有些指令需要访问常量池，常量池索引由两字节操作数给出
Index表示常量池索引
 */
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
















