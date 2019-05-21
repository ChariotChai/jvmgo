package loads

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

//指令索引来自操作数
type FLOAD struct { base.Index8Instruction }

//指令索引隐含在操作码中
type FLOAD_0 struct { base.NoOpInstruction }
type FLOAD_1 struct { base.NoOpInstruction }
type FLOAD_2 struct { base.NoOpInstruction }
type FLOAD_3 struct { base.NoOpInstruction }

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OpStack().PushFloat(val)
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, uint(self.Index))
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
