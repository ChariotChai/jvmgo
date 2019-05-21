package heap

import (
	"jvmgo/ch08/classfile"
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
	staticVars      Slots
	initStarted     bool
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) IsAccessibleTo(c *Class) bool {
	return self.IsPublic() || self.GetPackageName() == c.GetPackageName()
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (c *Class) InitStarted() bool {
	return c.initStarted
}

func (c *Class) StartInit() {
	c.initStarted = true
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (c *Class) Loader() *ClassLoader {
	return c.loader
}

func (c *Class) ArrayClass() *Class {
	arrClassName := getArrayClassName(c.name)
	return c.loader.LoadClass(arrClassName)
}

func (c *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(c.name)
	return c.loader.LoadClass(componentClassName)
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (c *Class) isJlObject() bool {
	return c.name == "java/lang/Object"
}

func (c *Class) isJlCloneable() bool {
	return c.name == "java/lang/Cloneable"
}

func (c *Class) isJlSerializable() bool {
	return c.name == "java/io/Serializable"
}
