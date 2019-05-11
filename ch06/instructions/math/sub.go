package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type ISUB struct { base.NoOpInstruction }
func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

type LSUB struct { base.NoOpInstruction }
func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}

type FSUB struct { base.NoOpInstruction }
func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

type DSUB struct { base.NoOpInstruction }
func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}
