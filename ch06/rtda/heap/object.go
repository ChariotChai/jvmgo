package heap

type Object struct {
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCnt),
	}
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) Fields() Slots {
	return o.fields
}

func (o *Object) InstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}
