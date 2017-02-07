package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (ua *UnparsedAttribute) readInfo(cr *ClassReader) {
	ua.info = cr.readBytes(ua.length)
}
