package class

var FIELD_ACCESS_FLAGS = map[string]uint16{
	"ACC_PUBLIC":    1 << 0,
	"ACC_PRIVATE":   1 << 1,
	"ACC_PROTECTED": 1 << 2,
	"ACC_STATIC":    1 << 3,
	"ACC_FINAL":     1 << 4,
	"ACC_VOLATILE":  1 << 6,
	"ACC_TRANSIENT": 1 << 7,
	"ACC_SYNTHETIC": 1 << 12,
	"ACC_ENUM":      1 << 14,
}

//	field_info {
//	    u2             access_flags;
//	    u2             name_index;
//	    u2             descriptor_index;
//	    u2             attributes_count;
//	    attribute_info attributes[attributes_count];
//	}
type FieldInfo struct {
	// The value of the access_flags item is a mask of
	// flags used to denote access permission to and properties of this field
	AccessFlags uint16
	// The value of the name_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	NameIndex uint16
	// The value of the descriptor_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	DescriptorIndex uint16
	// The value of the attributes_count item indicates the number of additional attributes of this field.
	AttributesCount uint16
	// Each value of the attributes table must be an attribute_info structure. A field can have any number
	// of optional attributes associated with it.
	AttributeInfo []*AttributeInfo
}

func (f *FieldInfo) GetName(constantPool *ConstantPool) string {
	return string(constantPool.Get(f.NameIndex).Info.(*CONSTANT_Utf8_info).Bytes)
}

func (f *FieldInfo) GetFieldAccessFlags() []string {
	var flags []string
	for k, v := range FIELD_ACCESS_FLAGS {
		if (f.AccessFlags & v) > 0 {
			flags = append(flags, k)
		}
	}
	return flags
}
