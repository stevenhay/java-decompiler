package class

const (
	ConstantValue                        = "ConstantValue"
	Code                                 = "Code"
	StackMapTable                        = "StackMapTable"
	Exceptions                           = "Exceptions"
	InnerClasses                         = "InnerClasses"
	EnclosingMethod                      = "EnclosingMethod"
	Synthetic                            = "Synthetic"
	Signature                            = "Signature"
	SourceFile                           = "SourceFile"
	SourceDebugExtension                 = "SourceDebugExtension"
	LineNumberTable                      = "LineNumberTable"
	LocalVariableTable                   = "LocalVariableTable"
	LocalVariableTypeTable               = "LocalVariableTypeTable"
	Deprecated                           = "Deprecated"
	RuntimeVisibleAnnotations            = "RuntimeVisibleAnnotations"
	RuntimeInvisibleAnnotations          = "RuntimeInvisibleAnnotations"
	RuntimeVisibleParameterAnnotations   = "RuntimeVisibleParameterAnnotations"
	RuntimeInvisibleParameterAnnotations = "RuntimeInvisibleParameterAnnotations"
	AnnotationDefault                    = "AnnotationDefault"
	BootstrapMethods                     = "BootstrapMethods"
)

type AttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               interface{}
}

func (a *AttributeInfo) GetName(constantPool *ConstantPool) string {
	return string(constantPool.Get(a.AttributeNameIndex).Info.(*CONSTANT_Utf8_info).Bytes)
}

func (a *AttributeInfo) GetAsCodeAttribute() *Code_attribute {
	return a.Info.(*Code_attribute)
}

func (a *AttributeInfo) GetAsLineNumberTableAttribute() *LineNumberTable_attribute {
	return a.Info.(*LineNumberTable_attribute)
}

func (a *AttributeInfo) GetAsSourceFileAttribute() *SourceFile_attribute {
	return a.Info.(*SourceFile_attribute)
}

type Code_attribute struct {
	MaxStack             uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte
	ExceptionTableLength uint16
	ExceptionTables      []*ExceptionTable
	AttributesCount      uint16
	AttributeInfo        []*AttributeInfo
}

type LineNumberTable_attribute struct {
	LineNumberTableLength uint16
	LineNumberTable       []*LineNumberTableData
}

type LineNumberTableData struct {
	StartPc    uint16
	LineNumber uint16
}

type SourceFile_attribute struct {
	SourceFileIndex uint16
}
