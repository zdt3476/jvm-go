package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (sfa *SourceFileAttribute) readInfo(cr *ClassReader) {
	sfa.sourceFileIndex = cr.readUint16()
}

func (sfa *SourceFileAttribute) FileName() string {
	return sfa.cp.getUtf8(sfa.sourceFileIndex)
}
