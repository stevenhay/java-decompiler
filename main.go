package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/stevenhay/GVM/bytecode"
	"github.com/stevenhay/GVM/class"
	"github.com/stevenhay/GVM/parser"
)

func main() {
	file := "Main.class"
	reader, err := parser.NewClassReader(file)
	if err != nil {
		log.Fatalf("could not load class file - %s", err)
	}

	cls := reader.Parse()

	magic := hex.EncodeToString(cls.Magic)

	fmt.Printf("Magic: %s\n", magic)
	fmt.Printf("Major: %d\n", cls.MajorVersion)
	fmt.Printf("Minor: %d\n", cls.MinorVersion)
	fmt.Printf("Constant Pool Count: %d\n", cls.ConstantPoolCount)
	fmt.Printf("Field Count: %d\n", cls.FieldCount)
	fmt.Printf("Interface Count: %d\n", cls.InterfaceCount)
	fmt.Printf("Method Count: %d\n", cls.MethodCount)
	for i, f := range cls.Methods {
		fmt.Printf("Method %d: \n", i)
		fmt.Printf("\tName: %s\n", f.GetName(cls.ConstantPool))
		fmt.Printf("\tDescriptor: %s\n", f.GetDescriptor(cls.ConstantPool))
		fmt.Printf("\tAccess Flags: %s\n", strings.Join(f.GetMethodAccessFlags(), ", "))

		codeAttributes := f.GetCodeAttributes(cls.ConstantPool)
		for i, ca := range codeAttributes {
			fmt.Printf("\tCode Attribute %d:\n", i)
			fmt.Printf("\tMax Locals: %d:\n", ca.MaxLocals)
			fmt.Printf("\tMax Stack: %d:\n", ca.MaxStack)
			fmt.Printf("\tAttributes Count: %d:\n", ca.AttributesCount)

			for i, f := range ca.AttributeInfo {
				fmt.Printf("\tAttribute %d:\n", i)

				name := f.GetName(cls.ConstantPool)
				fmt.Printf("\t\tName: %s\n", name)
				if name == class.LineNumberTable {
					lnt := f.GetAsLineNumberTableAttribute()
					for i := 0; i < int(lnt.LineNumberTableLength); i++ {
						fmt.Printf("\t\tLineNumber %d: %d\n", i, lnt.LineNumberTable[i].LineNumber)
					}
				}
			}
		}

		fmt.Printf("\tBytecode (bytes): \n")
		bc := bytecode.DecompileBytes(f.GetCode(cls.ConstantPool), cls.ConstantPool)
		for _, s := range bc {
			fmt.Printf("\t\t%s\n", s)
		}

		fmt.Printf("\tBytecode: \n")
		bc = bytecode.Decompile(f.GetCode(cls.ConstantPool), cls.ConstantPool)
		for _, s := range bc {
			fmt.Printf("\t\t%s\n", s)
		}
	}
	fmt.Printf("Attribute Count: %d\n", cls.AttributeCount)
	for i, f := range cls.Attributes {
		fmt.Printf("Attribute %d:\n", i)

		name := f.GetName(cls.ConstantPool)
		fmt.Printf("\tName: %s\n", name)
		if name == class.SourceFile {
			sourceFile := cls.ConstantPool.GetString(f.GetAsSourceFileAttribute().SourceFileIndex)
			fmt.Printf("\tValue: %s\n", sourceFile)
		}
	}
}
