package classfile
import "fmt"

/**
 * 常量池中的常量分为两类：字面量（literal，包括数字常量和字符串常量）和符号引用（symbolic reference，包括类、接口名、字段、方法信息）
 * 符号引用都是通过索引直接或者间接指向 CONSTANT_Utf8 常量
 */
const (
	CONSTANT_Class					= 7
	CONSTANT_Fieldref				= 9
	CONSTANT_Methodref				= 10
	CONSTANT_InterfaceMethodref		= 11
	CONSTANT_String					= 8
	CONSTANT_Integer				= 3
	CONSTANT_Float					= 4
	CONSTANT_Long					= 5
	CONSTANT_Double					= 6
	CONSTANT_NameAndType			= 12
	CONSTANT_Utf8					= 1
	CONSTANT_MethodHandle			= 15
	CONSTANT_MethodType				= 16
	CONSTANT_InvokeDynamic			= 18
)

type ConstantInfo interface {
	readInfo (reader *ClassReader)
}
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer: return &ConstantIntegerInfo{}
	case CONSTANT_Float: return &ConstantFloatInfo{}
	case CONSTANT_Long: return &ConstantLongInfo{}
	case CONSTANT_Double: return &ConstantDoubleInfo{}		
	case CONSTANT_Utf8: return &ConstantUtf8Info{}		
	case CONSTANT_String: return &ConstantStringInfo{}
	case CONSTANT_Class: return &ConstantClassInfo{}
	case CONSTANT_Fieldref: return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref: return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref: return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType: return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle: return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic: return &ConstantInvokeDynamicInfo{}
	default: 
	fmt.Printf("invalid tag: %d\n", tag)
	panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
