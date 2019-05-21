package loads

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

//指令索引来自操作数
type ILOAD struct{ base.Index8Instruction }

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

//指令索引隐含在操作码中
type ILOAD_0 struct{ base.NoOpInstruction }
type ILOAD_1 struct{ base.NoOpInstruction }
type ILOAD_2 struct{ base.NoOpInstruction }
type ILOAD_3 struct{ base.NoOpInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OpStack().PushInt(val)
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
