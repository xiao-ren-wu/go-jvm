package heap

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if s.IsArray() {
		if !s.IsInterface() {
			if !t.IsInterface() {
				return s.isSubInterfaceOf(t)
			} else {
				return s.isImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return t.IsJ1Object()
			} else {
				return t.isSubInterfaceOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJ1Object()
			} else {
				return t.isJ1Cloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			ts := t.ComponentClass()
			return sc == ts || ts.isAssignableFrom(sc)
		}
	}
	return false
}

func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
