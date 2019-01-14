// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/node_def.proto

package framework // import "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"

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

type NodeDef struct {
	// The name given to this operator. Used for naming inputs,
	// logging, visualization, etc.  Unique within a single GraphDef.
	// Must match the regexp "[A-Za-z0-9.][A-Za-z0-9_./]*".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The operation name.  There may be custom parameters in attrs.
	// Op names starting with an underscore are reserved for internal use.
	Op string `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"`
	// Each input is "node:src_output" with "node" being a string name and
	// "src_output" indicating which output tensor to use from "node". If
	// "src_output" is 0 the ":0" suffix can be omitted.  Regular inputs
	// may optionally be followed by control inputs that have the format
	// "^node".
	Input []string `protobuf:"bytes,3,rep,name=input,proto3" json:"input,omitempty"`
	// A (possibly partial) specification for the device on which this
	// node should be placed.
	// The expected syntax for this string is as follows:
	//
	// DEVICE_SPEC ::= PARTIAL_SPEC
	//
	// PARTIAL_SPEC ::= ("/" CONSTRAINT) *
	// CONSTRAINT ::= ("job:" JOB_NAME)
	//              | ("replica:" [1-9][0-9]*)
	//              | ("task:" [1-9][0-9]*)
	//              | ("device:" [A-Za-z]* ":" ([1-9][0-9]* | "*") )
	//
	// Valid values for this string include:
	// * "/job:worker/replica:0/task:1/device:GPU:3"  (full specification)
	// * "/job:worker/device:GPU:3"                   (partial specification)
	// * ""                                    (no specification)
	//
	// If the constraints do not resolve to a single device (or if this
	// field is empty or not present), the runtime will attempt to
	// choose a device automatically.
	Device string `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`
	// Operation-specific graph-construction-time configuration.
	// Note that this should include all attrs defined in the
	// corresponding OpDef, including those with a value matching
	// the default -- this allows the default to change and makes
	// NodeDefs easier to interpret on their own.  However, if
	// an attr with a default is not specified in this list, the
	// default will be used.
	// The "names" (keys) must match the regexp "[a-z][a-z0-9_]+" (and
	// one of the names from the corresponding OpDef's attr field).
	// The values must have a type matching the corresponding OpDef
	// attr's type field.
	// TODO(josh11b): Add some examples here showing best practices.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// This stores debug information associated with the node.
	ExperimentalDebugInfo *NodeDef_ExperimentalDebugInfo `protobuf:"bytes,6,opt,name=experimental_debug_info,json=experimentalDebugInfo,proto3" json:"experimental_debug_info,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                       `json:"-"`
	XXX_unrecognized      []byte                         `json:"-"`
	XXX_sizecache         int32                          `json:"-"`
}

func (m *NodeDef) Reset()         { *m = NodeDef{} }
func (m *NodeDef) String() string { return proto.CompactTextString(m) }
func (*NodeDef) ProtoMessage()    {}
func (*NodeDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_def_1b388980ffdf8b16, []int{0}
}
func (m *NodeDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeDef.Unmarshal(m, b)
}
func (m *NodeDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeDef.Marshal(b, m, deterministic)
}
func (dst *NodeDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDef.Merge(dst, src)
}
func (m *NodeDef) XXX_Size() int {
	return xxx_messageInfo_NodeDef.Size(m)
}
func (m *NodeDef) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDef.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDef proto.InternalMessageInfo

func (m *NodeDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *NodeDef) GetInput() []string {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *NodeDef) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *NodeDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *NodeDef) GetExperimentalDebugInfo() *NodeDef_ExperimentalDebugInfo {
	if m != nil {
		return m.ExperimentalDebugInfo
	}
	return nil
}

type NodeDef_ExperimentalDebugInfo struct {
	// Opaque string inserted into error messages created by the runtime.
	//
	// This is intended to store the list of names of the nodes from the
	// original graph that this node was derived. For example if this node, say
	// C, was result of a fusion of 2 nodes A and B, then 'original_node' would
	// be {A, B}. This information can be used to map errors originating at the
	// current node to some top level source code.
	OriginalNodeNames    []string `protobuf:"bytes,1,rep,name=original_node_names,json=originalNodeNames,proto3" json:"original_node_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeDef_ExperimentalDebugInfo) Reset()         { *m = NodeDef_ExperimentalDebugInfo{} }
func (m *NodeDef_ExperimentalDebugInfo) String() string { return proto.CompactTextString(m) }
func (*NodeDef_ExperimentalDebugInfo) ProtoMessage()    {}
func (*NodeDef_ExperimentalDebugInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_def_1b388980ffdf8b16, []int{0, 1}
}
func (m *NodeDef_ExperimentalDebugInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeDef_ExperimentalDebugInfo.Unmarshal(m, b)
}
func (m *NodeDef_ExperimentalDebugInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeDef_ExperimentalDebugInfo.Marshal(b, m, deterministic)
}
func (dst *NodeDef_ExperimentalDebugInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDef_ExperimentalDebugInfo.Merge(dst, src)
}
func (m *NodeDef_ExperimentalDebugInfo) XXX_Size() int {
	return xxx_messageInfo_NodeDef_ExperimentalDebugInfo.Size(m)
}
func (m *NodeDef_ExperimentalDebugInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDef_ExperimentalDebugInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDef_ExperimentalDebugInfo proto.InternalMessageInfo

func (m *NodeDef_ExperimentalDebugInfo) GetOriginalNodeNames() []string {
	if m != nil {
		return m.OriginalNodeNames
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeDef)(nil), "tensorflow.NodeDef")
	proto.RegisterMapType((map[string]*AttrValue)(nil), "tensorflow.NodeDef.AttrEntry")
	proto.RegisterType((*NodeDef_ExperimentalDebugInfo)(nil), "tensorflow.NodeDef.ExperimentalDebugInfo")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/node_def.proto", fileDescriptor_node_def_1b388980ffdf8b16)
}

var fileDescriptor_node_def_1b388980ffdf8b16 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x14, 0x64, 0x25, 0xdb, 0x45, 0x6b, 0x28, 0xed, 0xb6, 0x6e, 0x17, 0x43, 0x41, 0xf4, 0xa4, 0xb6,
	0x20, 0x51, 0xf7, 0x52, 0x0a, 0x3d, 0xc4, 0xd8, 0x84, 0x5c, 0x8c, 0xd1, 0x21, 0x87, 0x5c, 0x84,
	0x6c, 0x3d, 0x29, 0x8b, 0xa5, 0x7d, 0x62, 0xbd, 0xb2, 0xe3, 0xbf, 0x90, 0x5f, 0x9c, 0x63, 0xd8,
	0xb5, 0x63, 0x8b, 0xa0, 0xdc, 0xde, 0xc7, 0x8c, 0x46, 0x6f, 0x66, 0x69, 0xa0, 0x41, 0x6e, 0x51,
	0xe5, 0x25, 0xee, 0xa3, 0x35, 0x2a, 0x88, 0x72, 0x95, 0x56, 0xb0, 0x47, 0xb5, 0x89, 0x24, 0x66,
	0x90, 0x64, 0x90, 0x87, 0xb5, 0x42, 0x8d, 0x8c, 0x5e, 0x90, 0xe3, 0x9f, 0x6f, 0xb3, 0x52, 0xad,
	0x55, 0xb2, 0x4b, 0xcb, 0x06, 0x8e, 0xbc, 0xef, 0x8f, 0x2e, 0x7d, 0xb7, 0xc0, 0x0c, 0x66, 0x90,
	0x33, 0x46, 0x7b, 0x32, 0xad, 0x80, 0x13, 0x9f, 0x04, 0x5e, 0x6c, 0x6b, 0xf6, 0x9e, 0x3a, 0x58,
	0x73, 0xc7, 0x4e, 0x1c, 0xac, 0xd9, 0x67, 0xda, 0x17, 0xb2, 0x6e, 0x34, 0x77, 0x7d, 0x37, 0xf0,
	0xe2, 0x63, 0xc3, 0xbe, 0xd0, 0x41, 0x06, 0x3b, 0xb1, 0x06, 0xde, 0xb3, 0xc8, 0x53, 0xc7, 0x7e,
	0xd3, 0x9e, 0x51, 0xe4, 0x7d, 0xdf, 0x0d, 0x86, 0x93, 0x6f, 0xe1, 0xe5, 0xc7, 0xc2, 0x93, 0x68,
	0x78, 0xa5, 0xb5, 0x9a, 0x4b, 0xad, 0x0e, 0xb1, 0x85, 0xb2, 0x94, 0x7e, 0x85, 0x87, 0x1a, 0x94,
	0xa8, 0x40, 0xea, 0xb4, 0x4c, 0x32, 0x58, 0x35, 0x45, 0x22, 0x64, 0x8e, 0x7c, 0xe0, 0x93, 0x60,
	0x38, 0xf9, 0xd1, 0xf5, 0x95, 0x79, 0x8b, 0x32, 0x33, 0x8c, 0x1b, 0x99, 0x63, 0x3c, 0x82, 0xae,
	0xf1, 0x78, 0x41, 0xbd, 0xb3, 0x2a, 0xfb, 0x40, 0xdd, 0x0d, 0x1c, 0x4e, 0x37, 0x9b, 0x92, 0xfd,
	0xa2, 0x7d, 0xeb, 0x90, 0xbd, 0x7a, 0x38, 0x19, 0xb5, 0xf5, 0x0c, 0xef, 0xd6, 0x2c, 0xe3, 0x23,
	0xe6, 0x9f, 0xf3, 0x97, 0x8c, 0xaf, 0xe9, 0xa8, 0x53, 0x9f, 0x85, 0xf4, 0x13, 0x2a, 0x51, 0x08,
	0x99, 0x96, 0x89, 0xcd, 0xcb, 0x58, 0xba, 0xe5, 0xc4, 0x5a, 0xf7, 0xf1, 0x65, 0x65, 0x6e, 0x58,
	0x98, 0xc5, 0x54, 0x50, 0x8e, 0xaa, 0x68, 0xeb, 0x9d, 0x93, 0x9b, 0x7a, 0x06, 0xb6, 0x34, 0x99,
	0x2d, 0xc9, 0xdd, 0xff, 0x42, 0xe8, 0xfb, 0x66, 0x15, 0xae, 0xb1, 0x8a, 0x5a, 0x61, 0x77, 0x97,
	0x05, 0xbe, 0x7a, 0x05, 0x4f, 0x84, 0xac, 0x06, 0x36, 0xfe, 0x3f, 0xcf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x2f, 0xbc, 0xb0, 0x71, 0x62, 0x02, 0x00, 0x00,
}