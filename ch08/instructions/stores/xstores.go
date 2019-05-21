package stores

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

type AASTORE struct {
	base.NoOpInstruction
}

func (self *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = val
}

type BASTORE struct {
	base.NoOpInstruction
}

func (self *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	bytes[index] = int8(val)
}

type CASTORE struct {
	base.NoOpInstruction
}

func (self *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	chars[index] = uint16(val)
}

type DASTORE struct {
	base.NoOpInstruction
}

func (self *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	doubles[index] = float64(val)
}

type FASTORE struct {
	base.NoOpInstruction
}

func (self *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	floats[index] = float32(val)
}

type IASTORE struct {
	base.NoOpInstruction
}

func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}

type LASTORE struct {
	base.NoOpInstruction
}

func (self *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	longs[index] = int64(val)
}

type SASTORE struct {
	base.NoOpInstruction
}

func (self *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	shorts[index] = int16(val)
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
