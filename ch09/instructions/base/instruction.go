package base
import "jvmgo/ch09/rtda"

type Instruction interface {
	FetchOp(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOpInstruction struct {}
func (self *NoOpInstruction) FetchOp(reader *BytecodeReader) {
}

type BranchInstruction struct {
	Offset int
}
func (self *BranchInstruction) FetchOp(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}
func (self *Index8Instruction) FetchOp(reader *BytecodeReader) {
	self.Index = uint(reader.ReadInt8())
}

type Index16Instruction struct {
	Index uint
}
func (self *Index16Instruction) FetchOp(reader *BytecodeReader) {
	self.Index = uint(reader.ReadInt16())
}
