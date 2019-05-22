package math

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOp(reader *base.BytecodeReader)  {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
