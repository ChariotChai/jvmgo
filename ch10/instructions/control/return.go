package control

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type RETURN struct{ base.NoOpInstruction }

func (r *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type IRETURN struct{ base.NoOpInstruction }

func (r *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OpStack().PopInt()
	invokerFrame.OpStack().PushInt(retVal)
}

type FRETURN struct{ base.NoOpInstruction }

func (r *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OpStack().PopFloat()
	invokerFrame.OpStack().PushFloat(retVal)
}

type LRETURN struct{ base.NoOpInstruction }

func (r *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OpStack().PopLong()
	invokerFrame.OpStack().PushLong(retVal)
}

type DRETURN struct{ base.NoOpInstruction }

func (r *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OpStack().PopDouble()
	invokerFrame.OpStack().PushDouble(retVal)
}

type ARETURN struct{ base.NoOpInstruction }

func (r *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OpStack().PopRef()
	invokerFrame.OpStack().PushRef(retVal)
}
