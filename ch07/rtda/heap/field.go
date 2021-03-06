package heap

import "jvmgo/ch07/classfile"

type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if varAttr := cfField.ConstantValueAttribute(); varAttr != nil {
		self.constValueIndex = uint(varAttr.ConstantValueIndex())
	}
}

func (f *Field) SlotId() uint {
	return f.slotId
}

func (f *Field) ConstValueIndex() uint {
	return f.constValueIndex
}
