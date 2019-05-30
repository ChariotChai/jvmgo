package classfile

type UnparsedAttributes struct {
	attrName		string 
	attrLen			uint32
	info			[]byte				
}

func (self *UnparsedAttributes) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.attrLen)
}
