package constants

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type NOP struct { base.NoOpInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
	//nop
}
