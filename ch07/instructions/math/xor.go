package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type IXOR struct { base.NoOpInstruction }
func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

type LXOR struct { base.NoOpInstruction }
func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
