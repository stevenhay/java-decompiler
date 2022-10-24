package bytecode

import (
	"fmt"
	"log"
	"strings"

	"github.com/stevenhay/GVM/class"
	"github.com/stevenhay/GVM/parser"
)

func Decompile(b []byte, constantPool *class.ConstantPool) []string {
	buf := parser.NewByteReader(b)

	instructions := make([]string, 0)
	for buf.Len() > 0 {
		opb := buf.Read(1)[0]
		op, ok := Opcodes[opb]
		if !ok {
			log.Fatalf("OPCODE '%d' is unknown", opb)
		}

		// Assume it's a 0 byte operation and prepare to print
		// just the name
		var instruction = op.Name

		// If there's more data to read, then decompile the instruction bytes
		if op.Length > 0 {
			switch op.Op {
			case INVOKESTATIC:
				fallthrough
			case INVOKEVIRTUAL:
				fallthrough
			case INVOKESPECIAL:
				instruction = decompileInvokeInstruction(op, buf, constantPool)

			case BIPUSH:
				instruction = decompileBIPUSHInstruction(op, buf, constantPool)

			case GETSTATIC:
				instruction = decompileGetStaticInstruction(op, buf, constantPool)

			case LDC:
				instruction = decompileLDCInstruction(op, buf, constantPool)
			case NEW:
				instruction = decompileNEWInstruction(op, buf, constantPool)
			default:
				log.Fatalf("%s is not handled yet", op.Name)
			}
		}
		instructions = append(instructions, instruction)
	}
	return instructions
}

func DecompileBytes(b []byte, constantPool *class.ConstantPool) []string {
	buf := parser.NewByteReader(b)

	instructions := make([]string, 0)
	for buf.Len() > 0 {
		opb := buf.Read(1)[0]
		op, ok := Opcodes[opb]
		if !ok {
			log.Fatalf("OPCODE '0x%X' is unknown", opb)
		}

		if op.Length > 0 {
			data := buf.Read(int(op.Length))
			instructions = append(instructions, fmt.Sprintf("%s %s", op.Name, strings.Trim(fmt.Sprintf("%v", data), "[]")))
		} else {
			instructions = append(instructions, fmt.Sprintf("%s", op.Name))
		}
	}
	return instructions
}
