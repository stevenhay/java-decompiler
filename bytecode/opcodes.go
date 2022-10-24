package bytecode

const (
	GETSTATIC = 0xB2

	INVOKESTATIC  = 0xB8
	INVOKESPECIAL = 0xB7
	INVOKEVIRTUAL = 0xB6

	IRETURN = 0xAC
	RETURN  = 0xB1
	BIPUSH  = 0x10

	ICONST_M1 = 0x2
	ICONST_0  = 0x3
	ICONST_1  = 0x4
	ICONST_2  = 0x5
	ICONST_3  = 0x6
	ICONST_4  = 0x7
	ICONST_5  = 0x8

	ISTORE_0 = 0x3B
	ISTORE_1 = 0x3C
	ISTORE_2 = 0x3D
	ISTORE_3 = 0x3E

	ILOAD_0 = 0x1A
	ILOAD_1 = 0x1B
	ILOAD_2 = 0x1C
	ILOAD_3 = 0x1D

	IADD = 0x60

	ALOAD_0 = 0x2A
	ALOAD_1 = 0x2B
	ALOAD_2 = 0x2C
	ALOAD_3 = 0x2D

	LDC = 0x12
	NEW = 0xBB
	DUP = 0x59
)

var Opcodes = map[uint8]Opcode{
	// 0 byte operations
	ICONST_M1: {Op: ICONST_M1, Name: "ICONST_M1", Length: 0},
	ICONST_0:  {Op: ICONST_0, Name: "ICONST_0", Length: 0},
	ICONST_1:  {Op: ICONST_1, Name: "ICONST_1", Length: 0},
	ICONST_2:  {Op: ICONST_2, Name: "ICONST_2", Length: 0},
	ICONST_3:  {Op: ICONST_3, Name: "ICONST_3", Length: 0},
	ICONST_4:  {Op: ICONST_4, Name: "ICONST_4", Length: 0},
	ICONST_5:  {Op: ICONST_5, Name: "ICONST_5", Length: 0},
	ALOAD_0:   {Op: ALOAD_0, Name: "ALOAD_0", Length: 0},
	ALOAD_1:   {Op: ALOAD_1, Name: "ALOAD_1", Length: 0},
	ALOAD_2:   {Op: ALOAD_2, Name: "ALOAD_2", Length: 0},
	ALOAD_3:   {Op: ALOAD_3, Name: "ALOAD_3", Length: 0},
	ILOAD_0:   {Op: ILOAD_0, Name: "ILOAD_0", Length: 0},
	ILOAD_1:   {Op: ILOAD_1, Name: "ILOAD_1", Length: 0},
	ILOAD_2:   {Op: ILOAD_2, Name: "ILOAD_2", Length: 0},
	ILOAD_3:   {Op: ILOAD_3, Name: "ILOAD_3", Length: 0},
	ISTORE_0:  {Op: ISTORE_0, Name: "ISTORE_0", Length: 0},
	ISTORE_1:  {Op: ISTORE_1, Name: "ISTORE_1", Length: 0},
	ISTORE_2:  {Op: ISTORE_2, Name: "ISTORE_2", Length: 0},
	ISTORE_3:  {Op: ISTORE_3, Name: "ISTORE_3", Length: 0},
	IRETURN:   {Op: IRETURN, Name: "IRETURN", Length: 0},
	RETURN:    {Op: RETURN, Name: "RETURN", Length: 0},
	IADD:      {Op: IADD, Name: "IADD", Length: 0},
	DUP:       {Op: DUP, Name: "DUP", Length: 0},

	// 1 byte operations
	BIPUSH: {Op: BIPUSH, Name: "BIPUSH", Length: 1},
	LDC:    {Op: LDC, Name: "LDC", Length: 1},

	// 2 byte operations
	GETSTATIC:     {Op: GETSTATIC, Name: "GETSTATIC", Length: 2},
	INVOKESTATIC:  {Op: INVOKESTATIC, Name: "INVOKESTATIC", Length: 2},
	INVOKEVIRTUAL: {Op: INVOKEVIRTUAL, Name: "INVOKEVIRTUAL", Length: 2},
	INVOKESPECIAL: {Op: INVOKESPECIAL, Name: "INVOKESPECIAL", Length: 2},
	NEW:           {Op: NEW, Name: "NEW", Length: 2},
}

type Opcode struct {
	Op     uint8
	Name   string
	Length uint8
}
