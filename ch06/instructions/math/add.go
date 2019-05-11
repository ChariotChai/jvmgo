package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type IADD struct { base.NoOpInstruction }
func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

type LADD struct { base.NoOpInstruction }
func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}

type FADD struct { base.NoOpInstruction }
func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

type DADD struct { base.NoOpInstruction }
func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}
