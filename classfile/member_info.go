package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(cr *ClassReader, cp ConstantPool) []*MemberInfo {
	mc := cr.readUint16()
	mis := make([]*MemberInfo, mc)
	for i := range mis {
		mis[i] = readMember(cr, cp)
	}
	return mis
}

func readMember(cr *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlag:      cr.readUint16(),
		nameIndex:       cr.readUint16(),
		descriptorIndex: cr.readUint16(),
		attributes:      readAttributes(cr, cp),
	}
}

func (mi *MemberInfo) AccessFlag() uint16 {
	return mi.accessFlag
}

func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}
