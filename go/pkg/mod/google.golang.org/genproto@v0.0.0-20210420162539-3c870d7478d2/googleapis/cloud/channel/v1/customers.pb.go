// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: google/cloud/channel/v1/customers.proto

package channel

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	postaladdress "google.golang.org/genproto/googleapis/type/postaladdress"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Entity representing a customer of a reseller or distributor.
type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name of the customer.
	// Format: accounts/{account_id}/customers/{customer_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Name of the organization that the customer entity represents.
	OrgDisplayName string `protobuf:"bytes,2,opt,name=org_display_name,json=orgDisplayName,proto3" json:"org_display_name,omitempty"`
	// Required. Address of the organization of the customer entity.
	// Region and zip codes are required to enforce US laws and embargoes.
	// Valid address lines are required for all customers.
	// Language code is discarded. Use the Customer-level language code to set the
	// customer's language.
	OrgPostalAddress *postaladdress.PostalAddress `protobuf:"bytes,3,opt,name=org_postal_address,json=orgPostalAddress,proto3" json:"org_postal_address,omitempty"`
	// Primary contact info.
	PrimaryContactInfo *ContactInfo `protobuf:"bytes,4,opt,name=primary_contact_info,json=primaryContactInfo,proto3" json:"primary_contact_info,omitempty"`
	// Secondary contact email.
	// Alternate email and primary contact email are required to have different
	// domains if primary contact email is present.
	// When creating admin.google.com accounts, users get notified credentials at
	// this email. This email address is also used as a recovery email.
	AlternateEmail string `protobuf:"bytes,5,opt,name=alternate_email,json=alternateEmail,proto3" json:"alternate_email,omitempty"`
	// Required. Primary domain used by the customer.
	// Domain of primary contact email is required to be same as the provided
	// domain.
	Domain string `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain,omitempty"`
	// Output only. The time at which the customer is created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time at which the customer is updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Customer's cloud_identity_id.
	// Populated only if a Cloud Identity resource exists for this customer.
	CloudIdentityId string `protobuf:"bytes,9,opt,name=cloud_identity_id,json=cloudIdentityId,proto3" json:"cloud_identity_id,omitempty"`
	// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,10,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Output only. Cloud Identity information for the customer.
	// Populated only if a Cloud Identity account exists for this customer.
	CloudIdentityInfo *CloudIdentityInfo `protobuf:"bytes,12,opt,name=cloud_identity_info,json=cloudIdentityInfo,proto3" json:"cloud_identity_info,omitempty"`
	// Cloud Identity ID of the customer's channel partner.
	// Populated only if a channel partner exists for this customer.
	ChannelPartnerId string `protobuf:"bytes,13,opt,name=channel_partner_id,json=channelPartnerId,proto3" json:"channel_partner_id,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_customers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_customers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_customers_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetOrgDisplayName() string {
	if x != nil {
		return x.OrgDisplayName
	}
	return ""
}

func (x *Customer) GetOrgPostalAddress() *postaladdress.PostalAddress {
	if x != nil {
		return x.OrgPostalAddress
	}
	return nil
}

func (x *Customer) GetPrimaryContactInfo() *ContactInfo {
	if x != nil {
		return x.PrimaryContactInfo
	}
	return nil
}

func (x *Customer) GetAlternateEmail() string {
	if x != nil {
		return x.AlternateEmail
	}
	return ""
}

func (x *Customer) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Customer) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Customer) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Customer) GetCloudIdentityId() string {
	if x != nil {
		return x.CloudIdentityId
	}
	return ""
}

func (x *Customer) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *Customer) GetCloudIdentityInfo() *CloudIdentityInfo {
	if x != nil {
		return x.CloudIdentityInfo
	}
	return nil
}

func (x *Customer) GetChannelPartnerId() string {
	if x != nil {
		return x.ChannelPartnerId
	}
	return ""
}

// Contact information for a customer account.
type ContactInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// First name of the contact in the customer account.
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// Last name of the contact in the customer account.
	LastName string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// Output only. Display name of the contact in the customer account.
	// Populated by combining customer first name and last name.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Email of the contact in the customer account.
	// Email is required for entitlements that need creation of admin.google.com
	// accounts. The email will be the username used in credentials to access the
	// admin.google.com account.
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// Optional. Job title of the contact in the customer account.
	Title string `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	// Phone number of the contact in the customer account.
	Phone string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *ContactInfo) Reset() {
	*x = ContactInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_customers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactInfo) ProtoMessage() {}

func (x *ContactInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_customers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactInfo.ProtoReflect.Descriptor instead.
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_customers_proto_rawDescGZIP(), []int{1}
}

func (x *ContactInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ContactInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ContactInfo) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ContactInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ContactInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_google_cloud_channel_v1_customers_proto protoreflect.FileDescriptor

var file_google_cloud_channel_v1_customers_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x06, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x10, 0x6f,
	0x72, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x6f, 0x72,
	0x67, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x6f, 0x72, 0x67, 0x50, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x56, 0x0a, 0x14, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x11, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0d,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x5f, 0x0a, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x3a, 0x52, 0xea, 0x41, 0x4f, 0x0a, 0x24, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x27, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x22, 0xb8, 0x01, 0x0a, 0x0b, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x42, 0x6f, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_channel_v1_customers_proto_rawDescOnce sync.Once
	file_google_cloud_channel_v1_customers_proto_rawDescData = file_google_cloud_channel_v1_customers_proto_rawDesc
)

func file_google_cloud_channel_v1_customers_proto_rawDescGZIP() []byte {
	file_google_cloud_channel_v1_customers_proto_rawDescOnce.Do(func() {
		file_google_cloud_channel_v1_customers_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_channel_v1_customers_proto_rawDescData)
	})
	return file_google_cloud_channel_v1_customers_proto_rawDescData
}

var file_google_cloud_channel_v1_customers_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_channel_v1_customers_proto_goTypes = []interface{}{
	(*Customer)(nil),                    // 0: google.cloud.channel.v1.Customer
	(*ContactInfo)(nil),                 // 1: google.cloud.channel.v1.ContactInfo
	(*postaladdress.PostalAddress)(nil), // 2: google.type.PostalAddress
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
	(*CloudIdentityInfo)(nil),           // 4: google.cloud.channel.v1.CloudIdentityInfo
}
var file_google_cloud_channel_v1_customers_proto_depIdxs = []int32{
	2, // 0: google.cloud.channel.v1.Customer.org_postal_address:type_name -> google.type.PostalAddress
	1, // 1: google.cloud.channel.v1.Customer.primary_contact_info:type_name -> google.cloud.channel.v1.ContactInfo
	3, // 2: google.cloud.channel.v1.Customer.create_time:type_name -> google.protobuf.Timestamp
	3, // 3: google.cloud.channel.v1.Customer.update_time:type_name -> google.protobuf.Timestamp
	4, // 4: google.cloud.channel.v1.Customer.cloud_identity_info:type_name -> google.cloud.channel.v1.CloudIdentityInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_channel_v1_customers_proto_init() }
func file_google_cloud_channel_v1_customers_proto_init() {
	if File_google_cloud_channel_v1_customers_proto != nil {
		return
	}
	file_google_cloud_channel_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_channel_v1_customers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_channel_v1_customers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_channel_v1_customers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_channel_v1_customers_proto_goTypes,
		DependencyIndexes: file_google_cloud_channel_v1_customers_proto_depIdxs,
		MessageInfos:      file_google_cloud_channel_v1_customers_proto_msgTypes,
	}.Build()
	File_google_cloud_channel_v1_customers_proto = out.File
	file_google_cloud_channel_v1_customers_proto_rawDesc = nil
	file_google_cloud_channel_v1_customers_proto_goTypes = nil
	file_google_cloud_channel_v1_customers_proto_depIdxs = nil
}
