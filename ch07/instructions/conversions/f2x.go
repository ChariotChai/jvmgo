package conversions

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type F2I struct { base.NoOpInstruction }
func (self *F2I) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

type F2D struct { base.NoOpInstruction }
func (self *F2D) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

type F2L struct { base.NoOpInstruction }
func (self *F2L) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
