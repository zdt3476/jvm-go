package classfile

type ConstantPool []ConstantInfo

func readConstantPool(cr *ClassReader) ConstantPool {
	cnt := int(cr.readUint16())
	cp := make(ConstantPool, cnt)

	for i := 1; i < cnt; i++ {
		cp[i] = readConstantInfo(cr, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo: // 占两个位置
			i++
		}
	}

	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if info := cp[index]; info != nil {
		return info
	}
	panic("Invalid constant pool index!")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	info := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(info.nameIndex)
	typ := cp.getUtf8(info.descriptorIndex)
	return name, typ
}

func (cp ConstantPool) getClassName(index uint16) string {
	info := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(info.nameIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return info.str
}
