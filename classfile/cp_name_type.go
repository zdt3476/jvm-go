package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *ConstantNameAndTypeInfo) readInfo(cr *ClassReader) {
	c.nameIndex = cr.readUint16()
	c.descriptorIndex = cr.readUint16()
}
