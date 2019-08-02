package classfile

/*
LineNumberTable属性表用于存放方法的行号信息
LocalVariableTable属性表用于存放方法的局部变量信息
这两种属性在前面介绍的SourceFile属性都属于调试信息，都不是运行时必须的，在使用javac进行编译Java程序的时候
默认会在class文件中生成这些信息。
LineNumberTable_attribute{
	u2	attribute_name_index;
	u4	attribute_length;
	{
		u2	start_pc;
		u2	line_number;
	}
	line_number_table[line_number_table_length];
}
 */
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc 	uint16
	lineNumber 	uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	LineNumberTableLength:=reader.readUnit16()
	self.lineNumberTable = make([] *LineNumberTableEntry,LineNumberTableLength)
	for i:=range self.lineNumberTable{
		self.lineNumberTable[i]=&LineNumberTableEntry{
			startPc:		reader.readUnit16(),
			lineNumber:		reader.readUnit16(),
		}
	}
}