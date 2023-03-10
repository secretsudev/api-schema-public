// Code generated by protoc-gen-go-copy. DO NOT EDIT.
// source: kentik/cloud_export/v202101beta1/cloud_export.proto

package cloud_export

import "google.golang.org/protobuf/types/known/fieldmaskpb"
import "google.golang.org/protobuf/types/known/wrapperspb"

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *CloudExport) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *CloudExport:
		x.Id = v.GetId()
		x.Type = v.GetType()
		x.Enabled = v.GetEnabled()
		x.Name = v.GetName()
		x.Description = v.GetDescription()
		x.ApiRoot = v.GetApiRoot()
		x.FlowDest = v.GetFlowDest()
		x.PlanId = v.GetPlanId()
		x.CloudProvider = v.GetCloudProvider()
		x.Properties = v.GetProperties()
		x.Bgp = v.GetBgp()
		x.CurrentStatus = v.GetCurrentStatus()
	default:
		if v, ok := v.(interface{ GetId() string }); ok {
			x.Id = v.GetId()
		}
		if v, ok := v.(interface{ GetType() CloudExportType }); ok {
			x.Type = v.GetType()
		}
		if v, ok := v.(interface{ GetEnabled() bool }); ok {
			x.Enabled = v.GetEnabled()
		}
		if v, ok := v.(interface{ GetName() string }); ok {
			x.Name = v.GetName()
		}
		if v, ok := v.(interface{ GetDescription() string }); ok {
			x.Description = v.GetDescription()
		}
		if v, ok := v.(interface{ GetApiRoot() string }); ok {
			x.ApiRoot = v.GetApiRoot()
		}
		if v, ok := v.(interface{ GetFlowDest() string }); ok {
			x.FlowDest = v.GetFlowDest()
		}
		if v, ok := v.(interface{ GetPlanId() string }); ok {
			x.PlanId = v.GetPlanId()
		}
		if v, ok := v.(interface{ GetCloudProvider() string }); ok {
			x.CloudProvider = v.GetCloudProvider()
		}
		if v, ok := v.(interface {
			GetProperties() isCloudExport_Properties
		}); ok {
			x.Properties = v.GetProperties()
		} else {
			func() {
				if v, ok := v.(interface{ GetAws() *AwsProperties }); ok {
					var defaultValue *AwsProperties
					if v := v.GetAws(); v != defaultValue {
						x.Properties = &CloudExport_Aws{Aws: v}
						return
					}
				}
				if v, ok := v.(interface{ GetAzure() *AzureProperties }); ok {
					var defaultValue *AzureProperties
					if v := v.GetAzure(); v != defaultValue {
						x.Properties = &CloudExport_Azure{Azure: v}
						return
					}
				}
				if v, ok := v.(interface{ GetGce() *GceProperties }); ok {
					var defaultValue *GceProperties
					if v := v.GetGce(); v != defaultValue {
						x.Properties = &CloudExport_Gce{Gce: v}
						return
					}
				}
				if v, ok := v.(interface{ GetIbm() *IbmProperties }); ok {
					var defaultValue *IbmProperties
					if v := v.GetIbm(); v != defaultValue {
						x.Properties = &CloudExport_Ibm{Ibm: v}
						return
					}
				}
			}()
		}
		if v, ok := v.(interface{ GetBgp() *BgpProperties }); ok {
			x.Bgp = v.GetBgp()
		}
		if v, ok := v.(interface{ GetCurrentStatus() *Status }); ok {
			x.CurrentStatus = v.GetCurrentStatus()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *CloudExport) Proto_ShallowClone() (c *CloudExport) {
	if x != nil {
		c = new(CloudExport)
		c.Id = x.Id
		c.Type = x.Type
		c.Enabled = x.Enabled
		c.Name = x.Name
		c.Description = x.Description
		c.ApiRoot = x.ApiRoot
		c.FlowDest = x.FlowDest
		c.PlanId = x.PlanId
		c.CloudProvider = x.CloudProvider
		c.Properties = x.Properties
		c.Bgp = x.Bgp
		c.CurrentStatus = x.CurrentStatus
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *BgpProperties) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *BgpProperties:
		x.ApplyBgp = v.GetApplyBgp()
		x.UseBgpDeviceId = v.GetUseBgpDeviceId()
		x.DeviceBgpType = v.GetDeviceBgpType()
	default:
		if v, ok := v.(interface{ GetApplyBgp() bool }); ok {
			x.ApplyBgp = v.GetApplyBgp()
		}
		if v, ok := v.(interface{ GetUseBgpDeviceId() string }); ok {
			x.UseBgpDeviceId = v.GetUseBgpDeviceId()
		}
		if v, ok := v.(interface{ GetDeviceBgpType() string }); ok {
			x.DeviceBgpType = v.GetDeviceBgpType()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *BgpProperties) Proto_ShallowClone() (c *BgpProperties) {
	if x != nil {
		c = new(BgpProperties)
		c.ApplyBgp = x.ApplyBgp
		c.UseBgpDeviceId = x.UseBgpDeviceId
		c.DeviceBgpType = x.DeviceBgpType
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *AwsProperties) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *AwsProperties:
		x.Bucket = v.GetBucket()
		x.IamRoleArn = v.GetIamRoleArn()
		x.Region = v.GetRegion()
		x.DeleteAfterRead = v.GetDeleteAfterRead()
		x.MultipleBuckets = v.GetMultipleBuckets()
	default:
		if v, ok := v.(interface{ GetBucket() string }); ok {
			x.Bucket = v.GetBucket()
		}
		if v, ok := v.(interface{ GetIamRoleArn() string }); ok {
			x.IamRoleArn = v.GetIamRoleArn()
		}
		if v, ok := v.(interface{ GetRegion() string }); ok {
			x.Region = v.GetRegion()
		}
		if v, ok := v.(interface{ GetDeleteAfterRead() bool }); ok {
			x.DeleteAfterRead = v.GetDeleteAfterRead()
		}
		if v, ok := v.(interface{ GetMultipleBuckets() bool }); ok {
			x.MultipleBuckets = v.GetMultipleBuckets()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *AwsProperties) Proto_ShallowClone() (c *AwsProperties) {
	if x != nil {
		c = new(AwsProperties)
		c.Bucket = x.Bucket
		c.IamRoleArn = x.IamRoleArn
		c.Region = x.Region
		c.DeleteAfterRead = x.DeleteAfterRead
		c.MultipleBuckets = x.MultipleBuckets
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *AzureProperties) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *AzureProperties:
		x.Location = v.GetLocation()
		x.ResourceGroup = v.GetResourceGroup()
		x.StorageAccount = v.GetStorageAccount()
		x.SubscriptionId = v.GetSubscriptionId()
		x.SecurityPrincipalEnabled = v.GetSecurityPrincipalEnabled()
	default:
		if v, ok := v.(interface{ GetLocation() string }); ok {
			x.Location = v.GetLocation()
		}
		if v, ok := v.(interface{ GetResourceGroup() string }); ok {
			x.ResourceGroup = v.GetResourceGroup()
		}
		if v, ok := v.(interface{ GetStorageAccount() string }); ok {
			x.StorageAccount = v.GetStorageAccount()
		}
		if v, ok := v.(interface{ GetSubscriptionId() string }); ok {
			x.SubscriptionId = v.GetSubscriptionId()
		}
		if v, ok := v.(interface{ GetSecurityPrincipalEnabled() bool }); ok {
			x.SecurityPrincipalEnabled = v.GetSecurityPrincipalEnabled()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *AzureProperties) Proto_ShallowClone() (c *AzureProperties) {
	if x != nil {
		c = new(AzureProperties)
		c.Location = x.Location
		c.ResourceGroup = x.ResourceGroup
		c.StorageAccount = x.StorageAccount
		c.SubscriptionId = x.SubscriptionId
		c.SecurityPrincipalEnabled = x.SecurityPrincipalEnabled
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *GceProperties) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *GceProperties:
		x.Project = v.GetProject()
		x.Subscription = v.GetSubscription()
	default:
		if v, ok := v.(interface{ GetProject() string }); ok {
			x.Project = v.GetProject()
		}
		if v, ok := v.(interface{ GetSubscription() string }); ok {
			x.Subscription = v.GetSubscription()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *GceProperties) Proto_ShallowClone() (c *GceProperties) {
	if x != nil {
		c = new(GceProperties)
		c.Project = x.Project
		c.Subscription = x.Subscription
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *IbmProperties) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *IbmProperties:
		x.Bucket = v.GetBucket()
	default:
		if v, ok := v.(interface{ GetBucket() string }); ok {
			x.Bucket = v.GetBucket()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *IbmProperties) Proto_ShallowClone() (c *IbmProperties) {
	if x != nil {
		c = new(IbmProperties)
		c.Bucket = x.Bucket
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *Status) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *Status:
		x.Status = v.GetStatus()
		x.ErrorMessage = v.GetErrorMessage()
		x.FlowFound = v.GetFlowFound()
		x.ApiAccess = v.GetApiAccess()
		x.StorageAccountAccess = v.GetStorageAccountAccess()
	default:
		if v, ok := v.(interface{ GetStatus() string }); ok {
			x.Status = v.GetStatus()
		}
		if v, ok := v.(interface{ GetErrorMessage() string }); ok {
			x.ErrorMessage = v.GetErrorMessage()
		}
		if v, ok := v.(interface{ GetFlowFound() *wrapperspb.BoolValue }); ok {
			x.FlowFound = v.GetFlowFound()
		}
		if v, ok := v.(interface{ GetApiAccess() *wrapperspb.BoolValue }); ok {
			x.ApiAccess = v.GetApiAccess()
		}
		if v, ok := v.(interface{ GetStorageAccountAccess() *wrapperspb.BoolValue }); ok {
			x.StorageAccountAccess = v.GetStorageAccountAccess()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *Status) Proto_ShallowClone() (c *Status) {
	if x != nil {
		c = new(Status)
		c.Status = x.Status
		c.ErrorMessage = x.ErrorMessage
		c.FlowFound = x.FlowFound
		c.ApiAccess = x.ApiAccess
		c.StorageAccountAccess = x.StorageAccountAccess
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *CreateCloudExportRequest) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *CreateCloudExportRequest:
		x.Export = v.GetExport()
	default:
		if v, ok := v.(interface{ GetExport() *CloudExport }); ok {
			x.Export = v.GetExport()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *CreateCloudExportRequest) Proto_ShallowClone() (c *CreateCloudExportRequest) {
	if x != nil {
		c = new(CreateCloudExportRequest)
		c.Export = x.Export
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *CreateCloudExportResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *CreateCloudExportResponse:
		x.Export = v.GetExport()
	default:
		if v, ok := v.(interface{ GetExport() *CloudExport }); ok {
			x.Export = v.GetExport()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *CreateCloudExportResponse) Proto_ShallowClone() (c *CreateCloudExportResponse) {
	if x != nil {
		c = new(CreateCloudExportResponse)
		c.Export = x.Export
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ListCloudExportRequest) Proto_ShallowCopy(v interface{}) {
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ListCloudExportRequest) Proto_ShallowClone() (c *ListCloudExportRequest) {
	if x != nil {
		c = new(ListCloudExportRequest)
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *ListCloudExportResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *ListCloudExportResponse:
		x.Exports = v.GetExports()
		x.InvalidExportsCount = v.GetInvalidExportsCount()
	default:
		if v, ok := v.(interface{ GetExports() []*CloudExport }); ok {
			x.Exports = v.GetExports()
		}
		if v, ok := v.(interface{ GetInvalidExportsCount() uint32 }); ok {
			x.InvalidExportsCount = v.GetInvalidExportsCount()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *ListCloudExportResponse) Proto_ShallowClone() (c *ListCloudExportResponse) {
	if x != nil {
		c = new(ListCloudExportResponse)
		c.Exports = x.Exports
		c.InvalidExportsCount = x.InvalidExportsCount
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *GetCloudExportRequest) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *GetCloudExportRequest:
		x.Id = v.GetId()
	default:
		if v, ok := v.(interface{ GetId() string }); ok {
			x.Id = v.GetId()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *GetCloudExportRequest) Proto_ShallowClone() (c *GetCloudExportRequest) {
	if x != nil {
		c = new(GetCloudExportRequest)
		c.Id = x.Id
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *GetCloudExportResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *GetCloudExportResponse:
		x.Export = v.GetExport()
	default:
		if v, ok := v.(interface{ GetExport() *CloudExport }); ok {
			x.Export = v.GetExport()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *GetCloudExportResponse) Proto_ShallowClone() (c *GetCloudExportResponse) {
	if x != nil {
		c = new(GetCloudExportResponse)
		c.Export = x.Export
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *PatchCloudExportRequest) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *PatchCloudExportRequest:
		x.Export = v.GetExport()
		x.Mask = v.GetMask()
	default:
		if v, ok := v.(interface{ GetExport() *CloudExport }); ok {
			x.Export = v.GetExport()
		}
		if v, ok := v.(interface{ GetMask() *fieldmaskpb.FieldMask }); ok {
			x.Mask = v.GetMask()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *PatchCloudExportRequest) Proto_ShallowClone() (c *PatchCloudExportRequest) {
	if x != nil {
		c = new(PatchCloudExportRequest)
		c.Export = x.Export
		c.Mask = x.Mask
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *PatchCloudExportResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *PatchCloudExportResponse:
		x.Export = v.GetExport()
	default:
		if v, ok := v.(interface{ GetExport() *CloudExport }); ok {
			x.Export = v.GetExport()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *PatchCloudExportResponse) Proto_ShallowClone() (c *PatchCloudExportResponse) {
	if x != nil {
		c = new(PatchCloudExportResponse)
		c.Export = x.Export
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *UpdateCloudExportRequest) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *UpdateCloudExportRequest:
		x.Export = v.GetExport()
	default:
		if v, ok := v.(interface{ GetExport() *CloudExport }); ok {
			x.Export = v.GetExport()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *UpdateCloudExportRequest) Proto_ShallowClone() (c *UpdateCloudExportRequest) {
	if x != nil {
		c = new(UpdateCloudExportRequest)
		c.Export = x.Export
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *UpdateCloudExportResponse) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *UpdateCloudExportResponse:
		x.Export = v.GetExport()
	default:
		if v, ok := v.(interface{ GetExport() *CloudExport }); ok {
			x.Export = v.GetExport()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *UpdateCloudExportResponse) Proto_ShallowClone() (c *UpdateCloudExportResponse) {
	if x != nil {
		c = new(UpdateCloudExportResponse)
		c.Export = x.Export
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *DeleteCloudExportRequest) Proto_ShallowCopy(v interface{}) {
	switch v := v.(type) {
	case *DeleteCloudExportRequest:
		x.Id = v.GetId()
	default:
		if v, ok := v.(interface{ GetId() string }); ok {
			x.Id = v.GetId()
		}
	}
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *DeleteCloudExportRequest) Proto_ShallowClone() (c *DeleteCloudExportRequest) {
	if x != nil {
		c = new(DeleteCloudExportRequest)
		c.Id = x.Id
	}
	return
}

// Proto_ShallowCopy copies fields, from v to the receiver, using field getters.
// Note that v is of an arbitrary type, which may implement any number of the
// field getters, which are defined as any methods of the same signature as those
// generated for the receiver type, with a name starting with Get.
func (x *DeleteCloudExportResponse) Proto_ShallowCopy(v interface{}) {
}

// Proto_ShallowClone returns a shallow copy of the receiver or nil if it's nil.
func (x *DeleteCloudExportResponse) Proto_ShallowClone() (c *DeleteCloudExportResponse) {
	if x != nil {
		c = new(DeleteCloudExportResponse)
	}
	return
}
