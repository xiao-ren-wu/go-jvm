# go-jvm
:milky_way: Implementation of Java Virtual Machine through go Language

### class字节码
1. 魔数  
    class字节码的魔数是："0xCAFEBABY"
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
      > ~~~
      > cp_info{
      > 	u1		tag;
      > 	u1		info[];
      > }
      > ~~~

4. 访问类标志

  一个16位的`bitmask`指出class文件定义的是类还是接口，访问级别是`public`还是`private`等等

5. 类和超类索引

访问类标志之后是两个u2类型的常量池索引，分别给出类名和超类。class文件存储的类名类似完全限定名，但是把点换成了斜线，Java语言规范把这种名字叫做：二进制名，因为每个类都有名字，所以thisClass必须是有效的常量池索引。除了java.lang.Object之外，其他类都有超类，所以superClass只有在Object.class中是0，在其他的class文件中必须是有效的常量池索引。

6. 接口索引表

类和超类之后是接口索引表，表中存放的是常量池的索引，给出该类实现的所有接口的名字。classFileTest没有实现接口。

7. 字段和方法表

接口索引表之后是字段表和方法表，分别存储字段和方法信息，字段和方法的基本结构大致相同，仅差别于属性表。

~~~
file_info{
	u2		access_flags;
	u2		name_index;
	u2		descriptor_index;
	u2		attributes_count;
	attribute_info	attributes[attributes_count];
}
~~~

和类一样，方法也有自己的访问标志，访问标志之后是一个常量池索引，给出字段名或者方法名，然后又是一个常量池索引，给出字段或者方法的描述符，最后是属性表，













































