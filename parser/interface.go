package parser

func (r *ClassReader) readInterfaceIndexes() []uint16 {
	indexes := make([]uint16, r.class.InterfaceCount)
	for i := 0; i < int(r.class.InterfaceCount); i++ {
		indexes[i] = r.b.ReadUint16()
	}
	return indexes
}
