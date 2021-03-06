package comparisons

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type LCMP struct { base.NoOpInstruction }

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
