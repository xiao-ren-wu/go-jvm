package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser:=&MethodDescriptorParser{descriptor:descriptor}
	return parser.parse()
}

