package loads

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

type AALOAD struct {
	base.NoOpInstruction
}

func (a *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

type BALOAD struct {
	base.NoOpInstruction
}

func (a *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

type CALOAD struct {
	base.NoOpInstruction
}

func (a *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

type DALOAD struct {
	base.NoOpInstruction
}

func (a *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(float64(doubles[index]))
}

type FALOAD struct {
	base.NoOpInstruction
}

func (a *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(float32(floats[index]))
}

type IALOAD struct {
	base.NoOpInstruction
}

func (a *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(int32(ints[index]))
}

type LALOAD struct {
	base.NoOpInstruction
}

func (a *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(int64(longs[index]))
}

type SALOAD struct {
	base.NoOpInstruction
}

func (a *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
