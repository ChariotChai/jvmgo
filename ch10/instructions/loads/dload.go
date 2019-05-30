package loads

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

//指令索引来自操作数
type DLOAD struct { base.Index8Instruction }

//指令索引隐含在操作码中
type DLOAD_0 struct { base.NoOpInstruction }
type DLOAD_1 struct { base.NoOpInstruction }
type DLOAD_2 struct { base.NoOpInstruction }
type DLOAD_3 struct { base.NoOpInstruction }

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OpStack().PushDouble(val)
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(self.Index))
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
