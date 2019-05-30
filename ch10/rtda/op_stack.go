package rtda

import "math"
import "jvmgo/ch10/rtda/heap"

type OpStack struct {
	size  uint
	slots []Slot
}

func newOpStack(maxStack uint) *OpStack {
	if maxStack > 0 {
		return &OpStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}
func (self *OpStack) PushBoolean(val bool) {
	if val {
		self.PushInt(1)
	} else {
		self.PushInt(0)
	}
}

func (self *OpStack) PopBoolean() bool {
	return self.PopInt() == 1
}

func (self *OpStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}
func (self *OpStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}
func (self *OpStack) PushFloat(val float32) {
	self.slots[self.size].num = int32(math.Float32bits(val))
	self.size++
}
func (self *OpStack) PopFloat() float32 {
	self.size--
	return math.Float32frombits(uint32(self.slots[self.size].num))
}
func (self *OpStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}
func (self *OpStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return (int64(high) << 32) | int64(low)
}
func (self *OpStack) PushDouble(val float64) {
	self.PushLong(int64(math.Float64bits(val)))
}
func (self *OpStack) PopDouble() float64 {
	return math.Float64frombits(uint64(self.PopLong()))
}
func (self *OpStack) PushRef(ref *heap.Object) {
	self.slots[self.size].ref = ref
	self.size++
}
func (self *OpStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil //内存回收意识很重要
	return ref
}
func (self *OpStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}
func (self *OpStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}
func (self *OpStack) TopSlot() Slot {
	return self.slots[self.size-1]
}
func (self *OpStack) GetRefFromTop(n uint) *heap.Object {
	//????
	return self.slots[self.size-1-n].ref
}
func (self *OpStack) Clear() {
	self.size = 0
	for i := range self.slots {
		self.slots[i].ref = nil
	}
}
