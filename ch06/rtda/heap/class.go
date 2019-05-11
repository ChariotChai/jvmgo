package heap

import (
	"jvmgo/ch06/classfile"
	"strings"
)

type Class struct {
	accessFlags     uint16
	name            string
	superClassName  string
	interfaceNames  []string
	constantPool    *ConstantPool
	fields          []*Field
	methods         []*Method
	loader          *ClassLoader
	superClass      *Class
	interfaces      []*Class
	instanceSlotCnt uint
	staticSlotCnt   uint
	staticVars      *Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.accessFlags
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
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

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCnt),
	}
}

func (self *Class) StaticVars() *Slots {
	return self.staticVars
}
