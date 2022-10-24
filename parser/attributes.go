package parser

import (
	"log"

	"github.com/stevenhay/GVM/class"
)

func (r *ClassReader) readAttributeInfo(count uint16) []*class.AttributeInfo {
	attributes := make([]*class.AttributeInfo, count)
	for i := 0; i < int(count); i++ {
		nameIndex := r.b.ReadUint16()
		attributesLength := r.b.ReadUint32()

		cp := r.class.ConstantPool.Get(nameIndex)
		if cp.Tag != class.CONSTANT_Utf8 {
			log.Fatalf("failed to read attribute info (idx=%d), name_index does not point at CONSTANT_Utf8", i)
		}

		attributeName := string((cp.Info.(*class.CONSTANT_Utf8_info)).Bytes)
		ai := &class.AttributeInfo{
			AttributeNameIndex: nameIndex,
			AttributeLength:    attributesLength,
		}

		switch attributeName {
		case class.Code:
			ai.Info = r.readCodeAttribute()
		case class.LineNumberTable:
			ai.Info = r.readLineNumberTableAttribute()
		case class.SourceFile:
			ai.Info = r.readSourceFileAttribute()
		default:
			log.Fatalf("Unexpected attribute name %s, length = %d", attributeName, attributesLength)
		}

		attributes[i] = ai
	}
	return attributes
}

func (r *ClassReader) readCodeAttribute() *class.Code_attribute {
	c := &class.Code_attribute{
		MaxStack:   r.b.ReadUint16(),
		MaxLocals:  r.b.ReadUint16(),
		CodeLength: r.b.ReadUint32(),
	}
	c.Code = r.b.Read(int(c.CodeLength))
	c.ExceptionTableLength = r.b.ReadUint16()
	c.ExceptionTables = r.readExceptionTables(c.ExceptionTableLength)

	c.AttributesCount = r.b.ReadUint16()
	c.AttributeInfo = r.readAttributeInfo(c.AttributesCount)
	return c
}

func (r *ClassReader) readLineNumberTableAttribute() *class.LineNumberTable_attribute {
	a := &class.LineNumberTable_attribute{
		LineNumberTableLength: r.b.ReadUint16(),
	}

	a.LineNumberTable = make([]*class.LineNumberTableData, a.LineNumberTableLength)
	for i := 0; i < int(a.LineNumberTableLength); i++ {
		data := &class.LineNumberTableData{
			StartPc:    r.b.ReadUint16(),
			LineNumber: r.b.ReadUint16(),
		}
		a.LineNumberTable[i] = data
	}
	return a
}

func (r *ClassReader) readSourceFileAttribute() *class.SourceFile_attribute {
	return &class.SourceFile_attribute{
		SourceFileIndex: r.b.ReadUint16(),
	}
}

func (r *ClassReader) readExceptionTables(length uint16) []*class.ExceptionTable {
	exceptionTables := make([]*class.ExceptionTable, length)
	for i := 0; i < int(length); i++ {
		exceptionTables[i] = &class.ExceptionTable{
			StartPc:   r.b.ReadUint16(),
			EndPc:     r.b.ReadUint16(),
			HandlerPc: r.b.ReadUint16(),
			CatchType: r.b.ReadUint16(),
		}
	}
	return exceptionTables
}
