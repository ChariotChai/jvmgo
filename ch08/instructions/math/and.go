package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type IAND struct { base.NoOpInstruction }
func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

type LAND struct { base.NoOpInstruction }
func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
