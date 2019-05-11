package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.isProtected() {
		return d == c || d.getPackageName() == c.getPackageName()
	}
	if !self.isPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (c *ClassMember) Class() *Class {
	return c.class
}

func (c *ClassMember) IsStatic() {

}
