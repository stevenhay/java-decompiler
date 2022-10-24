package class

var METHOD_ACCESS_FLAGS = map[string]uint16{
	"ACC_PUBLIC":       1 << 0,
	"ACC_PRIVATE":      1 << 1,
	"ACC_PROTECTED":    1 << 2,
	"ACC_STATIC":       1 << 3,
	"ACC_FINAL":        1 << 4,
	"ACC_SYNCHRONIZED": 1 << 5,
	"ACC_BRIDGE":       1 << 6,
	"ACC_VARARGS":      1 << 7,
	"ACC_NATIVE":       1 << 8,
	"ACC_ABSTRACT":     1 << 10,
	"ACC_STRICT":       1 << 11,
	"ACC_SYNTHETIC":    1 << 12,
}

type MethodInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	AttributeInfo   []*AttributeInfo
}

func (f *MethodInfo) GetName(constantPool *ConstantPool) string {
	return string(constantPool.Get(f.NameIndex).Info.(*CONSTANT_Utf8_info).Bytes)
}

func (f *MethodInfo) GetDescriptor(constantPool *ConstantPool) string {
	return string(constantPool.Get(f.DescriptorIndex).Info.(*CONSTANT_Utf8_info).Bytes)
}

func (f *MethodInfo) GetCodeAttributes(ConstantPool *ConstantPool) []*Code_attribute {
	ca := make([]*Code_attribute, 0)
	for _, a := range f.AttributeInfo {
		name := a.GetName(ConstantPool)
		if name == Code {
			ca = append(ca, a.Info.(*Code_attribute))
		}
	}
	return ca
}

func (f *MethodInfo) GetCode(constantPool *ConstantPool) []byte {
	for _, a := range f.AttributeInfo {
		name := a.GetName(constantPool)
		if name == Code {
			return a.GetAsCodeAttribute().Code
		}
	}
	return nil
}

func (m *MethodInfo) GetMethodAccessFlags() []string {
	var flags []string
	for k, v := range METHOD_ACCESS_FLAGS {
		if (m.AccessFlags & v) > 0 {
			flags = append(flags, k)
		}
	}
	return flags
}
