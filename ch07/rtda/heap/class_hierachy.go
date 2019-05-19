package heap

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
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
