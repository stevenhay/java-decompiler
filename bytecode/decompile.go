package bytecode

import (
	"fmt"
	"log"

	"github.com/stevenhay/GVM/class"
	"github.com/stevenhay/GVM/parser"
)

func decompileNEWInstruction(op Opcode, buf *parser.ByteReader, constantPool *class.ConstantPool) string {
	data := buf.Read(int(op.Length))
	index := (data[0] << 8) | data[1]

	poolItem := constantPool.Get(uint16(index))
	ref := poolItem.Info.(*class.CONSTANT_Class_info)
	name := constantPool.GetString(ref.NameIndex)

	return fmt.Sprintf("%s %s", op.Name, name)
}

func decompileBIPUSHInstruction(op Opcode, buf *parser.ByteReader, constantPool *class.ConstantPool) string {
	data := buf.Read(int(op.Length))[0]
	return fmt.Sprintf("%s %d", op.Name, data)
}

func decompileALOADInstruction(op Opcode, buf *parser.ByteReader, constantPool *class.ConstantPool) string {
	if op.Op == ALOAD_0 || op.Op == ALOAD_1 || op.Op == ALOAD_2 || op.Op == ALOAD_3 {
		return op.Name
	}

	data := buf.Read(int(op.Length))[0]
	return fmt.Sprintf("%s %d", op.Name, data)
}

func decompileGetStaticInstruction(op Opcode, buf *parser.ByteReader, constantPool *class.ConstantPool) string {
	data := buf.Read(int(op.Length))
	index := (data[0] << 8) | data[1]

	poolItem := constantPool.Get(uint16(index))
	var ref *class.CONSTANT_Fieldref_info

	switch poolItem.Tag {
	case class.CONSTANT_Fieldref:
		ref = poolItem.Info.(*class.CONSTANT_Fieldref_info)
	default:
		log.Fatalf("unhandled tag type '%d' for GETSTATIC instruction", poolItem.Tag)
	}

	classIndex := ref.ClassIndex
	classRef := constantPool.Get(classIndex).Info.(*class.CONSTANT_Class_info)
	className := constantPool.GetString(classRef.NameIndex)

	nameAndTypeRef := constantPool.Get(ref.NameAndTypeIndex).Info.(*class.CONSTANT_NameAndType_info)
	nameIndex := nameAndTypeRef.NameIndex
	descriptorIndex := nameAndTypeRef.DescriptorIndex

	fieldName := constantPool.GetString(nameIndex)
	descriptor := constantPool.GetString(descriptorIndex)

	return fmt.Sprintf("%s %s/%s %s", op.Name, className, fieldName, descriptor)
}

func decompileLDCInstruction(op Opcode, buf *parser.ByteReader, constantPool *class.ConstantPool) string {
	data := buf.Read(int(op.Length))
	index := uint16(data[0])

	poolItem := constantPool.Get(index)

	var value interface{}
	switch poolItem.Tag {
	case class.CONSTANT_String:
		value = constantPool.GetString(poolItem.Info.(*class.CONSTANT_String_info).StringIndex)
	default:
		log.Fatalf("unhandled tag type '%d' for LDC instruction", poolItem.Tag)
	}
	return fmt.Sprintf("%s %v", op.Name, value)
}

func decompileInvokeInstruction(op Opcode, buf *parser.ByteReader, constantPool *class.ConstantPool) string {
	data := buf.Read(int(op.Length))
	index := (data[0] << 8) | data[1]

	poolItem := constantPool.Get(uint16(index))

	ref := poolItem.Info.(*class.CONSTANT_Methodref_info)
	classIndex := ref.ClassIndex
	classRef := constantPool.Get(classIndex).Info.(*class.CONSTANT_Class_info)
	className := constantPool.GetString(classRef.NameIndex)

	nameAndTypeRef := constantPool.Get(ref.NameAndTypeIndex).Info.(*class.CONSTANT_NameAndType_info)
	nameIndex := nameAndTypeRef.NameIndex
	descriptorIndex := nameAndTypeRef.DescriptorIndex

	methodName := constantPool.GetString(nameIndex)
	descriptor := constantPool.GetString(descriptorIndex)

	return fmt.Sprintf("%s %s.%s %s", op.Name, className, methodName, descriptor)
}
