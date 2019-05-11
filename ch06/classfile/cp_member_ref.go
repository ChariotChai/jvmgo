package classfile

type ConstantMemberrefInfo struct {
	cp					ConstantPool
	classIndex			uint16
	nameAndTypeIndex	uint16
}
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDiscriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// go语言中没有继承，这里通过结构体嵌套的方式实现
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
