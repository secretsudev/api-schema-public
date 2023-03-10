// Code generated by protoc-gen-go-copy. DO NOT EDIT.
// source: google/api/http.proto

package annotations

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Http) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Http:
		x.Rules = v.GetRules()
		x.FullyDecodeReservedExpansion = v.GetFullyDecodeReservedExpansion()
	default:
		if v, ok := v.(interface{ GetRules() []*HttpRule }); ok {
			x.Rules = v.GetRules()
		}
		if v, ok := v.(interface{ GetFullyDecodeReservedExpansion() bool }); ok {
			x.FullyDecodeReservedExpansion = v.GetFullyDecodeReservedExpansion()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Http) Proto_ShallowClone() (c *Http) {
	if x != nil {
		c = new(Http)
		c.Rules = x.Rules
		c.FullyDecodeReservedExpansion = x.FullyDecodeReservedExpansion
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *HttpRule) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *HttpRule:
		x.Selector = v.GetSelector()
		x.Pattern = v.GetPattern()
		x.Body = v.GetBody()
		x.ResponseBody = v.GetResponseBody()
		x.AdditionalBindings = v.GetAdditionalBindings()
	default:
		if v, ok := v.(interface{ GetSelector() string }); ok {
			x.Selector = v.GetSelector()
		}
		if v, ok := v.(interface{ GetPattern() isHttpRule_Pattern }); ok {
			x.Pattern = v.GetPattern()
		} else {
			func() {
				if v, ok := v.(interface{ GetGet() string }); ok {
					var defaultValue string
					if v := v.GetGet(); v != defaultValue {
						x.Pattern = &HttpRule_Get{Get: v}
						return
					}
				}
				if v, ok := v.(interface{ GetPut() string }); ok {
					var defaultValue string
					if v := v.GetPut(); v != defaultValue {
						x.Pattern = &HttpRule_Put{Put: v}
						return
					}
				}
				if v, ok := v.(interface{ GetPost() string }); ok {
					var defaultValue string
					if v := v.GetPost(); v != defaultValue {
						x.Pattern = &HttpRule_Post{Post: v}
						return
					}
				}
				if v, ok := v.(interface{ GetDelete() string }); ok {
					var defaultValue string
					if v := v.GetDelete(); v != defaultValue {
						x.Pattern = &HttpRule_Delete{Delete: v}
						return
					}
				}
				if v, ok := v.(interface{ GetPatch() string }); ok {
					var defaultValue string
					if v := v.GetPatch(); v != defaultValue {
						x.Pattern = &HttpRule_Patch{Patch: v}
						return
					}
				}
				if v, ok := v.(interface{ GetCustom() *CustomHttpPattern }); ok {
					var defaultValue *CustomHttpPattern
					if v := v.GetCustom(); v != defaultValue {
						x.Pattern = &HttpRule_Custom{Custom: v}
						return
					}
				}
			}()
		}
		if v, ok := v.(interface{ GetBody() string }); ok {
			x.Body = v.GetBody()
		}
		if v, ok := v.(interface{ GetResponseBody() string }); ok {
			x.ResponseBody = v.GetResponseBody()
		}
		if v, ok := v.(interface{ GetAdditionalBindings() []*HttpRule }); ok {
			x.AdditionalBindings = v.GetAdditionalBindings()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *HttpRule) Proto_ShallowClone() (c *HttpRule) {
	if x != nil {
		c = new(HttpRule)
		c.Selector = x.Selector
		c.Pattern = x.Pattern
		c.Body = x.Body
		c.ResponseBody = x.ResponseBody
		c.AdditionalBindings = x.AdditionalBindings
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *CustomHttpPattern) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *CustomHttpPattern:
		x.Kind = v.GetKind()
		x.Path = v.GetPath()
	default:
		if v, ok := v.(interface{ GetKind() string }); ok {
			x.Kind = v.GetKind()
		}
		if v, ok := v.(interface{ GetPath() string }); ok {
			x.Path = v.GetPath()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *CustomHttpPattern) Proto_ShallowClone() (c *CustomHttpPattern) {
	if x != nil {
		c = new(CustomHttpPattern)
		c.Kind = x.Kind
		c.Path = x.Path
	}
	return
}
