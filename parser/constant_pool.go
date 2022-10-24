package parser

import (
	"log"

	"github.com/stevenhay/GVM/class"
)

func (r *ClassReader) buildConstantPool() *class.ConstantPool {
	poolItems := make([]*class.ConstantPoolItem, r.class.ConstantPoolCount)
	for i := 0; i < int(r.class.ConstantPoolCount)-1; i++ {
		// Each entry in the constant_pool table must begin with a 1-byte tag
		// indicating the kind of constant denoted by the entry.
		tag := r.b.ReadUint8()

		// Each tag byte must be followed by two or more bytes giving information
		// about the specific constant. The format of the additional information
		// depends on the tag byte, that is, the content of the info array varies
		// with the value of tag.
		var info interface{}

		switch tag {
		case class.CONSTANT_Utf8:
			info = r.readUtf8Info()
		case class.CONSTANT_Class:
			info = r.readClassInfo()
		case class.CONSTANT_String:
			info = r.readStringInfo()
		case class.CONSTANT_Fieldref:
			info = r.readFieldrefInfo()
		case class.CONSTANT_Methodref:
			info = r.readMethodrefInfo()
		case class.CONSTANT_NameAndType:
			info = r.readNameAndTypeInfo()
		default:
			log.Fatalf("unexpected tag - %d", tag)
		}

		poolItems[i] = &class.ConstantPoolItem{
			Tag:  tag,
			Info: info,
		}
	}

	return &class.ConstantPool{Items: poolItems}
}

func (r *ClassReader) readUtf8Info() *class.CONSTANT_Utf8_info {
	length := r.b.ReadUint16()
	return &class.CONSTANT_Utf8_info{
		Length: length,
		Bytes:  r.b.Read(int(length)),
	}
}

func (r *ClassReader) readClassInfo() *class.CONSTANT_Class_info {
	return &class.CONSTANT_Class_info{
		NameIndex: r.b.ReadUint16(),
	}
}

func (r *ClassReader) readStringInfo() *class.CONSTANT_String_info {
	return &class.CONSTANT_String_info{
		StringIndex: r.b.ReadUint16(),
	}
}

func (r *ClassReader) readFieldrefInfo() *class.CONSTANT_Fieldref_info {
	return &class.CONSTANT_Fieldref_info{
		ClassIndex:       r.b.ReadUint16(),
		NameAndTypeIndex: r.b.ReadUint16(),
	}
}

func (r *ClassReader) readMethodrefInfo() *class.CONSTANT_Methodref_info {
	return &class.CONSTANT_Methodref_info{
		ClassIndex:       r.b.ReadUint16(),
		NameAndTypeIndex: r.b.ReadUint16(),
	}
}

func (r *ClassReader) readInterfaceMethodrefInfo() *class.CONSTANT_InterfaceMethodref_info {
	return &class.CONSTANT_InterfaceMethodref_info{
		ClassIndex:       r.b.ReadUint16(),
		NameAndTypeIndex: r.b.ReadUint16(),
	}
}

func (r *ClassReader) readIntegerInfo() *class.CONSTANT_Integer_info {
	return &class.CONSTANT_Integer_info{
		Bytes: r.b.ReadUint32(),
	}
}

func (r *ClassReader) readFloatInfo() *class.CONSTANT_Float_info {
	return &class.CONSTANT_Float_info{
		Bytes: r.b.ReadUint32(),
	}
}

func (r *ClassReader) readLongInfo() *class.CONSTANT_Long_info {
	return &class.CONSTANT_Long_info{
		HighBytes: r.b.ReadUint32(),
		LowBytes:  r.b.ReadUint32(),
	}
}

func (r *ClassReader) readDoubleInfo() *class.CONSTANT_Double_info {
	return &class.CONSTANT_Double_info{
		HighBytes: r.b.ReadUint32(),
		LowBytes:  r.b.ReadUint32(),
	}
}

func (r *ClassReader) readNameAndTypeInfo() *class.CONSTANT_NameAndType_info {
	return &class.CONSTANT_NameAndType_info{
		NameIndex:       r.b.ReadUint16(),
		DescriptorIndex: r.b.ReadUint16(),
	}
}
