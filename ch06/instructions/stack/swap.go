package stack

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type SWAP struct { base.NoOpInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
