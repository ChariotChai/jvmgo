package stores

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type FSTORE struct { base.Index8Instruction }
func (self *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(self.Index))
}

type FSTORE_0 struct { base.NoOpInstruction }
func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct { base.NoOpInstruction }
func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct { base.NoOpInstruction }
func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct { base.NoOpInstruction }
func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtda.Frame, index uint)  {
	val := frame.OpStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
