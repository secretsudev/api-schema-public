// Code generated by protoc-gen-go-copy. DO NOT EDIT.
// source: grpc/reflection/v1alpha/reflection.proto

package grpc_reflection_v1alpha

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ServerReflectionRequest) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ServerReflectionRequest:
		x.Host = v.GetHost()
		x.MessageRequest = v.GetMessageRequest()
	default:
		if v, ok := v.(interface{ GetHost() string }); ok {
			x.Host = v.GetHost()
		}
		if v, ok := v.(interface {
			GetMessageRequest() isServerReflectionRequest_MessageRequest
		}); ok {
			x.MessageRequest = v.GetMessageRequest()
		} else {
			func() {
				if v, ok := v.(interface{ GetFileByFilename() string }); ok {
					var defaultValue string
					if v := v.GetFileByFilename(); v != defaultValue {
						x.MessageRequest = &ServerReflectionRequest_FileByFilename{FileByFilename: v}
						return
					}
				}
				if v, ok := v.(interface{ GetFileContainingSymbol() string }); ok {
					var defaultValue string
					if v := v.GetFileContainingSymbol(); v != defaultValue {
						x.MessageRequest = &ServerReflectionRequest_FileContainingSymbol{FileContainingSymbol: v}
						return
					}
				}
				if v, ok := v.(interface{ GetFileContainingExtension() *ExtensionRequest }); ok {
					var defaultValue *ExtensionRequest
					if v := v.GetFileContainingExtension(); v != defaultValue {
						x.MessageRequest = &ServerReflectionRequest_FileContainingExtension{FileContainingExtension: v}
						return
					}
				}
				if v, ok := v.(interface{ GetAllExtensionNumbersOfType() string }); ok {
					var defaultValue string
					if v := v.GetAllExtensionNumbersOfType(); v != defaultValue {
						x.MessageRequest = &ServerReflectionRequest_AllExtensionNumbersOfType{AllExtensionNumbersOfType: v}
						return
					}
				}
				if v, ok := v.(interface{ GetListServices() string }); ok {
					var defaultValue string
					if v := v.GetListServices(); v != defaultValue {
						x.MessageRequest = &ServerReflectionRequest_ListServices{ListServices: v}
						return
					}
				}
			}()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ServerReflectionRequest) Proto_ShallowClone() (c *ServerReflectionRequest) {
	if x != nil {
		c = new(ServerReflectionRequest)
		c.Host = x.Host
		c.MessageRequest = x.MessageRequest
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ExtensionRequest) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ExtensionRequest:
		x.ContainingType = v.GetContainingType()
		x.ExtensionNumber = v.GetExtensionNumber()
	default:
		if v, ok := v.(interface{ GetContainingType() string }); ok {
			x.ContainingType = v.GetContainingType()
		}
		if v, ok := v.(interface{ GetExtensionNumber() rune }); ok {
			x.ExtensionNumber = v.GetExtensionNumber()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ExtensionRequest) Proto_ShallowClone() (c *ExtensionRequest) {
	if x != nil {
		c = new(ExtensionRequest)
		c.ContainingType = x.ContainingType
		c.ExtensionNumber = x.ExtensionNumber
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ServerReflectionResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ServerReflectionResponse:
		x.ValidHost = v.GetValidHost()
		x.OriginalRequest = v.GetOriginalRequest()
		x.MessageResponse = v.GetMessageResponse()
	default:
		if v, ok := v.(interface{ GetValidHost() string }); ok {
			x.ValidHost = v.GetValidHost()
		}
		if v, ok := v.(interface {
			GetOriginalRequest() *ServerReflectionRequest
		}); ok {
			x.OriginalRequest = v.GetOriginalRequest()
		}
		if v, ok := v.(interface {
			GetMessageResponse() isServerReflectionResponse_MessageResponse
		}); ok {
			x.MessageResponse = v.GetMessageResponse()
		} else {
			func() {
				if v, ok := v.(interface {
					GetFileDescriptorResponse() *FileDescriptorResponse
				}); ok {
					var defaultValue *FileDescriptorResponse
					if v := v.GetFileDescriptorResponse(); v != defaultValue {
						x.MessageResponse = &ServerReflectionResponse_FileDescriptorResponse{FileDescriptorResponse: v}
						return
					}
				}
				if v, ok := v.(interface {
					GetAllExtensionNumbersResponse() *ExtensionNumberResponse
				}); ok {
					var defaultValue *ExtensionNumberResponse
					if v := v.GetAllExtensionNumbersResponse(); v != defaultValue {
						x.MessageResponse = &ServerReflectionResponse_AllExtensionNumbersResponse{AllExtensionNumbersResponse: v}
						return
					}
				}
				if v, ok := v.(interface{ GetListServicesResponse() *ListServiceResponse }); ok {
					var defaultValue *ListServiceResponse
					if v := v.GetListServicesResponse(); v != defaultValue {
						x.MessageResponse = &ServerReflectionResponse_ListServicesResponse{ListServicesResponse: v}
						return
					}
				}
				if v, ok := v.(interface{ GetErrorResponse() *ErrorResponse }); ok {
					var defaultValue *ErrorResponse
					if v := v.GetErrorResponse(); v != defaultValue {
						x.MessageResponse = &ServerReflectionResponse_ErrorResponse{ErrorResponse: v}
						return
					}
				}
			}()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ServerReflectionResponse) Proto_ShallowClone() (c *ServerReflectionResponse) {
	if x != nil {
		c = new(ServerReflectionResponse)
		c.ValidHost = x.ValidHost
		c.OriginalRequest = x.OriginalRequest
		c.MessageResponse = x.MessageResponse
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *FileDescriptorResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *FileDescriptorResponse:
		x.FileDescriptorProto = v.GetFileDescriptorProto()
	default:
		if v, ok := v.(interface{ GetFileDescriptorProto() [][]byte }); ok {
			x.FileDescriptorProto = v.GetFileDescriptorProto()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *FileDescriptorResponse) Proto_ShallowClone() (c *FileDescriptorResponse) {
	if x != nil {
		c = new(FileDescriptorResponse)
		c.FileDescriptorProto = x.FileDescriptorProto
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ExtensionNumberResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ExtensionNumberResponse:
		x.BaseTypeName = v.GetBaseTypeName()
		x.ExtensionNumber = v.GetExtensionNumber()
	default:
		if v, ok := v.(interface{ GetBaseTypeName() string }); ok {
			x.BaseTypeName = v.GetBaseTypeName()
		}
		if v, ok := v.(interface{ GetExtensionNumber() []rune }); ok {
			x.ExtensionNumber = v.GetExtensionNumber()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ExtensionNumberResponse) Proto_ShallowClone() (c *ExtensionNumberResponse) {
	if x != nil {
		c = new(ExtensionNumberResponse)
		c.BaseTypeName = x.BaseTypeName
		c.ExtensionNumber = x.ExtensionNumber
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ListServiceResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ListServiceResponse:
		x.Service = v.GetService()
	default:
		if v, ok := v.(interface{ GetService() []*ServiceResponse }); ok {
			x.Service = v.GetService()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ListServiceResponse) Proto_ShallowClone() (c *ListServiceResponse) {
	if x != nil {
		c = new(ListServiceResponse)
		c.Service = x.Service
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ServiceResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ServiceResponse:
		x.Name = v.GetName()
	default:
		if v, ok := v.(interface{ GetName() string }); ok {
			x.Name = v.GetName()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ServiceResponse) Proto_ShallowClone() (c *ServiceResponse) {
	if x != nil {
		c = new(ServiceResponse)
		c.Name = x.Name
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ErrorResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ErrorResponse:
		x.ErrorCode = v.GetErrorCode()
		x.ErrorMessage = v.GetErrorMessage()
	default:
		if v, ok := v.(interface{ GetErrorCode() rune }); ok {
			x.ErrorCode = v.GetErrorCode()
		}
		if v, ok := v.(interface{ GetErrorMessage() string }); ok {
			x.ErrorMessage = v.GetErrorMessage()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ErrorResponse) Proto_ShallowClone() (c *ErrorResponse) {
	if x != nil {
		c = new(ErrorResponse)
		c.ErrorCode = x.ErrorCode
		c.ErrorMessage = x.ErrorMessage
	}
	return
}
