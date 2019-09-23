// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RPC.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// User Information proto.  Included in ConnectionHeader on connection setup
type UserInformation struct {
	EffectiveUser    *string `protobuf:"bytes,1,req,name=effective_user,json=effectiveUser" json:"effective_user,omitempty"`
	RealUser         *string `protobuf:"bytes,2,opt,name=real_user,json=realUser" json:"real_user,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserInformation) Reset()                    { *m = UserInformation{} }
func (m *UserInformation) String() string            { return proto.CompactTextString(m) }
func (*UserInformation) ProtoMessage()               {}
func (*UserInformation) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *UserInformation) GetEffectiveUser() string {
	if m != nil && m.EffectiveUser != nil {
		return *m.EffectiveUser
	}
	return ""
}

func (m *UserInformation) GetRealUser() string {
	if m != nil && m.RealUser != nil {
		return *m.RealUser
	}
	return ""
}

// This is sent on connection setup after the connection preamble is sent.
type ConnectionHeader struct {
	UserInfo    *UserInformation `protobuf:"bytes,1,opt,name=user_info,json=userInfo" json:"user_info,omitempty"`
	ServiceName *string          `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	// Cell block codec we will use sending over optional cell blocks.  Server throws exception
	// if cannot deal.  Null means no codec'ing going on so we are pb all the time (SLOW!!!)
	CellBlockCodecClass *string `protobuf:"bytes,3,opt,name=cell_block_codec_class,json=cellBlockCodecClass" json:"cell_block_codec_class,omitempty"`
	// Compressor we will use if cell block is compressed.  Server will throw exception if not supported.
	// Class must implement hadoop's CompressionCodec Interface.  Can't compress if no codec.
	CellBlockCompressorClass *string      `protobuf:"bytes,4,opt,name=cell_block_compressor_class,json=cellBlockCompressorClass" json:"cell_block_compressor_class,omitempty"`
	VersionInfo              *VersionInfo `protobuf:"bytes,5,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`
	XXX_unrecognized         []byte       `json:"-"`
}

func (m *ConnectionHeader) Reset()                    { *m = ConnectionHeader{} }
func (m *ConnectionHeader) String() string            { return proto.CompactTextString(m) }
func (*ConnectionHeader) ProtoMessage()               {}
func (*ConnectionHeader) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *ConnectionHeader) GetUserInfo() *UserInformation {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *ConnectionHeader) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *ConnectionHeader) GetCellBlockCodecClass() string {
	if m != nil && m.CellBlockCodecClass != nil {
		return *m.CellBlockCodecClass
	}
	return ""
}

func (m *ConnectionHeader) GetCellBlockCompressorClass() string {
	if m != nil && m.CellBlockCompressorClass != nil {
		return *m.CellBlockCompressorClass
	}
	return ""
}

func (m *ConnectionHeader) GetVersionInfo() *VersionInfo {
	if m != nil {
		return m.VersionInfo
	}
	return nil
}

// Optional Cell block Message.  Included in client RequestHeader
type CellBlockMeta struct {
	// Length of the following cell block.  Could calculate it but convenient having it too hand.
	Length           *uint32 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CellBlockMeta) Reset()                    { *m = CellBlockMeta{} }
func (m *CellBlockMeta) String() string            { return proto.CompactTextString(m) }
func (*CellBlockMeta) ProtoMessage()               {}
func (*CellBlockMeta) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

func (m *CellBlockMeta) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

// At the RPC layer, this message is used to carry
// the server side exception to the RPC client.
type ExceptionResponse struct {
	// Class name of the exception thrown from the server
	ExceptionClassName *string `protobuf:"bytes,1,opt,name=exception_class_name,json=exceptionClassName" json:"exception_class_name,omitempty"`
	// Exception stack trace from the server side
	StackTrace *string `protobuf:"bytes,2,opt,name=stack_trace,json=stackTrace" json:"stack_trace,omitempty"`
	// Optional hostname.  Filled in for some exceptions such as region moved
	// where exception gives clue on where the region may have moved.
	Hostname *string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	Port     *int32  `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	// Set if we are NOT to retry on receipt of this exception
	DoNotRetry       *bool  `protobuf:"varint,5,opt,name=do_not_retry,json=doNotRetry" json:"do_not_retry,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ExceptionResponse) Reset()                    { *m = ExceptionResponse{} }
func (m *ExceptionResponse) String() string            { return proto.CompactTextString(m) }
func (*ExceptionResponse) ProtoMessage()               {}
func (*ExceptionResponse) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{3} }

func (m *ExceptionResponse) GetExceptionClassName() string {
	if m != nil && m.ExceptionClassName != nil {
		return *m.ExceptionClassName
	}
	return ""
}

func (m *ExceptionResponse) GetStackTrace() string {
	if m != nil && m.StackTrace != nil {
		return *m.StackTrace
	}
	return ""
}

func (m *ExceptionResponse) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *ExceptionResponse) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ExceptionResponse) GetDoNotRetry() bool {
	if m != nil && m.DoNotRetry != nil {
		return *m.DoNotRetry
	}
	return false
}

// Header sent making a request.
type RequestHeader struct {
	// Monotonically increasing call_id to keep track of RPC requests and their response
	CallId     *uint32   `protobuf:"varint,1,opt,name=call_id,json=callId" json:"call_id,omitempty"`
	TraceInfo  *RPCTInfo `protobuf:"bytes,2,opt,name=trace_info,json=traceInfo" json:"trace_info,omitempty"`
	MethodName *string   `protobuf:"bytes,3,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	// If true, then a pb Message param follows.
	RequestParam *bool `protobuf:"varint,4,opt,name=request_param,json=requestParam" json:"request_param,omitempty"`
	// If present, then an encoded data block follows.
	CellBlockMeta *CellBlockMeta `protobuf:"bytes,5,opt,name=cell_block_meta,json=cellBlockMeta" json:"cell_block_meta,omitempty"`
	// 0 is NORMAL priority.  200 is HIGH.  If no priority, treat it as NORMAL.
	// See HConstants.
	Priority         *uint32 `protobuf:"varint,6,opt,name=priority" json:"priority,omitempty"`
	Timeout          *uint32 `protobuf:"varint,7,opt,name=timeout" json:"timeout,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{4} }

func (m *RequestHeader) GetCallId() uint32 {
	if m != nil && m.CallId != nil {
		return *m.CallId
	}
	return 0
}

func (m *RequestHeader) GetTraceInfo() *RPCTInfo {
	if m != nil {
		return m.TraceInfo
	}
	return nil
}

func (m *RequestHeader) GetMethodName() string {
	if m != nil && m.MethodName != nil {
		return *m.MethodName
	}
	return ""
}

func (m *RequestHeader) GetRequestParam() bool {
	if m != nil && m.RequestParam != nil {
		return *m.RequestParam
	}
	return false
}

func (m *RequestHeader) GetCellBlockMeta() *CellBlockMeta {
	if m != nil {
		return m.CellBlockMeta
	}
	return nil
}

func (m *RequestHeader) GetPriority() uint32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

func (m *RequestHeader) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

type ResponseHeader struct {
	CallId *uint32 `protobuf:"varint,1,opt,name=call_id,json=callId" json:"call_id,omitempty"`
	// If present, then request threw an exception and no response message (else we presume one)
	Exception *ExceptionResponse `protobuf:"bytes,2,opt,name=exception" json:"exception,omitempty"`
	// If present, then an encoded data block follows.
	CellBlockMeta    *CellBlockMeta `protobuf:"bytes,3,opt,name=cell_block_meta,json=cellBlockMeta" json:"cell_block_meta,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{5} }

func (m *ResponseHeader) GetCallId() uint32 {
	if m != nil && m.CallId != nil {
		return *m.CallId
	}
	return 0
}

func (m *ResponseHeader) GetException() *ExceptionResponse {
	if m != nil {
		return m.Exception
	}
	return nil
}

func (m *ResponseHeader) GetCellBlockMeta() *CellBlockMeta {
	if m != nil {
		return m.CellBlockMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInformation)(nil), "pb.UserInformation")
	proto.RegisterType((*ConnectionHeader)(nil), "pb.ConnectionHeader")
	proto.RegisterType((*CellBlockMeta)(nil), "pb.CellBlockMeta")
	proto.RegisterType((*ExceptionResponse)(nil), "pb.ExceptionResponse")
	proto.RegisterType((*RequestHeader)(nil), "pb.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "pb.ResponseHeader")
}

func init() { proto.RegisterFile("RPC.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x55, 0xba, 0x7f, 0xcd, 0x6d, 0xb3, 0xfd, 0xe6, 0xfd, 0x18, 0x51, 0xf7, 0xb0, 0x52, 0x84,
	0x98, 0x40, 0x8a, 0xa6, 0xed, 0x09, 0x09, 0x5e, 0x5a, 0x21, 0x6d, 0x0f, 0x4c, 0x95, 0xb5, 0xf1,
	0x1a, 0xb9, 0xce, 0x6d, 0x1b, 0x2d, 0xb1, 0x83, 0xed, 0x56, 0xec, 0x3b, 0x20, 0xf1, 0xcc, 0x1b,
	0x1f, 0x83, 0x8f, 0x87, 0x6c, 0x27, 0xd9, 0x06, 0x48, 0xf0, 0x96, 0x7b, 0xce, 0xbd, 0xf1, 0x39,
	0xe7, 0x5a, 0x86, 0x90, 0x4e, 0x27, 0x49, 0xa5, 0xa4, 0x91, 0xa4, 0x53, 0xcd, 0x06, 0xd1, 0xb5,
	0x62, 0x3c, 0x17, 0x0b, 0x0f, 0x0d, 0x7a, 0x17, 0x63, 0xa6, 0xd1, 0x17, 0xa3, 0x1b, 0xd8, 0xbb,
	0xd1, 0xa8, 0x2e, 0xc5, 0x5c, 0xaa, 0x92, 0x99, 0x5c, 0x0a, 0xf2, 0x02, 0x76, 0x71, 0x3e, 0x47,
	0x6e, 0xf2, 0x35, 0xa6, 0x2b, 0x8d, 0x2a, 0x0e, 0x86, 0x9d, 0x93, 0x90, 0x46, 0x2d, 0x6a, 0x27,
	0xc8, 0x11, 0x84, 0x0a, 0x59, 0xe1, 0x3b, 0x3a, 0xc3, 0xe0, 0x24, 0xa4, 0x5d, 0x0b, 0x58, 0x72,
	0xf4, 0xb5, 0x03, 0xff, 0x4d, 0xa4, 0x10, 0xb6, 0x5f, 0x8a, 0x0b, 0x64, 0x19, 0x2a, 0x72, 0x0a,
	0xa1, 0x6d, 0x4e, 0x73, 0x31, 0x97, 0x71, 0x30, 0x0c, 0x4e, 0x7a, 0x67, 0x07, 0x49, 0x35, 0x4b,
	0x7e, 0x11, 0x40, 0xbb, 0xab, 0x1a, 0x20, 0xcf, 0xa0, 0xaf, 0x51, 0xad, 0x73, 0x8e, 0xa9, 0x60,
	0x25, 0xd6, 0xc7, 0xf4, 0x6a, 0xec, 0x8a, 0x95, 0x48, 0xce, 0xe1, 0x90, 0x63, 0x51, 0xa4, 0xb3,
	0x42, 0xf2, 0xdb, 0x94, 0xcb, 0x0c, 0x79, 0xca, 0x0b, 0xa6, 0x75, 0xbc, 0xe1, 0x9a, 0x0f, 0x2c,
	0x3b, 0xb6, 0xe4, 0xc4, 0x72, 0x13, 0x4b, 0x91, 0x77, 0x70, 0xf4, 0x68, 0xa8, 0xac, 0x14, 0x6a,
	0x2d, 0x55, 0x3d, 0xb9, 0xe9, 0x26, 0xe3, 0x07, 0x93, 0x4d, 0x83, 0x1f, 0x3f, 0x83, 0xfe, 0x1a,
	0x95, 0xce, 0xa5, 0xf0, 0x5e, 0xb6, 0x9c, 0x97, 0x3d, 0xeb, 0xe5, 0xa3, 0xc7, 0xad, 0x7a, 0xda,
	0x5b, 0xdf, 0x17, 0xa3, 0x97, 0x10, 0x4d, 0x9a, 0xff, 0x7d, 0x40, 0xc3, 0xc8, 0x21, 0x6c, 0x17,
	0x28, 0x16, 0x66, 0xe9, 0xa2, 0x88, 0x68, 0x5d, 0x8d, 0x7e, 0x04, 0xb0, 0xff, 0xfe, 0x33, 0xc7,
	0xca, 0x65, 0x81, 0xba, 0x92, 0x42, 0x23, 0x39, 0x85, 0xff, 0xb1, 0x01, 0xbd, 0x4a, 0x9f, 0x48,
	0xe0, 0xa4, 0x92, 0x96, 0x73, 0x02, 0x5d, 0x30, 0xc7, 0xd0, 0xd3, 0x86, 0xf1, 0xdb, 0xd4, 0x28,
	0xc6, 0x9b, 0xe8, 0xc0, 0x41, 0xf6, 0x3e, 0x20, 0x19, 0x40, 0x77, 0x29, 0xb5, 0x71, 0xbf, 0xf1,
	0x59, 0xb5, 0x35, 0x21, 0xb0, 0x59, 0x49, 0x65, 0x5c, 0x12, 0x5b, 0xd4, 0x7d, 0x93, 0x21, 0xf4,
	0x33, 0x99, 0x0a, 0x69, 0x52, 0x85, 0x46, 0xdd, 0x39, 0xd7, 0x5d, 0x0a, 0x99, 0xbc, 0x92, 0x86,
	0x5a, 0x64, 0xf4, 0xa5, 0x03, 0x11, 0xc5, 0x4f, 0x2b, 0xd4, 0xa6, 0x5e, 0xf9, 0x53, 0xd8, 0xe1,
	0xac, 0x28, 0xd2, 0x3c, 0x6b, 0x5c, 0xda, 0xf2, 0x32, 0x23, 0xaf, 0x01, 0x9c, 0x2e, 0x1f, 0x60,
	0xc7, 0x05, 0xd8, 0xb7, 0x01, 0xd2, 0xe9, 0xe4, 0xda, 0xa5, 0x17, 0x3a, 0xde, 0x5d, 0x83, 0x63,
	0xe8, 0x95, 0x68, 0x96, 0x32, 0x4b, 0x1f, 0x88, 0x05, 0x0f, 0x39, 0xaf, 0xcf, 0x21, 0x52, 0xfe,
	0xdc, 0xb4, 0x62, 0x8a, 0x95, 0x4e, 0x77, 0x97, 0xf6, 0x6b, 0x70, 0x6a, 0x31, 0xf2, 0x06, 0xf6,
	0x1e, 0x2c, 0xbd, 0x44, 0xc3, 0xea, 0xc5, 0xed, 0xdb, 0x73, 0x1f, 0x2d, 0x87, 0x46, 0xfc, 0xd1,
	0xae, 0x06, 0xd0, 0xad, 0x54, 0x2e, 0x55, 0x6e, 0xee, 0xe2, 0x6d, 0xe7, 0xa3, 0xad, 0x49, 0x0c,
	0x3b, 0x26, 0x2f, 0x51, 0xae, 0x4c, 0xbc, 0xe3, 0xa8, 0xa6, 0x1c, 0x7d, 0x0b, 0x60, 0xb7, 0x59,
	0xe0, 0xdf, 0xf2, 0x38, 0x87, 0xb0, 0xdd, 0x61, 0x1d, 0xc7, 0x13, 0x2b, 0xeb, 0xb7, 0x9b, 0x40,
	0xef, 0xfb, 0xfe, 0xe4, 0x68, 0xe3, 0xdf, 0x1c, 0x8d, 0xdf, 0xc2, 0x2b, 0xa9, 0x16, 0x09, 0xab,
	0x18, 0x5f, 0x62, 0xb2, 0x64, 0x99, 0x94, 0x55, 0xb2, 0x9c, 0xb5, 0x0f, 0xc3, 0x6c, 0x35, 0x4f,
	0x16, 0x28, 0x50, 0x31, 0x83, 0xd9, 0xd8, 0x3e, 0x28, 0x53, 0x0b, 0xeb, 0x8b, 0xe0, 0x7b, 0x10,
	0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xde, 0x56, 0xde, 0xc7, 0x61, 0x04, 0x00, 0x00,
}
