package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type IDIV struct { base.NoOpInstruction }
func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushInt(result)
}

type LDIV struct { base.NoOpInstruction }
func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushLong(result)
}

type FDIV struct { base.NoOpInstruction }
func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

type DDIV struct { base.NoOpInstruction }
func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}
