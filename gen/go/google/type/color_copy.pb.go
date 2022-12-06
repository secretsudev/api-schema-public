// Code generated by protoc-gen-go-copy. DO NOT EDIT.
// source: google/type/color.proto

package color

import "google.golang.org/protobuf/types/known/wrapperspb"

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Color) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Color:
		x.Red = v.GetRed()
		x.Green = v.GetGreen()
		x.Blue = v.GetBlue()
		x.Alpha = v.GetAlpha()
	default:
		if v, ok := v.(interface{ GetRed() float32 }); ok {
			x.Red = v.GetRed()
		}
		if v, ok := v.(interface{ GetGreen() float32 }); ok {
			x.Green = v.GetGreen()
		}
		if v, ok := v.(interface{ GetBlue() float32 }); ok {
			x.Blue = v.GetBlue()
		}
		if v, ok := v.(interface{ GetAlpha() *wrapperspb.FloatValue }); ok {
			x.Alpha = v.GetAlpha()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Color) Proto_ShallowClone() (c *Color) {
	if x != nil {
		c = new(Color)
		c.Red = x.Red
		c.Green = x.Green
		c.Blue = x.Blue
		c.Alpha = x.Alpha
	}
	return
}