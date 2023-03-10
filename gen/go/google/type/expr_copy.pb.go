// Code generated by protoc-gen-go-copy. DO NOT EDIT.
// source: google/type/expr.proto

package expr

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Expr) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Expr:
		x.Expression = v.GetExpression()
		x.Title = v.GetTitle()
		x.Description = v.GetDescription()
		x.Location = v.GetLocation()
	default:
		if v, ok := v.(interface{ GetExpression() string }); ok {
			x.Expression = v.GetExpression()
		}
		if v, ok := v.(interface{ GetTitle() string }); ok {
			x.Title = v.GetTitle()
		}
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetLocation() string }); ok {
			x.Location = v.GetLocation()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Expr) Proto_ShallowClone() (c *Expr) {
	if x != nil {
		c = new(Expr)
		c.Expression = x.Expression
		c.Title = x.Title
		c.Description = x.Description
		c.Location = x.Location
	}
	return
}
