package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/rtda/heap"
)

type PUT_STATIC struct {
	base.Index16Instruction
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currMethod := frame.Method()
	currClass := currMethod.Class()
	cp := currClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	//必须是静态字段
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//final字段，只能在类初始化方法中赋值
	if field.IsFinal() {
		if currClass != class || currMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
}
