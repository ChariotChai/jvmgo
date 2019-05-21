package constants

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type NOP struct { base.NoOpInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
	//nop
}
