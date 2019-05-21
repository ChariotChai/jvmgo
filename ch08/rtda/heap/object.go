package heap

type Object struct {
	class *Class
	data  interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCnt),
	}
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

func (o *Object) InstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}

func (o *Object) GetRefVar(name, descriptor string) *Object {
	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (o *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
