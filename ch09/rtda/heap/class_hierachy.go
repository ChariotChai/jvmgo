package heap

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}

	if !s.IsArray() {
		if !t.IsInterface() {
			if !t.IsInterface() {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return t.isJlObject() //todo
			} else {
				return t.IsSuperInterfaceOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isJlCloneable() || t.isJlSerializable() //数组可以强制转换成Object类型、Cloneable类型、Serializable类型
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}
	return false
}

func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) IsImplements(other *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == other || i.IsSubClassOf(other) {
				return true
			}
		}
	}
	return false
}

func (self *Class) IsSubInterfaceOf(other *Class) bool {
	for _, i := range self.interfaces {
		if i == other || i.IsSubInterfaceOf(other) {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

func (self *Class) IsSuperInterfaceOf(other *Class) bool {
	return other.IsSubInterfaceOf(self)
}
