package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOpInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
