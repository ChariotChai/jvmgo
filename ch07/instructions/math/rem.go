package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "math"

//求余
type IREM struct { base.NoOpInstruction }
func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmaticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

type LREM struct { base.NoOpInstruction }
func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmaticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

type DREM struct { base.NoOpInstruction }
func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type FREM struct { base.NoOpInstruction }
func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := math.Mod(float64(v1), float64(v2))
	stack.PushFloat(float32(result))
}
