package heap

import "jvmgo/ch09/classfile"

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
	if self.IsProtected() {
		return d == c || d.GetPackageName() == c.GetPackageName()
	}
	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}

func (c *ClassMember) Descriptor() string {
	return c.descriptor
}

func (c *ClassMember) Name() string {
	return c.name
}

func (c *ClassMember) Class() *Class {
	return c.class
}

func (c *ClassMember) IsPublic() bool {
	return 0 != c.accessFlags&ACC_PUBLIC
}

func (c *ClassMember) IsProtected() bool {
	return 0 != c.accessFlags&ACC_PROTECTED
}

func (c *ClassMember) IsPrivate() bool {
	return 0 != c.accessFlags&ACC_PRIVATE
}

func (c *ClassMember) IsStatic() bool {
	return 0 != c.accessFlags&ACC_STATIC
}

func (c *ClassMember) IsFinal() bool {
	return 0 != c.accessFlags&ACC_SYNTHETIC
}
