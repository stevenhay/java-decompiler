package parser

import (
	"log"
	"os"

	"github.com/stevenhay/GVM/class"
)

type ClassReader struct {
	b *ByteReader

	class *class.ClassFile
}

func NewClassReader(file string) (*ClassReader, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return &ClassReader{
		b:     NewByteReader(b),
		class: &class.ClassFile{},
	}, nil
}

func (r *ClassReader) Parse() *class.ClassFile {
	r.class.Magic = r.b.Read(4)
	r.class.MinorVersion = r.b.ReadUint16()
	r.class.MajorVersion = r.b.ReadUint16()
	r.class.ConstantPoolCount = r.b.ReadUint16()
	r.class.ConstantPool = r.buildConstantPool()
	r.class.AccessFlags = r.b.ReadUint16()
	r.class.ThisClass = r.b.ReadUint16()
	r.class.SuperClass = r.b.ReadUint16()
	r.class.InterfaceCount = r.b.ReadUint16()
	r.class.Interfaces = r.readInterfaceIndexes()
	r.class.FieldCount = r.b.ReadUint16()
	r.class.Fields = r.readFieldInfo()
	r.class.MethodCount = r.b.ReadUint16()
	r.class.Methods = r.readMethodInfo()
	r.class.AttributeCount = r.b.ReadUint16()
	r.class.Attributes = r.readAttributeInfo(r.class.AttributeCount)
	return r.class
}

func (r *ClassReader) readFieldInfo() []*class.FieldInfo {
	fields := make([]*class.FieldInfo, r.class.FieldCount)
	for i := 0; i < int(r.class.FieldCount); i++ {
		accessFlags := r.b.ReadUint16()
		nameIndex := r.b.ReadUint16()

		cp := r.class.ConstantPool.Get(nameIndex)
		if cp.Tag != class.CONSTANT_Utf8 {
			log.Fatalf("failed to read field info (idx=%d), name_index does not point at CONSTANT_Utf8", i)
		}

		descriptorIndex := r.b.ReadUint16()
		cp = r.class.ConstantPool.Get(descriptorIndex)
		if cp.Tag != class.CONSTANT_Utf8 {
			log.Fatalf("failed to read field info (idx=%d), descriptor_index does not point at CONSTANT_Utf8", i)
		}

		attributeCount := r.b.ReadUint16()
		attributes := r.readAttributeInfo(attributeCount)

		fields[i] = &class.FieldInfo{
			AccessFlags:     accessFlags,
			NameIndex:       nameIndex,
			DescriptorIndex: descriptorIndex,
			AttributesCount: attributeCount,
			AttributeInfo:   attributes,
		}
	}
	return fields
}

func (r *ClassReader) readMethodInfo() []*class.MethodInfo {
	methods := make([]*class.MethodInfo, r.class.MethodCount)
	for i := 0; i < int(r.class.MethodCount); i++ {
		accessFlags := r.b.ReadUint16()
		nameIndex := r.b.ReadUint16()

		cp := r.class.ConstantPool.Get(nameIndex)
		if cp.Tag != class.CONSTANT_Utf8 {
			log.Fatalf("failed to read method info (idx=%d), name_index does not point at CONSTANT_Utf8", i)
		}

		descriptorIndex := r.b.ReadUint16()
		cp = r.class.ConstantPool.Get(descriptorIndex)
		if cp.Tag != class.CONSTANT_Utf8 {
			log.Fatalf("failed to read method info (idx=%d), descriptor_index does not point at CONSTANT_Utf8", i)
		}

		attributeCount := r.b.ReadUint16()
		attributeInfo := r.readAttributeInfo(attributeCount)

		methods[i] = &class.MethodInfo{
			AccessFlags:     accessFlags,
			NameIndex:       nameIndex,
			DescriptorIndex: descriptorIndex,
			AttributesCount: attributeCount,
			AttributeInfo:   attributeInfo,
		}
	}
	return methods
}
