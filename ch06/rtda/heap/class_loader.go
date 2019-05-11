package heap

import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class //loaded classes
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return *ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return self.loadNonArrayClass(name) //数组类与普通类不同，其数据不是来自class文件，而是有Java虚拟机在运行期间生成
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name) //找到class文件并读入内存
	class := self.defineClass(data)     //解析class文件，生成虚拟机可用的类数据，并放入方法区
	link(class)                         //进行链接
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry //entry仅用于打印信息
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCnt := len(class.interfaceNames)
	if interfaceCnt > 0 {
		class.interfaces = make([]*Class, interfaceCnt)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

//对类规范进行验证
func verify(class *Class) {

}

//给变量分配空间并赋初始值
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class) //计算实例字段个数
	calcStaticFieldSlotIds(class)   //计算静态字段个数
	allocAndInitStaticVars(class)   //给类变量分配空间并赋初值
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCnt
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCnt = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCnt)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descritor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
