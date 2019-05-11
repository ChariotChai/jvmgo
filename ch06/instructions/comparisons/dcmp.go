package comparisons

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type DCMPG struct { base.NoOpInstruction }
type DCMPL struct { base.NoOpInstruction }

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, flag bool) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
