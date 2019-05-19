package conversions

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type L2I struct { base.NoOpInstruction }
func (self *L2I) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}

type L2F struct { base.NoOpInstruction }
func (self *L2F) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

type L2D struct { base.NoOpInstruction }
func (self *L2D) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}
