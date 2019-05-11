package rtda

/**
 * 执行方法所需要的局部变量表大小和操作数栈大小是由编译器提前计算好的
 * 存储在class文件的method_info结构的Code属性中
 */
type Frame struct {
	lower		*Frame
	localVars	LocalVars //局部变量表
	opStack		*OpStack  //操作数栈
	thread		*Thread
	nextPC		int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame {
		thread: thread,
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

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) NextPC() int {
	return self.nextPC
}
