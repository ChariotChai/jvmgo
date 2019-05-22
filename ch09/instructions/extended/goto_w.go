package extended

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"


type GOTO_W struct { 
	offset int
}
func (self *GOTO_W) FetchOp(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
