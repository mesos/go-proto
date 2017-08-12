// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quota.proto

/*
Package mesos_v1_quota is a generated protocol buffer package.

It is generated from these files:
	quota.proto

It has these top-level messages:
	QuotaInfo
	QuotaRequest
	QuotaStatus
*/
package mesos_v1_quota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mesos_v1 "github.com/mesos/go-proto/mesos/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TODO(joerg84): Add limits, i.e. upper bounds of resources that a
// role is allowed to use.
type QuotaInfo struct {
	// Quota is granted per role and not per framework, similar to
	// dynamic reservations.
	Role *string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	// Principal which set the quota. Currently only operators can set quotas.
	Principal *string `protobuf:"bytes,2,opt,name=principal" json:"principal,omitempty"`
	// The guarantee that these resources are allocatable for the above role.
	// NOTE: `guarantee.role` should not specify any role except '*',
	// because quota does not reserve specific resources.
	Guarantee        []*mesos_v1.Resource `protobuf:"bytes,3,rep,name=guarantee" json:"guarantee,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *QuotaInfo) Reset()                    { *m = QuotaInfo{} }
func (m *QuotaInfo) String() string            { return proto.CompactTextString(m) }
func (*QuotaInfo) ProtoMessage()               {}
func (*QuotaInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QuotaInfo) GetRole() string {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return ""
}

func (m *QuotaInfo) GetPrincipal() string {
	if m != nil && m.Principal != nil {
		return *m.Principal
	}
	return ""
}

func (m *QuotaInfo) GetGuarantee() []*mesos_v1.Resource {
	if m != nil {
		return m.Guarantee
	}
	return nil
}

// *
// `QuotaRequest` provides a schema for set quota JSON requests.
type QuotaRequest struct {
	// Disables the capacity heuristic check if set to `true`.
	Force *bool `protobuf:"varint,1,opt,name=force,def=0" json:"force,omitempty"`
	// The role for which to set quota.
	Role *string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	// The requested guarantee that these resources will be allocatable for
	// the above role.
	Guarantee        []*mesos_v1.Resource `protobuf:"bytes,3,rep,name=guarantee" json:"guarantee,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *QuotaRequest) Reset()                    { *m = QuotaRequest{} }
func (m *QuotaRequest) String() string            { return proto.CompactTextString(m) }
func (*QuotaRequest) ProtoMessage()               {}
func (*QuotaRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_QuotaRequest_Force bool = false

func (m *QuotaRequest) GetForce() bool {
	if m != nil && m.Force != nil {
		return *m.Force
	}
	return Default_QuotaRequest_Force
}

func (m *QuotaRequest) GetRole() string {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return ""
}

func (m *QuotaRequest) GetGuarantee() []*mesos_v1.Resource {
	if m != nil {
		return m.Guarantee
	}
	return nil
}

// *
// `QuotaStatus` describes the internal representation for the /quota/status
// response.
type QuotaStatus struct {
	// Quotas which are currently set, i.e. known to the master.
	Infos            []*QuotaInfo `protobuf:"bytes,1,rep,name=infos" json:"infos,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *QuotaStatus) Reset()                    { *m = QuotaStatus{} }
func (m *QuotaStatus) String() string            { return proto.CompactTextString(m) }
func (*QuotaStatus) ProtoMessage()               {}
func (*QuotaStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QuotaStatus) GetInfos() []*QuotaInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func init() {
	proto.RegisterType((*QuotaInfo)(nil), "mesos.v1.quota.QuotaInfo")
	proto.RegisterType((*QuotaRequest)(nil), "mesos.v1.quota.QuotaRequest")
	proto.RegisterType((*QuotaStatus)(nil), "mesos.v1.quota.QuotaStatus")
}

func init() { proto.RegisterFile("quota.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x96, 0x20, 0x72, 0x41, 0x0c, 0x16, 0x43, 0x0a, 0x0c, 0x55, 0x59, 0x3a, 0x39,
	0x94, 0x91, 0x81, 0xa1, 0x1b, 0x1b, 0x98, 0x5f, 0x70, 0x8a, 0x2e, 0x25, 0x52, 0xc8, 0x25, 0x3e,
	0xbb, 0xbf, 0x1f, 0xf5, 0x22, 0x25, 0xc0, 0xd6, 0xcd, 0x7e, 0xef, 0x9d, 0xbe, 0xa7, 0x07, 0xf9,
	0x10, 0x39, 0xa0, 0xed, 0x3d, 0x07, 0x36, 0x37, 0xdf, 0x24, 0x2c, 0xf6, 0xb8, 0xb3, 0xaa, 0xde,
	0xdd, 0xea, 0xbf, 0x3c, 0xee, 0xca, 0xd1, 0xd0, 0xd4, 0x86, 0x21, 0xfb, 0x38, 0xd9, 0x6f, 0x5d,
	0xcd, 0xc6, 0xc0, 0x85, 0xe7, 0x96, 0x8a, 0x64, 0x9d, 0x6c, 0x33, 0xa7, 0x6f, 0xf3, 0x00, 0x59,
	0xef, 0x9b, 0xae, 0x6a, 0x7a, 0x6c, 0x8b, 0x85, 0x1a, 0xb3, 0x60, 0x9e, 0x20, 0x3b, 0x44, 0xf4,
	0xd8, 0x05, 0xa2, 0x62, 0xb9, 0x5e, 0x6e, 0xf3, 0x67, 0x63, 0x27, 0xb0, 0x23, 0xe1, 0xe8, 0x2b,
	0x72, 0x73, 0x68, 0x33, 0xc0, 0xb5, 0x02, 0x1d, 0x0d, 0x91, 0x24, 0x98, 0x7b, 0x48, 0x6b, 0xf6,
	0xd5, 0x08, 0xbd, 0x7a, 0x49, 0x6b, 0x6c, 0x85, 0xdc, 0xa8, 0x4d, 0x85, 0x16, 0xbf, 0x0a, 0x9d,
	0x8f, 0x7c, 0x85, 0x5c, 0x91, 0x9f, 0x01, 0x43, 0x14, 0x53, 0x42, 0xda, 0x74, 0x35, 0x4b, 0x91,
	0xe8, 0xf1, 0xca, 0xfe, 0x1d, 0xca, 0x4e, 0x7b, 0xb8, 0x31, 0xb7, 0x7f, 0x84, 0x15, 0xfb, 0x83,
	0xc5, 0x1e, 0xab, 0x2f, 0xfa, 0x97, 0xde, 0x5f, 0xbe, 0x9f, 0x76, 0x94, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x55, 0x0b, 0x2c, 0x31, 0x7c, 0x01, 0x00, 0x00,
}
