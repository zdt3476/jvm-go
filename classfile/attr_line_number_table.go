package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (l *LineNumberTableAttribute) readInfo(cr *ClassReader) {
	lens := cr.readUint16()
	l.lineNumberTable = make([]*LineNumberTableEntry, lens)
	for i := range l.lineNumberTable {
		l.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    cr.readUint16(),
			lineNumber: cr.readUint16(),
		}
	}
}
