package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type IMUL struct { base.NoOpInstruction }
func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

type LMUL struct { base.NoOpInstruction }
func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}

type FMUL struct { base.NoOpInstruction }
func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

type DMUL struct { base.NoOpInstruction }
func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}
