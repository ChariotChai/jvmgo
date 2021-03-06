package control

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/

type TABLE_SWITCH struct {
	defaultOffset 	int32
	low				int32
	high			int32
	jumpOffsets		[]int32
}

func (self *TABLE_SWITCH) FetchOp(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCnt := self.high - self.low
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCnt)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OpStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index - self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
