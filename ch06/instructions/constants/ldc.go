package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type LDC struct {
	base.Index8Instruction
}

func (l *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, l.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (l *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, l.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OpStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	default:
		panic("todo: ldc")
	}
}

type LDC2_W struct {
	base.Index16Instruction
}

func (l *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(l.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("todo: ldc")
	}
}
