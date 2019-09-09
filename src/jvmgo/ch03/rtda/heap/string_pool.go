package heap

import "unicode/utf16"

/*
用map表示字符串常量池，key是Go字符串，value 是Java字符串
 */
var internedStrings = map[string]*Object{}

/*
JString()函数根据Go字符串返回相应的Java字符串实例，如果Java字符串已经在池中
直接返回即可，否则先把Go字符串（UTF8）转换成（UTF16），然后创建一个Java字符串实例
把他的value变量设置成刚刚转换而来的字符数组，最后把Java字符串放入池中
 */
func JString(loader *ClassLoader, goStr string) *Object {
	if internedStrs,ok:=internedStrings[goStr];ok{
		return internedStrs
	}
	chars:=stringToUTF16(goStr)
	jChars:=&Object{loader.LoadClass("[C"),chars}

	jStr:=loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value","[C",jChars)

	internedStrings[goStr]=jStr
	return jStr
}

func stringToUTF16(s string) []uint16 {
	runes:=[]rune(s)//utf32
	return utf16.Encode(runes)
}




