package main

import (
	"jvmgo/ch03/classfile"
	"jvmgo/ch03/rtda"
)

/*
调用MemberInfo结构体的CodeAttribute方法可以获取他的Code属性
 */
func interpreter(methodInfo *classfile.MemberInfo) {
	codeAttr:=methodInfo.CodeAttribute()
	maxLocals:=codeAttr.MaxLocals()
	maxStack:=codeAttr.MaxStack()
	bytecode:=codeAttr.Code()

	thread:=rtda.NewThread()
	frame:=thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread,bytecode)
}
