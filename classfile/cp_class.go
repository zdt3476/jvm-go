package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ConstantClassInfo) readInfo(cr *ClassReader) {
	c.nameIndex = cr.readUint16()
}

func (c *ConstantClassInfo) String() string {
	return c.cp.getUtf8(c.nameIndex)
}
