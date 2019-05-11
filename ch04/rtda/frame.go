package rtda

/**
 * 执行方法所需要的局部变量表大小和操作数栈大小是由编译器提前计算好的
 * 存储在class文件的method_info结构的Code属性中
 */
type Frame struct {
	lower		*Frame
	localVars	LocalVars //局部变量表
	opStack		*OpStack  //操作数栈
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame {
		localVars: newLocalVars(maxLocals),
		opStack: newOpStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OpStack() *OpStack {
	return self.opStack
}
