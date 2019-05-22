package heap

import (
	"unicode/utf16"
)

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C]"), chars, nil}
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16 {
	r := []rune(s)
	return utf16.Encode(r)
}

func utf16ToString(s []uint16) string {
	r := utf16.Decode(s)
	return string(r)
}

func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}
