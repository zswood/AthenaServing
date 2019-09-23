// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Tracing.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Used to pass through the information necessary to continue
// a trace after an RPC is made. All we need is the traceid
// (so we know the overarching trace this message is a part of), and
// the id of the current span when this message was sent, so we know
// what span caused the new span we will create when this message is received.
type RPCTInfo struct {
	TraceId          *int64 `protobuf:"varint,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	ParentId         *int64 `protobuf:"varint,2,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RPCTInfo) Reset()                    { *m = RPCTInfo{} }
func (m *RPCTInfo) String() string            { return proto.CompactTextString(m) }
func (*RPCTInfo) ProtoMessage()               {}
func (*RPCTInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *RPCTInfo) GetTraceId() int64 {
	if m != nil && m.TraceId != nil {
		return *m.TraceId
	}
	return 0
}

func (m *RPCTInfo) GetParentId() int64 {
	if m != nil && m.ParentId != nil {
		return *m.ParentId
	}
	return 0
}

func init() {
	proto.RegisterType((*RPCTInfo)(nil), "pb.RPCTInfo")
}

func init() { proto.RegisterFile("Tracing.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0d, 0x29, 0x4a, 0x4c,
	0xce, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe2,
	0xe2, 0x08, 0x0a, 0x70, 0x0e, 0xf1, 0xcc, 0x4b, 0xcb, 0x17, 0x92, 0xe4, 0xe2, 0x28, 0x29, 0x4a,
	0x4c, 0x4e, 0x8d, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0x07, 0xf3, 0x3d,
	0x53, 0x84, 0xa4, 0xb9, 0x38, 0x0b, 0x12, 0x8b, 0x52, 0xf3, 0x4a, 0x40, 0x72, 0x4c, 0x60, 0x39,
	0x0e, 0x88, 0x80, 0x67, 0x8a, 0x93, 0x03, 0x97, 0x56, 0x7e, 0x51, 0xba, 0x5e, 0x62, 0x41, 0x62,
	0x72, 0x46, 0xaa, 0x5e, 0x46, 0x62, 0x4a, 0x7e, 0x7e, 0x81, 0x5e, 0x46, 0x52, 0x62, 0x71, 0x2a,
	0xc4, 0xa2, 0xa4, 0xd2, 0x34, 0xbd, 0xf4, 0xd4, 0xbc, 0xd4, 0xa2, 0xc4, 0x92, 0xd4, 0x14, 0x27,
	0x98, 0x23, 0x02, 0x40, 0x52, 0xc5, 0x1e, 0x8c, 0x0b, 0x18, 0x19, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0x30, 0xae, 0x8e, 0x99, 0x00, 0x00, 0x00,
}
