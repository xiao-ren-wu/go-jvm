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
func interpreter(method *heap.Method) {
	thread:=rtda.NewThread()
	frame:=thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread,method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r:=recover();r!=nil{
		fmt.Printf("LocalVars:%v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		panic(r)
	}
}
func loop(thread *rtda.Thread,bytecode []byte){
	frame:=thread.PopFrame()
	reader:=&base.BytecodeReader{}
	for{
		pc:=frame.NextPC()
		thread.SetPc(pc)

		//decode
		reader.Reset(bytecode,pc)
		//获取操作码
		opcode:=reader.ReadUint8()
		//通过操作码获取对应的指令
		inst:=instructions.NewInstruction(opcode)
		//获取操作数
		inst.FetchOperands(reader)
		//程序计数器下移，指向下一条指令
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc:%2d inst:%T %v\n",pc,inst,inst)
		//execute
		inst.Execute(frame)
	}
}
