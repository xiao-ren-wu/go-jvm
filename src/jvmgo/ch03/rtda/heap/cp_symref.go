package heap

type SymRef struct {
	cp        *RunConstantPool //存放运行时常量池指针
	className string           //类的完全限定名
	class     *Class           //解析后类结构体指针
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
