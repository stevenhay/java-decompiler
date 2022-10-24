package class

var CLASS_ACCESS_FLAGS = map[string]uint16{
	"ACC_PUBLIC":     1 << 0,
	"ACC_FINAL":      1 << 4,
	"ACC_SUPER":      1 << 5,
	"ACC_INTERFACE":  1 << 9,
	"ACC_ABSTRACT":   1 << 10,
	"ACC_SYNTHETIC":  1 << 12,
	"ACC_ANNOTATION": 1 << 13,
	"ACC_ENUM":       1 << 14,
}

type ClassFile struct {
	Magic             []byte
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPool      *ConstantPool
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	InterfaceCount    uint16
	Interfaces        []uint16
	FieldCount        uint16
	Fields            []*FieldInfo
	MethodCount       uint16
	Methods           []*MethodInfo
	AttributeCount    uint16
	Attributes        []*AttributeInfo
}

type ConstantPool struct {
	Items []*ConstantPoolItem
}

func (c *ConstantPool) Get(index uint16) *ConstantPoolItem {
	return c.Items[index-1]
}

func (c *ConstantPool) GetString(index uint16) string {
	return string(c.Items[index-1].Info.(*CONSTANT_Utf8_info).Bytes)
}

type ConstantPoolItem struct {
	Tag  uint8
	Info interface{}
}

func (c *ClassFile) GetClassAccessFlags() []string {
	var flags []string
	for k, v := range CLASS_ACCESS_FLAGS {
		if (c.AccessFlags & v) > 0 {
			flags = append(flags, k)
		}
	}
	return flags
}
