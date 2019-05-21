package conversions

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type D2I struct { base.NoOpInstruction }
func (self *D2I) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

type D2F struct { base.NoOpInstruction }
func (self *D2F) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

type D2L struct { base.NoOpInstruction }
func (self *D2L) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
