package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type INEG struct { base.NoOpInstruction }
func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v1 := stack.PopInt()
	result := -v1
	stack.PushInt(result)
}

type LNEG struct { base.NoOpInstruction }
func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v1 := stack.PopLong()
	result := -v1
	stack.PushLong(result)
}

type FNEG struct { base.NoOpInstruction }
func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v1 := stack.PopFloat()
	result := -v1
	stack.PushFloat(result)
}

type DNEG struct { base.NoOpInstruction }
func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v1 := stack.PopDouble()
	result := -v1
	stack.PushDouble(result)
}
