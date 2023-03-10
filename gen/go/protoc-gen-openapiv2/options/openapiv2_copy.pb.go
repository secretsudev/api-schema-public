// Code generated by protoc-gen-go-copy. DO NOT EDIT.
// source: protoc-gen-openapiv2/options/openapiv2.proto

package options

import "google.golang.org/protobuf/types/known/structpb"

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Swagger) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Swagger:
		x.Swagger = v.GetSwagger()
		x.Info = v.GetInfo()
		x.Host = v.GetHost()
		x.BasePath = v.GetBasePath()
		x.Schemes = v.GetSchemes()
		x.Consumes = v.GetConsumes()
		x.Produces = v.GetProduces()
		x.Responses = v.GetResponses()
		x.SecurityDefinitions = v.GetSecurityDefinitions()
		x.Security = v.GetSecurity()
		x.ExternalDocs = v.GetExternalDocs()
		x.Extensions = v.GetExtensions()
	default:
		if v, ok := v.(interface{ GetSwagger() string }); ok {
			x.Swagger = v.GetSwagger()
		}
		if v, ok := v.(interface{ GetInfo() *Info }); ok {
			x.Info = v.GetInfo()
		}
		if v, ok := v.(interface{ GetHost() string }); ok {
			x.Host = v.GetHost()
		}
		if v, ok := v.(interface{ GetBasePath() string }); ok {
			x.BasePath = v.GetBasePath()
		}
		if v, ok := v.(interface{ GetSchemes() []Scheme }); ok {
			x.Schemes = v.GetSchemes()
		}
		if v, ok := v.(interface{ GetConsumes() []string }); ok {
			x.Consumes = v.GetConsumes()
		}
		if v, ok := v.(interface{ GetProduces() []string }); ok {
			x.Produces = v.GetProduces()
		}
		if v, ok := v.(interface{ GetResponses() map[string]*Response }); ok {
			x.Responses = v.GetResponses()
		}
		if v, ok := v.(interface{ GetSecurityDefinitions() *SecurityDefinitions }); ok {
			x.SecurityDefinitions = v.GetSecurityDefinitions()
		}
		if v, ok := v.(interface{ GetSecurity() []*SecurityRequirement }); ok {
			x.Security = v.GetSecurity()
		}
		if v, ok := v.(interface{ GetExternalDocs() *ExternalDocumentation }); ok {
			x.ExternalDocs = v.GetExternalDocs()
		}
		if v, ok := v.(interface {
			GetExtensions() map[string]*structpb.Value
		}); ok {
			x.Extensions = v.GetExtensions()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Swagger) Proto_ShallowClone() (c *Swagger) {
	if x != nil {
		c = new(Swagger)
		c.Swagger = x.Swagger
		c.Info = x.Info
		c.Host = x.Host
		c.BasePath = x.BasePath
		c.Schemes = x.Schemes
		c.Consumes = x.Consumes
		c.Produces = x.Produces
		c.Responses = x.Responses
		c.SecurityDefinitions = x.SecurityDefinitions
		c.Security = x.Security
		c.ExternalDocs = x.ExternalDocs
		c.Extensions = x.Extensions
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Operation) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Operation:
		x.Tags = v.GetTags()
		x.Summary = v.GetSummary()
		x.Description = v.GetDescription()
		x.ExternalDocs = v.GetExternalDocs()
		x.OperationId = v.GetOperationId()
		x.Consumes = v.GetConsumes()
		x.Produces = v.GetProduces()
		x.Responses = v.GetResponses()
		x.Schemes = v.GetSchemes()
		x.Deprecated = v.GetDeprecated()
		x.Security = v.GetSecurity()
		x.Extensions = v.GetExtensions()
	default:
		if v, ok := v.(interface{ GetTags() []string }); ok {
			x.Tags = v.GetTags()
		}
		if v, ok := v.(interface{ GetSummary() string }); ok {
			x.Summary = v.GetSummary()
		}
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetExternalDocs() *ExternalDocumentation }); ok {
			x.ExternalDocs = v.GetExternalDocs()
		}
		if v, ok := v.(interface{ GetOperationId() string }); ok {
			x.OperationId = v.GetOperationId()
		}
		if v, ok := v.(interface{ GetConsumes() []string }); ok {
			x.Consumes = v.GetConsumes()
		}
		if v, ok := v.(interface{ GetProduces() []string }); ok {
			x.Produces = v.GetProduces()
		}
		if v, ok := v.(interface{ GetResponses() map[string]*Response }); ok {
			x.Responses = v.GetResponses()
		}
		if v, ok := v.(interface{ GetSchemes() []Scheme }); ok {
			x.Schemes = v.GetSchemes()
		}
		if v, ok := v.(interface{ GetDeprecated() bool }); ok {
			x.Deprecated = v.GetDeprecated()
		}
		if v, ok := v.(interface{ GetSecurity() []*SecurityRequirement }); ok {
			x.Security = v.GetSecurity()
		}
		if v, ok := v.(interface {
			GetExtensions() map[string]*structpb.Value
		}); ok {
			x.Extensions = v.GetExtensions()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Operation) Proto_ShallowClone() (c *Operation) {
	if x != nil {
		c = new(Operation)
		c.Tags = x.Tags
		c.Summary = x.Summary
		c.Description = x.Description
		c.ExternalDocs = x.ExternalDocs
		c.OperationId = x.OperationId
		c.Consumes = x.Consumes
		c.Produces = x.Produces
		c.Responses = x.Responses
		c.Schemes = x.Schemes
		c.Deprecated = x.Deprecated
		c.Security = x.Security
		c.Extensions = x.Extensions
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Header) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Header:
		x.Description = v.GetDescription()
		x.Type = v.GetType()
		x.Format = v.GetFormat()
		x.Default = v.GetDefault()
		x.Pattern = v.GetPattern()
	default:
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetType() string }); ok {
			x.Type = v.GetType()
		}
		if v, ok := v.(interface{ GetFormat() string }); ok {
			x.Format = v.GetFormat()
		}
		if v, ok := v.(interface{ GetDefault() string }); ok {
			x.Default = v.GetDefault()
		}
		if v, ok := v.(interface{ GetPattern() string }); ok {
			x.Pattern = v.GetPattern()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Header) Proto_ShallowClone() (c *Header) {
	if x != nil {
		c = new(Header)
		c.Description = x.Description
		c.Type = x.Type
		c.Format = x.Format
		c.Default = x.Default
		c.Pattern = x.Pattern
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Response) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Response:
		x.Description = v.GetDescription()
		x.Schema = v.GetSchema()
		x.Headers = v.GetHeaders()
		x.Examples = v.GetExamples()
		x.Extensions = v.GetExtensions()
	default:
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetSchema() *Schema }); ok {
			x.Schema = v.GetSchema()
		}
		if v, ok := v.(interface{ GetHeaders() map[string]*Header }); ok {
			x.Headers = v.GetHeaders()
		}
		if v, ok := v.(interface{ GetExamples() map[string]string }); ok {
			x.Examples = v.GetExamples()
		}
		if v, ok := v.(interface {
			GetExtensions() map[string]*structpb.Value
		}); ok {
			x.Extensions = v.GetExtensions()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Response) Proto_ShallowClone() (c *Response) {
	if x != nil {
		c = new(Response)
		c.Description = x.Description
		c.Schema = x.Schema
		c.Headers = x.Headers
		c.Examples = x.Examples
		c.Extensions = x.Extensions
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Info) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Info:
		x.Title = v.GetTitle()
		x.Description = v.GetDescription()
		x.TermsOfService = v.GetTermsOfService()
		x.Contact = v.GetContact()
		x.License = v.GetLicense()
		x.Version = v.GetVersion()
		x.Extensions = v.GetExtensions()
	default:
		if v, ok := v.(interface{ GetTitle() string }); ok {
			x.Title = v.GetTitle()
		}
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetTermsOfService() string }); ok {
			x.TermsOfService = v.GetTermsOfService()
		}
		if v, ok := v.(interface{ GetContact() *Contact }); ok {
			x.Contact = v.GetContact()
		}
		if v, ok := v.(interface{ GetLicense() *License }); ok {
			x.License = v.GetLicense()
		}
		if v, ok := v.(interface{ GetVersion() string }); ok {
			x.Version = v.GetVersion()
		}
		if v, ok := v.(interface {
			GetExtensions() map[string]*structpb.Value
		}); ok {
			x.Extensions = v.GetExtensions()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Info) Proto_ShallowClone() (c *Info) {
	if x != nil {
		c = new(Info)
		c.Title = x.Title
		c.Description = x.Description
		c.TermsOfService = x.TermsOfService
		c.Contact = x.Contact
		c.License = x.License
		c.Version = x.Version
		c.Extensions = x.Extensions
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Contact) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Contact:
		x.Name = v.GetName()
		x.Url = v.GetUrl()
		x.Email = v.GetEmail()
	default:
		if v, ok := v.(interface{ GetName() string }); ok {
			x.Name = v.GetName()
		}
		if v, ok := v.(interface{ GetUrl() string }); ok {
			x.Url = v.GetUrl()
		}
		if v, ok := v.(interface{ GetEmail() string }); ok {
			x.Email = v.GetEmail()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Contact) Proto_ShallowClone() (c *Contact) {
	if x != nil {
		c = new(Contact)
		c.Name = x.Name
		c.Url = x.Url
		c.Email = x.Email
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *License) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *License:
		x.Name = v.GetName()
		x.Url = v.GetUrl()
	default:
		if v, ok := v.(interface{ GetName() string }); ok {
			x.Name = v.GetName()
		}
		if v, ok := v.(interface{ GetUrl() string }); ok {
			x.Url = v.GetUrl()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *License) Proto_ShallowClone() (c *License) {
	if x != nil {
		c = new(License)
		c.Name = x.Name
		c.Url = x.Url
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ExternalDocumentation) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ExternalDocumentation:
		x.Description = v.GetDescription()
		x.Url = v.GetUrl()
	default:
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetUrl() string }); ok {
			x.Url = v.GetUrl()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ExternalDocumentation) Proto_ShallowClone() (c *ExternalDocumentation) {
	if x != nil {
		c = new(ExternalDocumentation)
		c.Description = x.Description
		c.Url = x.Url
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Schema) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Schema:
		x.JsonSchema = v.GetJsonSchema()
		x.Discriminator = v.GetDiscriminator()
		x.ReadOnly = v.GetReadOnly()
		x.ExternalDocs = v.GetExternalDocs()
		x.Example = v.GetExample()
	default:
		if v, ok := v.(interface{ GetJsonSchema() *JSONSchema }); ok {
			x.JsonSchema = v.GetJsonSchema()
		}
		if v, ok := v.(interface{ GetDiscriminator() string }); ok {
			x.Discriminator = v.GetDiscriminator()
		}
		if v, ok := v.(interface{ GetReadOnly() bool }); ok {
			x.ReadOnly = v.GetReadOnly()
		}
		if v, ok := v.(interface{ GetExternalDocs() *ExternalDocumentation }); ok {
			x.ExternalDocs = v.GetExternalDocs()
		}
		if v, ok := v.(interface{ GetExample() string }); ok {
			x.Example = v.GetExample()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Schema) Proto_ShallowClone() (c *Schema) {
	if x != nil {
		c = new(Schema)
		c.JsonSchema = x.JsonSchema
		c.Discriminator = x.Discriminator
		c.ReadOnly = x.ReadOnly
		c.ExternalDocs = x.ExternalDocs
		c.Example = x.Example
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *JSONSchema) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *JSONSchema:
		x.Ref = v.GetRef()
		x.Title = v.GetTitle()
		x.Description = v.GetDescription()
		x.Default = v.GetDefault()
		x.ReadOnly = v.GetReadOnly()
		x.Example = v.GetExample()
		x.MultipleOf = v.GetMultipleOf()
		x.Maximum = v.GetMaximum()
		x.ExclusiveMaximum = v.GetExclusiveMaximum()
		x.Minimum = v.GetMinimum()
		x.ExclusiveMinimum = v.GetExclusiveMinimum()
		x.MaxLength = v.GetMaxLength()
		x.MinLength = v.GetMinLength()
		x.Pattern = v.GetPattern()
		x.MaxItems = v.GetMaxItems()
		x.MinItems = v.GetMinItems()
		x.UniqueItems = v.GetUniqueItems()
		x.MaxProperties = v.GetMaxProperties()
		x.MinProperties = v.GetMinProperties()
		x.Required = v.GetRequired()
		x.Array = v.GetArray()
		x.Type = v.GetType()
		x.Format = v.GetFormat()
		x.Enum = v.GetEnum()
	default:
		if v, ok := v.(interface{ GetRef() string }); ok {
			x.Ref = v.GetRef()
		}
		if v, ok := v.(interface{ GetTitle() string }); ok {
			x.Title = v.GetTitle()
		}
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetDefault() string }); ok {
			x.Default = v.GetDefault()
		}
		if v, ok := v.(interface{ GetReadOnly() bool }); ok {
			x.ReadOnly = v.GetReadOnly()
		}
		if v, ok := v.(interface{ GetExample() string }); ok {
			x.Example = v.GetExample()
		}
		if v, ok := v.(interface{ GetMultipleOf() float64 }); ok {
			x.MultipleOf = v.GetMultipleOf()
		}
		if v, ok := v.(interface{ GetMaximum() float64 }); ok {
			x.Maximum = v.GetMaximum()
		}
		if v, ok := v.(interface{ GetExclusiveMaximum() bool }); ok {
			x.ExclusiveMaximum = v.GetExclusiveMaximum()
		}
		if v, ok := v.(interface{ GetMinimum() float64 }); ok {
			x.Minimum = v.GetMinimum()
		}
		if v, ok := v.(interface{ GetExclusiveMinimum() bool }); ok {
			x.ExclusiveMinimum = v.GetExclusiveMinimum()
		}
		if v, ok := v.(interface{ GetMaxLength() uint64 }); ok {
			x.MaxLength = v.GetMaxLength()
		}
		if v, ok := v.(interface{ GetMinLength() uint64 }); ok {
			x.MinLength = v.GetMinLength()
		}
		if v, ok := v.(interface{ GetPattern() string }); ok {
			x.Pattern = v.GetPattern()
		}
		if v, ok := v.(interface{ GetMaxItems() uint64 }); ok {
			x.MaxItems = v.GetMaxItems()
		}
		if v, ok := v.(interface{ GetMinItems() uint64 }); ok {
			x.MinItems = v.GetMinItems()
		}
		if v, ok := v.(interface{ GetUniqueItems() bool }); ok {
			x.UniqueItems = v.GetUniqueItems()
		}
		if v, ok := v.(interface{ GetMaxProperties() uint64 }); ok {
			x.MaxProperties = v.GetMaxProperties()
		}
		if v, ok := v.(interface{ GetMinProperties() uint64 }); ok {
			x.MinProperties = v.GetMinProperties()
		}
		if v, ok := v.(interface{ GetRequired() []string }); ok {
			x.Required = v.GetRequired()
		}
		if v, ok := v.(interface{ GetArray() []string }); ok {
			x.Array = v.GetArray()
		}
		if v, ok := v.(interface {
			GetType() []JSONSchema_JSONSchemaSimpleTypes
		}); ok {
			x.Type = v.GetType()
		}
		if v, ok := v.(interface{ GetFormat() string }); ok {
			x.Format = v.GetFormat()
		}
		if v, ok := v.(interface{ GetEnum() []string }); ok {
			x.Enum = v.GetEnum()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *JSONSchema) Proto_ShallowClone() (c *JSONSchema) {
	if x != nil {
		c = new(JSONSchema)
		c.Ref = x.Ref
		c.Title = x.Title
		c.Description = x.Description
		c.Default = x.Default
		c.ReadOnly = x.ReadOnly
		c.Example = x.Example
		c.MultipleOf = x.MultipleOf
		c.Maximum = x.Maximum
		c.ExclusiveMaximum = x.ExclusiveMaximum
		c.Minimum = x.Minimum
		c.ExclusiveMinimum = x.ExclusiveMinimum
		c.MaxLength = x.MaxLength
		c.MinLength = x.MinLength
		c.Pattern = x.Pattern
		c.MaxItems = x.MaxItems
		c.MinItems = x.MinItems
		c.UniqueItems = x.UniqueItems
		c.MaxProperties = x.MaxProperties
		c.MinProperties = x.MinProperties
		c.Required = x.Required
		c.Array = x.Array
		c.Type = x.Type
		c.Format = x.Format
		c.Enum = x.Enum
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Tag) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Tag:
		x.Description = v.GetDescription()
		x.ExternalDocs = v.GetExternalDocs()
	default:
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetExternalDocs() *ExternalDocumentation }); ok {
			x.ExternalDocs = v.GetExternalDocs()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Tag) Proto_ShallowClone() (c *Tag) {
	if x != nil {
		c = new(Tag)
		c.Description = x.Description
		c.ExternalDocs = x.ExternalDocs
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *SecurityDefinitions) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *SecurityDefinitions:
		x.Security = v.GetSecurity()
	default:
		if v, ok := v.(interface {
			GetSecurity() map[string]*SecurityScheme
		}); ok {
			x.Security = v.GetSecurity()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *SecurityDefinitions) Proto_ShallowClone() (c *SecurityDefinitions) {
	if x != nil {
		c = new(SecurityDefinitions)
		c.Security = x.Security
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *SecurityScheme) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *SecurityScheme:
		x.Type = v.GetType()
		x.Description = v.GetDescription()
		x.Name = v.GetName()
		x.In = v.GetIn()
		x.Flow = v.GetFlow()
		x.AuthorizationUrl = v.GetAuthorizationUrl()
		x.TokenUrl = v.GetTokenUrl()
		x.Scopes = v.GetScopes()
		x.Extensions = v.GetExtensions()
	default:
		if v, ok := v.(interface{ GetType() SecurityScheme_Type }); ok {
			x.Type = v.GetType()
		}
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetName() string }); ok {
			x.Name = v.GetName()
		}
		if v, ok := v.(interface{ GetIn() SecurityScheme_In }); ok {
			x.In = v.GetIn()
		}
		if v, ok := v.(interface{ GetFlow() SecurityScheme_Flow }); ok {
			x.Flow = v.GetFlow()
		}
		if v, ok := v.(interface{ GetAuthorizationUrl() string }); ok {
			x.AuthorizationUrl = v.GetAuthorizationUrl()
		}
		if v, ok := v.(interface{ GetTokenUrl() string }); ok {
			x.TokenUrl = v.GetTokenUrl()
		}
		if v, ok := v.(interface{ GetScopes() *Scopes }); ok {
			x.Scopes = v.GetScopes()
		}
		if v, ok := v.(interface {
			GetExtensions() map[string]*structpb.Value
		}); ok {
			x.Extensions = v.GetExtensions()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *SecurityScheme) Proto_ShallowClone() (c *SecurityScheme) {
	if x != nil {
		c = new(SecurityScheme)
		c.Type = x.Type
		c.Description = x.Description
		c.Name = x.Name
		c.In = x.In
		c.Flow = x.Flow
		c.AuthorizationUrl = x.AuthorizationUrl
		c.TokenUrl = x.TokenUrl
		c.Scopes = x.Scopes
		c.Extensions = x.Extensions
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *SecurityRequirement) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *SecurityRequirement:
		x.SecurityRequirement = v.GetSecurityRequirement()
	default:
		if v, ok := v.(interface {
			GetSecurityRequirement() map[string]*SecurityRequirement_SecurityRequirementValue
		}); ok {
			x.SecurityRequirement = v.GetSecurityRequirement()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *SecurityRequirement) Proto_ShallowClone() (c *SecurityRequirement) {
	if x != nil {
		c = new(SecurityRequirement)
		c.SecurityRequirement = x.SecurityRequirement
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *SecurityRequirement_SecurityRequirementValue) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *SecurityRequirement_SecurityRequirementValue:
		x.Scope = v.GetScope()
	default:
		if v, ok := v.(interface{ GetScope() []string }); ok {
			x.Scope = v.GetScope()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *SecurityRequirement_SecurityRequirementValue) Proto_ShallowClone() (c *SecurityRequirement_SecurityRequirementValue) {
	if x != nil {
		c = new(SecurityRequirement_SecurityRequirementValue)
		c.Scope = x.Scope
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Scopes) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Scopes:
		x.Scope = v.GetScope()
	default:
		if v, ok := v.(interface{ GetScope() map[string]string }); ok {
			x.Scope = v.GetScope()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Scopes) Proto_ShallowClone() (c *Scopes) {
	if x != nil {
		c = new(Scopes)
		c.Scope = x.Scope
	}
	return
}
