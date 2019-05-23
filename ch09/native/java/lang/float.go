package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRowIntBits)
}

func floatToRowIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OpStack().PushInt(int32(bits))
}
