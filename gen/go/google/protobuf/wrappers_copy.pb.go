// Code generated by protoc-gen-go-copy. DO NOT EDIT.
// source: google/protobuf/wrappers.proto

package wrapperspb

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *DoubleValue) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *DoubleValue:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() float64 }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *DoubleValue) Proto_ShallowClone() (c *DoubleValue) {
	if x != nil {
		c = new(DoubleValue)
		c.Value = x.Value
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *FloatValue) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *FloatValue:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() float32 }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *FloatValue) Proto_ShallowClone() (c *FloatValue) {
	if x != nil {
		c = new(FloatValue)
		c.Value = x.Value
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Int64Value) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Int64Value:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() int64 }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Int64Value) Proto_ShallowClone() (c *Int64Value) {
	if x != nil {
		c = new(Int64Value)
		c.Value = x.Value
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *UInt64Value) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *UInt64Value:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() uint64 }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *UInt64Value) Proto_ShallowClone() (c *UInt64Value) {
	if x != nil {
		c = new(UInt64Value)
		c.Value = x.Value
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Int32Value) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Int32Value:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() rune }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Int32Value) Proto_ShallowClone() (c *Int32Value) {
	if x != nil {
		c = new(Int32Value)
		c.Value = x.Value
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *UInt32Value) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *UInt32Value:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() uint32 }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *UInt32Value) Proto_ShallowClone() (c *UInt32Value) {
	if x != nil {
		c = new(UInt32Value)
		c.Value = x.Value
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *BoolValue) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *BoolValue:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() bool }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *BoolValue) Proto_ShallowClone() (c *BoolValue) {
	if x != nil {
		c = new(BoolValue)
		c.Value = x.Value
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *StringValue) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *StringValue:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() string }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *StringValue) Proto_ShallowClone() (c *StringValue) {
	if x != nil {
		c = new(StringValue)
		c.Value = x.Value
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *BytesValue) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *BytesValue:
		x.Value = v.GetValue()
	default:
		if v, ok := v.(interface{ GetValue() []byte }); ok {
			x.Value = v.GetValue()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *BytesValue) Proto_ShallowClone() (c *BytesValue) {
	if x != nil {
		c = new(BytesValue)
		c.Value = x.Value
	}
	return
}
