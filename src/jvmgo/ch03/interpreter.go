package main

import (
	"fmt"
	"jvmgo/ch03/instructions"
	"jvmgo/ch03/instructions/base"
	"jvmgo/ch03/rtda"
	"jvmgo/ch03/rtda/heap"
)

/*
调用MemberInfo结构体的CodeAttribute方法可以获取他的Code属性
*/
func interpreter(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}
//打印虚拟机栈信息
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

//每次循环开始，先拿到栈帧，然后根据pc从当前方法中解码一条指令
//指令执行完毕之后，判断Java虚拟机中是否有栈帧，如果没有直接退出循环，否则继续
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPc(pc)

		//decode
		reader.Reset(frame.Method().Code(), pc)
		//获取操作码
		opcode := reader.ReadUint8()
		//通过操作码获取对应的指令
		inst := instructions.NewInstruction(opcode)
		//获取操作数
		inst.FetchOperands(reader)
		//程序计数器下移，指向下一条指令
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame,inst)
		}
		//execute
		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}
}
//方法执行时打印指令信息
func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v #2d %T $v\n", className, methodName, pc, inst, inst)
}
