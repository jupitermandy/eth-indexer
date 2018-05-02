// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/maichain/eth-indexer/service/pb/account.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		github.com/maichain/eth-indexer/service/pb/account.proto
		github.com/maichain/eth-indexer/service/pb/block.proto
		github.com/maichain/eth-indexer/service/pb/transaction.proto

	It has these top-level messages:
		GetBalanceRequest
		GetOffsetBalanceRequest
		GetBalanceResponse
		BlockQueryRequest
		BlockQueryResponse
		Block
		TransactionQueryRequest
		TransactionQueryResponse
		Transaction
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetBalanceRequest struct {
	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	BlockNumber int64  `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (m *GetBalanceRequest) Reset()                    { *m = GetBalanceRequest{} }
func (*GetBalanceRequest) ProtoMessage()               {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{0} }

type GetOffsetBalanceRequest struct {
	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Offset  int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (m *GetOffsetBalanceRequest) Reset()                    { *m = GetOffsetBalanceRequest{} }
func (*GetOffsetBalanceRequest) ProtoMessage()               {}
func (*GetOffsetBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{1} }

type GetBalanceResponse struct {
	Amount      string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount"`
	BlockNumber int64  `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
}

func (m *GetBalanceResponse) Reset()                    { *m = GetBalanceResponse{} }
func (*GetBalanceResponse) ProtoMessage()               {}
func (*GetBalanceResponse) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{2} }

func init() {
	proto.RegisterType((*GetBalanceRequest)(nil), "pb.GetBalanceRequest")
	proto.RegisterType((*GetOffsetBalanceRequest)(nil), "pb.GetOffsetBalanceRequest")
	proto.RegisterType((*GetBalanceResponse)(nil), "pb.GetBalanceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccountService service

type AccountServiceClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetOffsetBalance(ctx context.Context, in *GetOffsetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := grpc.Invoke(ctx, "/pb.AccountService/GetBalance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetOffsetBalance(ctx context.Context, in *GetOffsetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := grpc.Invoke(ctx, "/pb.AccountService/GetOffsetBalance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetOffsetBalance(context.Context, *GetOffsetBalanceRequest) (*GetBalanceResponse, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccountService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetOffsetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOffsetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetOffsetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccountService/GetOffsetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetOffsetBalance(ctx, req.(*GetOffsetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _AccountService_GetBalance_Handler,
		},
		{
			MethodName: "GetOffsetBalance",
			Handler:    _AccountService_GetOffsetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/maichain/eth-indexer/service/pb/account.proto",
}

func (m *GetBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if m.BlockNumber != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAccount(dAtA, i, uint64(m.BlockNumber))
	}
	return i, nil
}

func (m *GetOffsetBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetOffsetBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if m.Offset != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAccount(dAtA, i, uint64(m.Offset))
	}
	return i, nil
}

func (m *GetBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Amount)))
		i += copy(dAtA[i:], m.Amount)
	}
	if m.BlockNumber != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAccount(dAtA, i, uint64(m.BlockNumber))
	}
	return i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetBalanceRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovAccount(uint64(m.BlockNumber))
	}
	return n
}

func (m *GetOffsetBalanceRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovAccount(uint64(m.Offset))
	}
	return n
}

func (m *GetBalanceResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovAccount(uint64(m.BlockNumber))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetBalanceRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetBalanceRequest{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`BlockNumber:` + fmt.Sprintf("%v", this.BlockNumber) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetOffsetBalanceRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetOffsetBalanceRequest{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`Offset:` + fmt.Sprintf("%v", this.Offset) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetBalanceResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetBalanceResponse{`,
		`Amount:` + fmt.Sprintf("%v", this.Amount) + `,`,
		`BlockNumber:` + fmt.Sprintf("%v", this.BlockNumber) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAccount(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetBalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetOffsetBalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetOffsetBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetOffsetBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetBalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAccount
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAccount
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAccount(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAccount = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/maichain/eth-indexer/service/pb/account.proto", fileDescriptorAccount)
}

var fileDescriptorAccount = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xee, 0xa4, 0x58, 0x71, 0x14, 0x59, 0x07, 0x5d, 0x43, 0x95, 0xd9, 0x9a, 0xd3, 0x5e, 0x36,
	0x83, 0x2e, 0xa2, 0x20, 0x7b, 0x30, 0x1e, 0xf6, 0xa6, 0x10, 0x7f, 0x80, 0xcc, 0xa4, 0x93, 0x34,
	0x6c, 0x33, 0x13, 0x33, 0x93, 0x45, 0x08, 0x85, 0xe2, 0xc1, 0x3f, 0xe0, 0xc5, 0x83, 0x3f, 0xc6,
	0x63, 0x8f, 0x82, 0x17, 0x4f, 0xc5, 0xa6, 0x9e, 0xfa, 0x2b, 0xa4, 0x33, 0x29, 0xb6, 0xd1, 0x9e,
	0xf6, 0xf4, 0xe6, 0xfb, 0x78, 0xbc, 0xef, 0x7b, 0xdf, 0x1b, 0xf8, 0x3c, 0x49, 0xf5, 0xa8, 0x64,
	0x7e, 0x24, 0x33, 0x92, 0xd1, 0x34, 0x1a, 0xd1, 0x54, 0x10, 0xae, 0x47, 0x27, 0xa9, 0x18, 0xf2,
	0x0f, 0xbc, 0x20, 0x8a, 0x17, 0x97, 0x69, 0xc4, 0x49, 0xce, 0x08, 0x8d, 0x22, 0x59, 0x0a, 0xed,
	0xe7, 0x85, 0xd4, 0x12, 0x39, 0x39, 0xeb, 0x3f, 0x4c, 0xa4, 0x4c, 0xc6, 0x9c, 0xd0, 0x3c, 0x25,
	0x54, 0x08, 0xa9, 0xa9, 0x4e, 0xa5, 0x50, 0xb6, 0xa3, 0x7f, 0xb2, 0x35, 0x3b, 0x91, 0x89, 0x24,
	0x86, 0x66, 0x65, 0x6c, 0x90, 0x01, 0xe6, 0x65, 0xdb, 0xbd, 0x18, 0xde, 0x39, 0xe7, 0x3a, 0xa0,
	0x63, 0x2a, 0x22, 0x1e, 0xf2, 0xf7, 0x25, 0x57, 0x1a, 0xdd, 0x85, 0xd7, 0xb4, 0xbc, 0xe0, 0xc2,
	0x05, 0x03, 0x70, 0x7c, 0x23, 0xb4, 0x00, 0xb9, 0xf0, 0x3a, 0x1d, 0x0e, 0x0b, 0xae, 0x94, 0xeb,
	0x18, 0x7e, 0x03, 0xd1, 0x23, 0x78, 0x8b, 0x8d, 0x65, 0x74, 0xf1, 0x4e, 0x94, 0x19, 0xe3, 0x85,
	0xdb, 0x1d, 0x80, 0xe3, 0x6e, 0x78, 0xd3, 0x70, 0xaf, 0x0d, 0xe5, 0x51, 0x78, 0xff, 0x9c, 0xeb,
	0x37, 0x71, 0xac, 0xae, 0xac, 0x76, 0x08, 0x7b, 0xd2, 0xcc, 0x69, 0x74, 0x1a, 0xe4, 0x65, 0x10,
	0x6d, 0xaf, 0xa2, 0x72, 0x29, 0x14, 0x47, 0x1e, 0xec, 0xd1, 0x6c, 0x9d, 0xa0, 0x1d, 0x1f, 0xc0,
	0xd5, 0xfc, 0xa8, 0x61, 0xc2, 0xa6, 0xa2, 0xd3, 0x96, 0xff, 0xb5, 0x60, 0x37, 0x38, 0x58, 0xcd,
	0x8f, 0x76, 0xf8, 0x9d, 0x8d, 0x9e, 0x7c, 0x75, 0xe0, 0xed, 0x97, 0xf6, 0x38, 0x6f, 0xed, 0xb9,
	0xd0, 0x14, 0x40, 0xf8, 0xd7, 0x02, 0xba, 0xe7, 0xe7, 0xcc, 0xff, 0x27, 0xdd, 0xfe, 0x61, 0x9b,
	0xb6, 0x4e, 0xbd, 0x57, 0x1f, 0x7f, 0xfc, 0xfe, 0xec, 0x9c, 0xa1, 0x17, 0xe4, 0xf2, 0xf1, 0xe6,
	0xec, 0x8a, 0x54, 0xcd, 0xda, 0x13, 0x62, 0xc4, 0x15, 0xa9, 0xb6, 0x2d, 0x4d, 0x88, 0x49, 0x4b,
	0x91, 0xca, 0xd4, 0x09, 0xfa, 0x04, 0xe0, 0x41, 0x3b, 0x68, 0xf4, 0xa0, 0x51, 0xfc, 0x5f, 0xfc,
	0x7b, 0xed, 0x9c, 0x19, 0x3b, 0xcf, 0xd0, 0xd3, 0x3d, 0x76, 0x6c, 0xea, 0xa4, 0xb2, 0xb5, 0x6d,
	0x24, 0x18, 0xcc, 0x16, 0xb8, 0xf3, 0x73, 0x81, 0x3b, 0xd3, 0x1a, 0x83, 0x59, 0x8d, 0xc1, 0xf7,
	0x1a, 0x83, 0x5f, 0x35, 0x06, 0x5f, 0x96, 0xb8, 0xf3, 0x6d, 0x89, 0x01, 0xeb, 0x99, 0x1f, 0x78,
	0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xd3, 0xd4, 0x69, 0x0e, 0x03, 0x00, 0x00,
}
