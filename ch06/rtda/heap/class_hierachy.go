package heap

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

func (self *Class) isSubClassOf(other *Class) bool {
	for c := self; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isImplements(other *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == other || i.isSubClassOf(other) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(other *Class) bool {
	for _, i := range self.interfaces {
		if i == other || i.isSubInterfaceOf(other) {
			return true
		}
	}
	return false
}
