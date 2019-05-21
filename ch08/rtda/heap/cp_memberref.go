package heap

import "jvmgo/ch08/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (m *MemberRef) Name() string {
	return m.name
}

func (m *MemberRef) Descriptor() string {
	return m.descriptor
}
