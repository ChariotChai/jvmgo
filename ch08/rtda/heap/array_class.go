package heap

func (self *Class) NewArray(count uint) *Object {
	if !self.IsArray() {
		panic("not an array")
	}
	switch self.Name() {
	case "[Z": //bools?
		return &Object{self, make([]int8, count)}
	case "[B": //bytes
		return &Object{self, make([]int8, count)}
	case "[C": //chars
		return &Object{self, make([]uint16, count)}
	case "[S": //shorts
		return &Object{self, make([]int16, count)}
	case "[I": //ints
		return &Object{self, make([]int32, count)}
	case "[J": //longs
		return &Object{self, make([]int64, count)}
	case "[F": //floats
		return &Object{self, make([]float32, count)}
	case "[D": //doubles
		return &Object{self, make([]float64, count)}
	default:
		return &Object{self, make([]*Object, count)}
	}
}

func (self *Class) IsArray() bool {
	return self.name[0] == '['
}
