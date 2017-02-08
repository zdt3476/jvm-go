package classfile

import "fmt"

const (
	classMagic uint32 = 0xCAFEBABE
)

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlag   uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cf = &ClassFile{}
	cr := &ClassReader{classData}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(cr *ClassReader) {
	cf.readAndCheckMagic(cr)
	cf.readAndCheckVersion(cr)
	cf.constantPool = readConstantPool(cr)
	cf.accessFlag = cr.readUint16()
	cf.thisClass = cr.readUint16()
	cf.superClass = cr.readUint16()
	cf.interfaces = cr.readUint16s()
	cf.fields = readMembers(cr, cf.constantPool)
	cf.methods = readMembers(cr, cf.constantPool)
	cf.attributes = readAttributes(cr, cf.constantPool)
}

func (cf *ClassFile) readAndCheckMagic(cr *ClassReader) {
	magic := cr.readUint32()
	if magic != classMagic {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(cr *ClassReader) {
	minor := cr.readUint16()
	major := cr.readUint16()

	switch major {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if minor == 0 {
			return
		}
	}

	panic("java.lang.UnsupportClassVersionError!")
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlag() uint16 {
	return cf.accessFlag
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return nil
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return nil
}

func (cf *ClassFile) InterfaceNames() []string {
	ins := make([]string, len(cf.interfaces))
	for i, idx := range cf.interfaces {
		ins[i] = cf.constantPool.getClassName(idx)
	}

	return ins
}
