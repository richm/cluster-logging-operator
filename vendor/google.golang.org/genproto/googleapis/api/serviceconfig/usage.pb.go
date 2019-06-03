// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/usage.proto

package serviceconfig // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configuration controlling usage of a service.
type Usage struct {
	// Requirements that must be satisfied before a consumer project can use the
	// service. Each requirement is of the form <service.name>/<requirement-id>;
	// for example 'serviceusage.googleapis.com/billing-enabled'.
	Requirements []string `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty"`
	// A list of usage rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*UsageRule `protobuf:"bytes,6,rep,name=rules,proto3" json:"rules,omitempty"`
	// The full resource name of a channel used for sending notifications to the
	// service producer.
	//
	// Google Service Management currently only supports
	// [Google Cloud Pub/Sub](https://cloud.google.com/pubsub) as a notification
	// channel. To use Google Cloud Pub/Sub as the channel, this must be the name
	// of a Cloud Pub/Sub topic that uses the Cloud Pub/Sub topic name format
	// documented in https://cloud.google.com/pubsub/docs/overview.
	ProducerNotificationChannel string   `protobuf:"bytes,7,opt,name=producer_notification_channel,json=producerNotificationChannel,proto3" json:"producer_notification_channel,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *Usage) Reset()         { *m = Usage{} }
func (m *Usage) String() string { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()    {}
func (*Usage) Descriptor() ([]byte, []int) {
	return fileDescriptor_usage_5ff3c0173739d6b5, []int{0}
}
func (m *Usage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Usage.Unmarshal(m, b)
}
func (m *Usage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Usage.Marshal(b, m, deterministic)
}
func (dst *Usage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Usage.Merge(dst, src)
}
func (m *Usage) XXX_Size() int {
	return xxx_messageInfo_Usage.Size(m)
}
func (m *Usage) XXX_DiscardUnknown() {
	xxx_messageInfo_Usage.DiscardUnknown(m)
}

var xxx_messageInfo_Usage proto.InternalMessageInfo

func (m *Usage) GetRequirements() []string {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *Usage) GetRules() []*UsageRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Usage) GetProducerNotificationChannel() string {
	if m != nil {
		return m.ProducerNotificationChannel
	}
	return ""
}

// Usage configuration rules for the service.
//
// NOTE: Under development.
//
//
// Use this rule to configure unregistered calls for the service. Unregistered
// calls are calls that do not contain consumer project identity.
// (Example: calls that do not contain an API key).
// By default, API methods do not allow unregistered calls, and each method call
// must be identified by a consumer project identity. Use this rule to
// allow/disallow unregistered calls.
//
// Example of an API that wants to allow unregistered calls for entire service.
//
//     usage:
//       rules:
//       - selector: "*"
//         allow_unregistered_calls: true
//
// Example of a method that wants to allow unregistered calls.
//
//     usage:
//       rules:
//       - selector: "google.example.library.v1.LibraryService.CreateBook"
//         allow_unregistered_calls: true
type UsageRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// If true, the selected method allows unregistered calls, e.g. calls
	// that don't identify any user or application.
	AllowUnregisteredCalls bool `protobuf:"varint,2,opt,name=allow_unregistered_calls,json=allowUnregisteredCalls,proto3" json:"allow_unregistered_calls,omitempty"`
	// If true, the selected method should skip service control and the control
	// plane features, such as quota and billing, will not be available.
	// This flag is used by Google Cloud Endpoints to bypass checks for internal
	// methods, such as service health check methods.
	SkipServiceControl   bool     `protobuf:"varint,3,opt,name=skip_service_control,json=skipServiceControl,proto3" json:"skip_service_control,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsageRule) Reset()         { *m = UsageRule{} }
func (m *UsageRule) String() string { return proto.CompactTextString(m) }
func (*UsageRule) ProtoMessage()    {}
func (*UsageRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_usage_5ff3c0173739d6b5, []int{1}
}
func (m *UsageRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsageRule.Unmarshal(m, b)
}
func (m *UsageRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsageRule.Marshal(b, m, deterministic)
}
func (dst *UsageRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsageRule.Merge(dst, src)
}
func (m *UsageRule) XXX_Size() int {
	return xxx_messageInfo_UsageRule.Size(m)
}
func (m *UsageRule) XXX_DiscardUnknown() {
	xxx_messageInfo_UsageRule.DiscardUnknown(m)
}

var xxx_messageInfo_UsageRule proto.InternalMessageInfo

func (m *UsageRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *UsageRule) GetAllowUnregisteredCalls() bool {
	if m != nil {
		return m.AllowUnregisteredCalls
	}
	return false
}

func (m *UsageRule) GetSkipServiceControl() bool {
	if m != nil {
		return m.SkipServiceControl
	}
	return false
}

func init() {
	proto.RegisterType((*Usage)(nil), "google.api.Usage")
	proto.RegisterType((*UsageRule)(nil), "google.api.UsageRule")
}

func init() { proto.RegisterFile("google/api/usage.proto", fileDescriptor_usage_5ff3c0173739d6b5) }

var fileDescriptor_usage_5ff3c0173739d6b5 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x49, 0xfb, 0xb7, 0x7f, 0x1b, 0xc5, 0x45, 0xd0, 0x12, 0x14, 0x61, 0xe8, 0x6a, 0x40,
	0x98, 0x11, 0xdd, 0x08, 0xae, 0x6c, 0x11, 0x71, 0x23, 0x65, 0xa4, 0x1b, 0x37, 0x43, 0x4c, 0x6f,
	0x63, 0x30, 0xcd, 0x1d, 0x93, 0x8c, 0x3e, 0x88, 0x5b, 0x57, 0x3e, 0xa9, 0x4c, 0xa6, 0xd6, 0x76,
	0x79, 0xef, 0x77, 0xce, 0x49, 0x72, 0x42, 0x47, 0x0a, 0x51, 0x19, 0xc8, 0x45, 0xa5, 0xf3, 0xda,
	0x0b, 0x05, 0x59, 0xe5, 0x30, 0x20, 0xa3, 0xed, 0x3e, 0x13, 0x95, 0x1e, 0x7f, 0x11, 0xda, 0x9b,
	0x37, 0x8c, 0x8d, 0xe9, 0xbe, 0x83, 0xb7, 0x5a, 0x3b, 0x58, 0x81, 0x0d, 0x9e, 0x93, 0xa4, 0x9b,
	0x0e, 0x8b, 0x9d, 0x1d, 0x3b, 0xa3, 0x3d, 0x57, 0x1b, 0xf0, 0xbc, 0x9f, 0x74, 0xd3, 0xbd, 0x8b,
	0xa3, 0xec, 0x2f, 0x29, 0x8b, 0x29, 0x45, 0x6d, 0xa0, 0x68, 0x35, 0x6c, 0x42, 0x4f, 0x2b, 0x87,
	0x8b, 0x5a, 0x82, 0x2b, 0x2d, 0x06, 0xbd, 0xd4, 0x52, 0x04, 0x8d, 0xb6, 0x94, 0x2f, 0xc2, 0x5a,
	0x30, 0xfc, 0x7f, 0x42, 0xd2, 0x61, 0x71, 0xf2, 0x2b, 0x7a, 0xd8, 0xd2, 0x4c, 0x5b, 0xc9, 0xf8,
	0x93, 0xd0, 0xe1, 0x26, 0x98, 0x1d, 0xd3, 0x81, 0x07, 0x03, 0x32, 0xa0, 0xe3, 0x24, 0x9a, 0x37,
	0x33, 0xbb, 0xa2, 0x5c, 0x18, 0x83, 0x1f, 0x65, 0x6d, 0x1d, 0x28, 0xed, 0x03, 0x38, 0x58, 0x94,
	0x52, 0x18, 0xe3, 0x79, 0x27, 0x21, 0xe9, 0xa0, 0x18, 0x45, 0x3e, 0xdf, 0xc2, 0xd3, 0x86, 0xb2,
	0x73, 0x7a, 0xe8, 0x5f, 0x75, 0x55, 0x7a, 0x70, 0xef, 0x5a, 0x42, 0x29, 0xd1, 0x06, 0x87, 0x86,
	0x77, 0xa3, 0x8b, 0x35, 0xec, 0xb1, 0x45, 0xd3, 0x96, 0x4c, 0x0c, 0x3d, 0x90, 0xb8, 0xda, 0x7a,
	0xfc, 0x84, 0xc6, 0x4b, 0xce, 0x9a, 0x7a, 0x67, 0xe4, 0xe9, 0x76, 0x4d, 0x14, 0x1a, 0x61, 0x55,
	0x86, 0x4e, 0xe5, 0x0a, 0x6c, 0x2c, 0x3f, 0x6f, 0x91, 0xa8, 0xb4, 0x8f, 0xff, 0xb2, 0x3e, 0x54,
	0xa2, 0x5d, 0x6a, 0x75, 0xbd, 0x33, 0x7d, 0x77, 0xfe, 0xdd, 0xdd, 0xcc, 0xee, 0x9f, 0xfb, 0xd1,
	0x78, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x64, 0x56, 0x3e, 0xcf, 0x01, 0x00, 0x00,
}