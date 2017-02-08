package classfile

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (c *CodeAttribute) readInfo(cr *ClassReader) {
	c.maxStack = cr.readUint16()
	c.maxLocals = cr.readUint16()
	codeLen := cr.readUint32()
	c.code = cr.readBytes(codeLen)
	c.exceptionTable = readExceptionTable(cr)
	c.attributes = readAttributes(cr, c.cp)
}

func readExceptionTable(cr *ClassReader) []*ExceptionTableEntry {
	lens := cr.readUint16()
	etes := make([]*ExceptionTableEntry, lens)

	for i := range etes {
		etes[i] = &ExceptionTableEntry{
			startPc:   cr.readUint16(),
			endPc:     cr.readUint16(),
			handlerPc: cr.readUint16(),
			catchType: cr.readUint16(),
		}
	}

	return etes
}
