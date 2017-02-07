package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (c *ConstantIntegerInfo) readInfo(cr *ClassReader) {
	c.val = int32(cr.readUint32())
}

type ConstantFloatInfo struct {
	val float32
}

func (c *ConstantFloatInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint32()
	c.val = math.Float32frombits(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantDoubleInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint64()
	c.val = math.Float64frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (c *ConstantLongInfo) readInfo(cr *ClassReader) {
	c.val = int64(cr.readUint64())
}
