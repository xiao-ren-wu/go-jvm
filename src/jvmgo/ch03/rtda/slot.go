package rtda

import "jvmgo/ch03/rtda/heap"

/*
局部变量表：
	执行方法所需的局部变量表大小和操作数栈深度是由编译器提前计算好的
存储在class文件中method_info结构的Code属性中
	Java虚拟机规范中，
局部变量表的每个元素至少可以容纳一个int类型或者引用值，
			两个连续的元素至少可以容纳一个long或者double
*/
type Slot struct {
	num int32        //存放整数
	ref *heap.Object //存放引用
}

