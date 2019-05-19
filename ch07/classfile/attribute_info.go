package classfile
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

/**
  共有23种预定义属性
  第一类：JVM使用，共5种(ConstantValue, Code, Exceptions, StackMapTable, BootstrapMethods)
  第二类：Java类库使用，共12种(InnerClasses, Synthetic<源代码不存在，由编译器生成的类成员>, EnclosingMethod, Signature, RuntimeVisibleAnnotations, RuntimeInvisibleAnnotations, RuntimeVisibleParameterAnnotations, RuntimeInvisibleParameterAnnotations, AnnotationDefault, RuntimeVisibleTypeAnnotations, RuntimeInvisibleTypeAnnotations, MethodParameters)
  第三类：提供给工具使用，是可选的，可以不出现在class文件中(LineNumberTable, LocalVariableTable, Depreciated, SourceDebugExtension, LocalVariableTypeTable)
 */

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code": return &CodeAttribute{cp: cp}
	default: return &UnparsedAttributes{attrName, attrLen, nil}
	}
}
