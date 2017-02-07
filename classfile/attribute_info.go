package classfile

type AttributeInfo interface {
	readInfo(cr *ClassReader)
}

func readAttributes(cr *ClassReader, cp ConstantPool) []AttributeInfo {
	n := cr.readUint16()
	ais := make([]AttributeInfo, n)

	for i := range ais {
		ais[i] = readAttribute(cr, cp)
	}
	return ais
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	return nil
}

func readAttribute(cr *ClassReader, cp ConstantPool) AttributeInfo {
	nameIdx := cr.readUint16()
	attrName := cp.getUtf8(nameIdx)
	attrLen := cr.readUint32()
	info := newAttributeInfo(attrName, attrLen, cp)
	info.readInfo(cr)
	return info
}
