package class

const (
	CONSTANT_Utf8               = 1
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_Class              = 7
	CONSTANT_String             = 8
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_NameAndType        = 12
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_Dynamic            = 17
	CONSTANT_InvokeDynamic      = 18
	CONSTANT_Module             = 19
	CONSTANT_Package            = 20
)

//	CONSTANT_Class_info {
//	    u1 tag;
//	    u2 name_index;
//	}
type CONSTANT_Class_info struct {
	// The value of the name_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	NameIndex uint16
}

//	CONSTANT_Fieldref_info {
//	    u1 tag;
//	    u2 class_index;
//	    u2 name_and_type_index;
//	}
type CONSTANT_Fieldref_info struct {
	// In a CONSTANT_Fieldref_info structure, the class_index item may be either a class type or an interface type.
	ClassIndex uint16
	// The value of the name_and_type_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_NameAndType_info structure.
	// This constant_pool entry indicates the name and descriptor of the field or method.
	NameAndTypeIndex uint16
}

//	CONSTANT_Methodref_info {
//	    u1 tag;
//	    u2 class_index;
//	    u2 name_and_type_index;
//	}
type CONSTANT_Methodref_info struct {
	// In a CONSTANT_Methodref_info structure, the class_index item must be a class type, not an interface type.
	ClassIndex uint16
	// The value of the name_and_type_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_NameAndType_info structure.
	// This constant_pool entry indicates the name and descriptor of the field or method.
	NameAndTypeIndex uint16
}

//	CONSTANT_InterfaceMethodref_info {
//	    u1 tag;
//	    u2 class_index;
//	    u2 name_and_type_index;
//	}
type CONSTANT_InterfaceMethodref_info struct {
	// In a CONSTANT_InterfaceMethodref_info structure, the class_index item must be an interface type, not a class type.
	ClassIndex uint16
	// The value of the name_and_type_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_NameAndType_info structure.
	// This constant_pool entry indicates the name and descriptor of the field or method.
	NameAndTypeIndex uint16
}

//	CONSTANT_String_info {
//	    u1 tag;
//	    u2 string_index;
//	}
type CONSTANT_String_info struct {
	// The value of the string_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	StringIndex uint16
}

//	CONSTANT_Integer_info {
//	    u1 tag;
//	    u4 bytes;
//	}
type CONSTANT_Integer_info struct {
	// The bytes item of the CONSTANT_Integer_info structure represents the value of the int constant.
	// The bytes of the value are stored in big-endian (high byte first) order.
	Bytes uint32
}

//	CONSTANT_Float_info {
//	    u1 tag;
//	    u4 bytes;
//	}
type CONSTANT_Float_info struct {
	// The bytes item of the CONSTANT_Float_info structure represents the value of the float constant
	// in IEEE 754 floating-point single format. The bytes of the single format representation are
	// stored in big-endian (high byte first) order.
	Bytes uint32
}

//	CONSTANT_Long_info {
//	    u1 tag;
//	    u4 high_bytes;
//	    u4 low_bytes;
//	}
//
// The unsigned high_bytes and low_bytes items of the CONSTANT_Long_info structure together represent
// the value of the long constant ((long) high_bytes << 32) + low_bytes where the bytes of each of
// high_bytes and low_bytes are stored in big-endian (high byte first) order.
type CONSTANT_Long_info struct {
	HighBytes uint32
	LowBytes  uint32
}

//	CONSTANT_Double_info {
//	    u1 tag;
//	    u4 high_bytes;
//	    u4 low_bytes;
//	}
//
// The high_bytes and low_bytes items of the CONSTANT_Double_info structure together represent the double
// value in IEEE 754 floating-point double format. The bytes of each item are stored in big-endian
// (high byte first) order.
type CONSTANT_Double_info struct {
	HighBytes uint32
	LowBytes  uint32
}

//	CONSTANT_NameAndType_info {
//	    u1 tag;
//	    u2 name_index;
//	    u2 descriptor_index;
//	}
type CONSTANT_NameAndType_info struct {
	// The value of the name_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	NameIndex uint16
	// The value of the descriptor_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	DescriptorIndex uint16
}

//	CONSTANT_Utf8_info {
//		u1 tag;
//		u2 length;
//		u1 bytes[length];
//	}
type CONSTANT_Utf8_info struct {
	// The value of the length item gives the number of bytes in the bytes array
	// (not the length of the resulting string).
	Length uint16
	// The bytes array contains the bytes of the string.
	// No byte may have the value (byte)0.
	// No byte may lie in the range (byte)0xf0 to (byte)0xff.
	Bytes []uint8
}

//	CONSTANT_MethodHandle_info {
//		u1 tag;
//		u1 reference_kind;
//		u2 reference_index;
//	}
type CONSTANT_MethodHandle_info struct {
	// The value of the reference_kind item must be in the range 1 to 9. The value denotes
	// the kind of this method handle, which characterizes its bytecode behavior
	ReferenceKind uint8
	// The value of the reference_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be as follows:
	//
	// If the value of the reference_kind item is 1 (REF_getField), 2 (REF_getStatic),
	// 3 (REF_putField), or 4 (REF_putStatic), then the constant_pool entry at that index must
	// be a CONSTANT_Fieldref_info structure
	//
	// If the value of the reference_kind item is 5 (REF_invokeVirtual) or 8 (REF_newInvokeSpecial),
	// then the constant_pool entry at that index must be a CONSTANT_Methodref_info structure
	//
	// If the value of the reference_kind item is 6 (REF_invokeStatic) or 7 (REF_invokeSpecial),
	// then if the class file version number is less than 52.0, the constant_pool entry at that index
	// must be a CONSTANT_Methodref_info structure representing a class's method for which a method
	// handle is to be created; if the class file version number is 52.0 or above, the constant_pool
	// entry at that index must be either a CONSTANT_Methodref_info structure or a CONSTANT_InterfaceMethodref_info
	// structure
	//
	// If the value of the reference_kind item is 9 (REF_invokeInterface), then the constant_pool
	//entry at that index must be a CONSTANT_InterfaceMethodref_info
	ReferenceIndex uint16
}

//	CONSTANT_MethodType_info {
//		u1 tag;
//		u2 descriptor_index;
//	}
type CONSTANT_MethodType_info struct {
	// The value of the descriptor_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	DescriptorIndex uint16
}

//	CONSTANT_Dynamic_info {
//		u1 tag;
//		u2 bootstrap_method_attr_index;
//		u2 name_and_type_index;
//	}
//
// The CONSTANT_Dynamic_info structure is used to represent a dynamically-computed constant,
// an arbitrary value that is produced by invocation of a bootstrap method in the course of an
// ldc instruction, among others. The auxiliary type specified by the structure constrains
// the type of the dynamically-computed constant.
type CONSTANT_Dyanmic_info struct {
	// The value of the bootstrap_method_attr_index item must be a valid index into the
	// bootstrap_methods array of the bootstrap method table of this class file
	BootstrapMethodAttrIndex uint16
	// The value of the name_and_type_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_NameAndType_info structure
	//
	// In a CONSTANT_Dynamic_info structure, the indicated descriptor must be a field descriptor
	NameAndTypeIndex uint16
}

//	CONSTANT_InvokeDynamic_info {
//		u1 tag;
//		u2 bootstrap_method_attr_index;
//		u2 name_and_type_index;
//	}
//
// The CONSTANT_InvokeDynamic_info structure is used to represent a dynamically-computed call site,
// an instance of java.lang.invoke.CallSite that is produced by invocation of a bootstrap method in
// the course of an invokedynamic instruction. The auxiliary type specified by the structure constrains
// the method type of the dynamically-computed call site.
type CONSTANT_InvokeDynamic_info struct {
	// The value of the bootstrap_method_attr_index item must be a valid index into the
	// bootstrap_methods array of the bootstrap method table of this class file
	BootstrapMethodAttrIndex uint16
	// The value of the name_and_type_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_NameAndType_info structure
	//
	// In a CONSTANT_InvokeDynamic_info structure, the indicated descriptor must be a method descriptor
	NameAndTypeIndex uint16
}

//	CONSTANT_Module_info {
//		u1 tag;
//		u2 name_index;
//	}
type CONSTANT_Module_info struct {
	// The value of the name_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	NameIndex uint16
}

//	CONSTANT_Package_info {
//		u1 tag;
//		u2 name_index;
//	}
type CONSTANT_Package_info struct {
	// The value of the name_index item must be a valid index into the constant_pool table.
	// The constant_pool entry at that index must be a CONSTANT_Utf8_info structure
	NameIndex uint16
}
