package classfile

type ExceptionAttribute struct {
	exceptionIndexTable 	[]uint16
}

func (self *ExceptionAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionAttribute) ExceptionIndextable() []uint16 {
	return self.exceptionIndexTable
}
