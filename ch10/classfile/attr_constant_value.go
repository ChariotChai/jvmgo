package classfile

type ConstantValueAttributes struct {
	constantValueIndex 		uint16
}

func (self *ConstantValueAttributes) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttributes) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
