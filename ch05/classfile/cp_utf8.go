package classfile
// import "fmt"
// import "unicode/utf16"

type ConstantUtf8Info struct {
	str string
}
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16()) // todo why transform into u32 -- readBytes(u32)
	bytes := reader.readBytes(length)
	self.str = decodeMTUF8(bytes)
}
func decodeMTUF8(bytes []byte) string {
	return string(bytes) //todo simplified, not compatible for null or complementary-chars
}

