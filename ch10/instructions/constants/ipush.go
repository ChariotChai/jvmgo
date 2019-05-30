package constants

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

//byte
type BIPUSH struct { val int8 }
func (self *BIPUSH) FetchOp(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OpStack().PushInt(i)
}

//short
type SIPUSH struct { val int16 }
func (self *SIPUSH) FetchOp(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OpStack().PushInt(i)
}
