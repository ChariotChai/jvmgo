package loads

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

//指令索引来自操作数
type LLOAD struct { base.Index8Instruction }

//指令索引隐含在操作码中
type LLOAD_0 struct { base.NoOpInstruction }
type LLOAD_1 struct { base.NoOpInstruction }
type LLOAD_2 struct { base.NoOpInstruction }
type LLOAD_3 struct { base.NoOpInstruction }

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OpStack().PushLong(val)
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(self.Index))
}

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
