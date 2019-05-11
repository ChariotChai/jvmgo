package heap

import "jvmgo/ch06/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}
func newInterfaceMethodRef(cp *ConstantPool, 
	refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
