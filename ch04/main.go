package main
import "fmt"
import "jvmgo/ch04/rtda"

func main()  {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOpStack(frame.OpStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 123412341234)
	vars.SetLong(4, -123412341234)
	vars.SetFloat(6, 1.23412341234)
	vars.SetDouble(7, -1.23412341234)
	vars.SetRef(9, nil)

	fmt.Println(vars.GetInt(0))
	fmt.Println(vars.GetInt(1))
	fmt.Println(vars.GetLong(2))
	fmt.Println(vars.GetLong(4))
	fmt.Println(vars.GetFloat(6))
	fmt.Println(vars.GetDouble(7))
	fmt.Println(vars.GetRef(9))
}

func testOpStack(ops *rtda.OpStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(123412341234)
	ops.PushLong(-123412341234)
	ops.PushFloat(1.23412341234)
	ops.PushDouble(-1.23412341234)
	ops.PushRef(nil)

	fmt.Println(ops.PopRef())
	fmt.Println(ops.PopDouble())
	fmt.Println(ops.PopFloat())
	fmt.Println(ops.PopLong())
	fmt.Println(ops.PopLong())
	fmt.Println(ops.PopInt())
	fmt.Println(ops.PopInt())
}
