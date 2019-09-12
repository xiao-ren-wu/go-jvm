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
	if internedStrs, ok := internedStrings[goStr]; ok {
		return internedStrs
	}
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetFieldValue("value", "[C").(*Object)
	return _utf16ToString(charArr.Chars())
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func _utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}


func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	internedStrings[goStr] = jStr
	return jStr
}
