package classfile

type ExceptionAttribute struct {
	exceptionTableIndex uint16
}

func (e *ExceptionAttribute) readInfo(cr *ClassReader) {
	e.exceptionTableIndex = cr.readUint16()
}

func (e *ExceptionAttribute) ExceptionTableIndex() uint16 {
	return e.exceptionTableIndex
}
