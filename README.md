# go-jvm
**:milky_way: Implementation of Java Virtual Machine through go Language**

### class字节码

#### ClassFile结构

~~~c
ClassFile{
	u4					magic;						//魔数
	u2					minor_version;					//次版本号
	u2 					major_version;					//主版本号
	u2					constant_pool_count;				//常量池计数器
	cp_info					constant_pool[constant_pool_count-1]		//常量池
	u2					access_flags;					//访问标志
	u2					this_class;					//类索引
	u2					super_class;					//父类索引
	u2					interfaces_count;				//接口计数器
	u2					interfaces[interfaces_count];			//接口表
	u2					fields_count;					//字段计数器
	field_info				fields[fields_count];				//字段表
	u2					methods_count;					//方法计数器
	method_info				methods[methods_count];				//方法表
	u2					attributes_count;				//属性计数器
	attribute_info				attributes[attributes_count];			//属性表
}
~~~

> 1. 为了描述class文件，Java虚拟机定义了u1,u2,u4三种数据类型来表示1,2,4字节无符号整数，分别对应Go语言中的uint8,uint16和uint32

1. 魔数  
    class字节码的魔数是：`0xCAFEBABY`                 : - ) 这可能就是为啥Java是个咖啡杯的原因了~
2. 版本号  
    魔数之后是class文件的次版本号和主版本号，如果某class字节码文件的主版本号是M，次版本号是m,那么完整的版本号可以表示成“M.m”,次版本号只有在Java1.2之前才用过，之后都是0，每次新版本发布都会在主版本号上加1。
3. 常量池

常量池占据了class文件很大一部分数据，里面存放着各式各样的常量信息包括数字和字符串常量，类和接口名，字段和方法名等等。

常量池实际上也是一个表，但是~

   1. 表头给出的常量池大小比实际大1

   2. 有效的常量池索引是1 ~ n-1 ,0是无效索引，不表示任何常量

   3. `CONSTANT_Long_info`	和 `CONSTANT_Double_info`各占两个位置，如果常量池中存在这两种常量，实际常量池常量数量比n-1还要少。而且1~n-1中某些数字也会变成无效索引。

      >  常量池中存放的信息各不相同，所以每种常量的格式也不相同，常量数据的第一个字节是tag,用来区分常量类型
      >
      >  ~~~
      >  cp_info{
      >  	u1		tag;
      >  	u1		info[];
      >  }
      >  ~~~
      >
      >  常量池小结：
      >
      >  可以把常量池中的常量分成两类：字面量和符号引用，字面量包括数字和字符串常量，符号引用包括类和接口名字段和方法信息等。除了字面量其他常量都是通过索引直接或者间接指向`CONSTANT_Utf8_info`常量。



4. 访问类标志

  一个16位的`bitmask`指出class文件定义的是类还是接口，访问级别是`public`还是`private`等等

5. 类和超类索引

   1. 访问类标志之后是两个u2类型的常量池索引，分别给出类名和超类。
   2. class文件存储的类名类似完全限定名，但是把点换成了斜线，Java语言规范把这种名字叫做：二进制名
   3. 因为每个类都有名字，所以thisClass必须是有效的常量池索引。
   4. 除了java.lang.Object之外，其他类都有超类，所以superClass只有在Object.class中是0，在其他的class文件中必须是有效的常量池索引。
6. 接口索引表

类和超类之后是接口索引表，表中存放的是常量池的索引，给出该类实现的所有接口的名字。

7. 字段和方法表

接口索引表之后是字段表和方法表，分别存储字段和方法信息，字段和方法的基本结构大致相同，仅差别于属性表。

~~~c
file_info{
	u2		access_flags;
	u2		name_index;
	u2		descriptor_index;
	u2		attributes_count;
	attribute_info	attributes[attributes_count];
}
~~~

和类一样，方法也有自己的访问标志，访问标志之后是一个常量池索引，给出字段名或者方法名，然后又是一个常量池索引，给出字段或者方法的描述符，最后是属性表。

### 运行时数据区
java和Go数据类型对应如下：
![](https://github.com/xiao-ren-wu/go-jvm/blob/master/images/data_type.png)

Java虚拟机规范对Java虚拟机栈约束相当宽松，

1. 可以是连续的空间，也可以是不连续的空间；

2. 可以是固定大小，也可以是在运行时动态拓展；如果大小有限制，且执行任务的所需的栈空间超出了这个限制，会抛出`StackOverFlow`，如果是动态拓展的（运行时指定-Xss）,但是内存耗尽，会导致`OutOfMemoryError`

Java的线程私有的运行时数据区如下：

![](https://github.com/xiao-ren-wu/go-jvm/blob/master/images/stack.png)

**说明**

1. 在进行`javac`编译时期,便会确定栈的最大深度，即`maxSize`,`size`表示当前栈的深度。
2. `_top`表示栈顶指针，指向栈顶元素。
3. 这里实现“栈”使用的是链表的数据结构，栈的每一个节点称为“栈帧”，即`Frame`。
4. 栈帧中包含`lower`栈顶节点，`localVars`局部变量表，`operandStack`操作数栈。
5. 操作数栈的大小也是在编译时期就确定的。`size`字段用于记录栈顶位置。
6. 局部变量表和操作数栈都是按照索引访问`[]Slot`的,这个数组的每个元素至少可以容纳一个`int`或者引用值，连续的两个元素可以容纳一个`long`或者`double`









































