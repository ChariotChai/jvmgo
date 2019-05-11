package constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type NOP struct { base.NoOpInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
	//nop
}
