package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (c *ConstantStringInfo) readInfo(cr *ClassReader) {
	c.stringIndex = cr.readUint16()
}

func (c *ConstantStringInfo) String() string {
	return c.cp.getUtf8(c.stringIndex)
}
