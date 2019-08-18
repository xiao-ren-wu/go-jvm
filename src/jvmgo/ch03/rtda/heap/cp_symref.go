package heap

type SymRef struct {
	cp        *RunConstantPool //存放运行时常量池指针
	className string //类的完全限定名
	class     *Class //解析后类结构体指针
}


