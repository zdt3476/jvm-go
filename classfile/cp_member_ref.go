package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *ConstantMemberRefInfo) readInfo(cr *ClassReader) {
	c.classIndex = cr.readUint16()
	c.nameAndTypeIndex = cr.readUint16()
}

func (c *ConstantMemberRefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}
