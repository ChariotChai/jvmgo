package stack

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

//for 32-bts
type POP struct { base.NoOpInstruction }
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	stack.PopSlot()
}

//for 64-bits
type POP2 struct { base.NoOpInstruction }
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	stack.PopSlot()
	stack.PopSlot()
}
