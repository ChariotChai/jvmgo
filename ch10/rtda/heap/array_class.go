package heap

func (self *Class) NewArray(count uint) *Object {
	if !self.IsArray() {
		panic("not an array")
	}
	switch self.Name() {
	case "[Z": //bools?
		return &Object{self, make([]int8, count), nil}
	case "[B": //bytes
		return &Object{self, make([]int8, count), nil}
	case "[C": //chars
		return &Object{self, make([]uint16, count), nil}
	case "[S": //shorts
		return &Object{self, make([]int16, count), nil}
	case "[I": //ints
		return &Object{self, make([]int32, count), nil}
	case "[J": //longs
		return &Object{self, make([]int64, count), nil}
	case "[F": //floats
		return &Object{self, make([]float32, count), nil}
	case "[D": //doubles
		return &Object{self, make([]float64, count), nil}
	default:
		return &Object{self, make([]*Object, count), nil}
	}
}

func (self *Class) IsArray() bool {
	return self.name[0] == '['
}
