package constants

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type NOP struct { base.NoOpInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
	//nop
}
