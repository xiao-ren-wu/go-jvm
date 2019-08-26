package heap

func LookupMethodInClass(class *Class, descriptor, name string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func LookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := LookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
