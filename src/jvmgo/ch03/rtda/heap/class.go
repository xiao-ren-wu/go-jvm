package heap

import (
	. "jvmgo/ch03/classfile"
	"strings"
)

/*
Class 结构体
*/
type Class struct {
	AccessFlags
	name              string //this class name
	superClassName    string
	interfaceNames    []string
	constantPool      *RunConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool    //class init flag
	jClass            *Object // java.lang.Class实例
}

func newClass(cf *ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfacesNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}
func (self *Class) NewObject() *Object {
	return newObject(self)
}
func (self *Class) ConstantPool() *RunConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name {
			return method
		}
	}
	return nil
}
func (self *Class) Name() string {
	return self.name
}
func (self *Class) InitStarted() bool {
	return self.initStarted
}
func (self *Class) StartInit() {
	self.initStarted = true
}
func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}

/**
根据字段名和描述符查找字段
*/
func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range self.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}
func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}
func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}
func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}
func (self *Class) GetInstanceField(name, descriptor string) *Field {
	return self.getField(name, descriptor, false)
}
