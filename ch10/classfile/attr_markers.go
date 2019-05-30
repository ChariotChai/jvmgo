package classfile

type MarkerAttribute struct {}
func (self *MarkerAttribute) readInfo(reader *ClassReader)  {
	//nothing
}
type DepreciatedAttribute struct { MarkerAttribute }
type SyntheticAttrbute struct { MarkerAttribute }
