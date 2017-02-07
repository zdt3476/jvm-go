package classfile

type MarkerAttribute struct{}

func (am *MarkerAttribute) readInfo(cr *ClassReader) {
	// nothing
}

type DeprecatedAttribute struct{ MarkerAttribute }

type SyntheticAttribute struct{ MarkerAttribute }
