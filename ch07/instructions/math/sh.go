package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type ISHL struct { base.NoOpInstruction }
func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f //位移值需要为无符号整数，且32位只需要5个bit位
	result := v1 << s
	stack.PushInt(result)
}

type ISHR struct { base.NoOpInstruction } //算术右位移
func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

type IUSHR struct { base.NoOpInstruction } //逻辑右位移
func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

type LSHL struct { base.NoOpInstruction }
func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f //位移值需要为无符号整数，且64位只需要6个bit位
	result := v1 << s
	stack.PushLong(result)
}

type LSHR struct { base.NoOpInstruction }
func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f //位移值需要为无符号整数，且64位只需要6个bit位
	result := v1 >> s
	stack.PushLong(result)
}

type LUSHR struct { base.NoOpInstruction }
func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
