package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

//类D通过符号引用N引用类C，解析N，先用D类的加载器加载C，然后检查D是否有权限访问C
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
