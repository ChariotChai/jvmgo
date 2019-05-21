package conversions

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type I2B struct { base.NoOpInstruction }
func (self *I2B) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

type I2C struct { base.NoOpInstruction }
func (self *I2C) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}

type I2L struct { base.NoOpInstruction }
func (self *I2L) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	i := stack.PopInt()
	c := int64(i)
	stack.PushLong(c)
}

type I2S struct { base.NoOpInstruction }
func (self *I2S) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

type I2F struct { base.NoOpInstruction }
func (self *I2F) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

type I2D struct { base.NoOpInstruction }
func (self *I2D) Execute(frame *rtda.Frame)  {
	stack := frame.OpStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
