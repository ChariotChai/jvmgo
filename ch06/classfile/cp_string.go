package classfile

type ConstantStringInfo struct {
	//本身不存放数据，仅保存了常量池索引
	cp 				ConstantPool
	stringIndex		uint16 	//该索引指向utf8常量
}
func (self *ConstantStringInfo) readInfo(reader *ClassReader)  {
	self.stringIndex = reader.readUint16()
}
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}