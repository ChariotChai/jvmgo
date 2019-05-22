package constants

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

type NOP struct { base.NoOpInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
	//nop
}
