package constants

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

type NOP struct { base.NoOpInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
	//nop
}
