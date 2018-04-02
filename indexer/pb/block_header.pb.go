// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/maichain/eth-indexer/indexer/pb/block_header.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BlockHeader struct {
	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id,omitempty" gorm:"primary_key" deepequal:"-"`
	Hash        string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty" gorm:"column:hash;size:255"`
	ParentHash  string `protobuf:"bytes,3,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty" gorm:"column:parent_hash;size:255"`
	UncleHash   string `protobuf:"bytes,4,opt,name=uncle_hash,json=uncleHash,proto3" json:"uncle_hash,omitempty" gorm:"column:uncle_hash;size:255"`
	Coinbase    string `protobuf:"bytes,5,opt,name=coinbase,proto3" json:"coinbase,omitempty" gorm:"column:coinbase;size:255"`
	Root        string `protobuf:"bytes,6,opt,name=root,proto3" json:"root,omitempty" gorm:"column:root;size:255"`
	TxHash      string `protobuf:"bytes,7,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" gorm:"column:tx_hash;size:255"`
	ReceiptHash string `protobuf:"bytes,8,opt,name=receipt_hash,json=receiptHash,proto3" json:"receipt_hash,omitempty" gorm:"column:receipt_hash;size:255"`
	Bloom       []byte `protobuf:"bytes,9,opt,name=bloom,proto3" json:"bloom,omitempty" gorm:"column:bloom"`
	Difficulty  int64  `protobuf:"varint,10,opt,name=difficulty,proto3" json:"difficulty,omitempty" gorm:"column:difficulty"`
	Number      int64  `protobuf:"varint,11,opt,name=number,proto3" json:"number,omitempty" gorm:"column:number"`
	GasLimit    uint64 `protobuf:"varint,12,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty" gorm:"column:gas_limit"`
	GasUsed     uint64 `protobuf:"varint,13,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty" gorm:"column:gas_used"`
	Time        uint64 `protobuf:"varint,14,opt,name=time,proto3" json:"time,omitempty" gorm:"column:time"`
	ExtraData   []byte `protobuf:"bytes,15,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty" gorm:"column:extra_data"`
	MixDigest   string `protobuf:"bytes,16,opt,name=mix_digest,json=mixDigest,proto3" json:"mix_digest,omitempty" gorm:"column:mix_digest"`
	Nonce       []byte `protobuf:"bytes,17,opt,name=nonce,proto3" json:"nonce,omitempty" gorm:"column:nonce"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlockHeader, []int{0} }

func init() {
	proto.RegisterType((*BlockHeader)(nil), "pb.BlockHeader")
}
func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.ID))
	}
	if len(m.Hash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if len(m.ParentHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.ParentHash)))
		i += copy(dAtA[i:], m.ParentHash)
	}
	if len(m.UncleHash) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.UncleHash)))
		i += copy(dAtA[i:], m.UncleHash)
	}
	if len(m.Coinbase) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Coinbase)))
		i += copy(dAtA[i:], m.Coinbase)
	}
	if len(m.Root) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Root)))
		i += copy(dAtA[i:], m.Root)
	}
	if len(m.TxHash) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.TxHash)))
		i += copy(dAtA[i:], m.TxHash)
	}
	if len(m.ReceiptHash) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.ReceiptHash)))
		i += copy(dAtA[i:], m.ReceiptHash)
	}
	if len(m.Bloom) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Bloom)))
		i += copy(dAtA[i:], m.Bloom)
	}
	if m.Difficulty != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.Difficulty))
	}
	if m.Number != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.Number))
	}
	if m.GasLimit != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.GasLimit))
	}
	if m.GasUsed != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.GasUsed))
	}
	if m.Time != 0 {
		dAtA[i] = 0x70
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.Time))
	}
	if len(m.ExtraData) > 0 {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.ExtraData)))
		i += copy(dAtA[i:], m.ExtraData)
	}
	if len(m.MixDigest) > 0 {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.MixDigest)))
		i += copy(dAtA[i:], m.MixDigest)
	}
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	return i, nil
}

func encodeVarintBlockHeader(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BlockHeader) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovBlockHeader(uint64(m.ID))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.UncleHash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.Coinbase)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.ReceiptHash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	if m.Difficulty != 0 {
		n += 1 + sovBlockHeader(uint64(m.Difficulty))
	}
	if m.Number != 0 {
		n += 1 + sovBlockHeader(uint64(m.Number))
	}
	if m.GasLimit != 0 {
		n += 1 + sovBlockHeader(uint64(m.GasLimit))
	}
	if m.GasUsed != 0 {
		n += 1 + sovBlockHeader(uint64(m.GasUsed))
	}
	if m.Time != 0 {
		n += 1 + sovBlockHeader(uint64(m.Time))
	}
	l = len(m.ExtraData)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.MixDigest)
	if l > 0 {
		n += 2 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 2 + l + sovBlockHeader(uint64(l))
	}
	return n
}

func sovBlockHeader(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBlockHeader(x uint64) (n int) {
	return sovBlockHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *BlockHeader) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BlockHeader{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`ParentHash:` + fmt.Sprintf("%v", this.ParentHash) + `,`,
		`UncleHash:` + fmt.Sprintf("%v", this.UncleHash) + `,`,
		`Coinbase:` + fmt.Sprintf("%v", this.Coinbase) + `,`,
		`Root:` + fmt.Sprintf("%v", this.Root) + `,`,
		`TxHash:` + fmt.Sprintf("%v", this.TxHash) + `,`,
		`ReceiptHash:` + fmt.Sprintf("%v", this.ReceiptHash) + `,`,
		`Bloom:` + fmt.Sprintf("%v", this.Bloom) + `,`,
		`Difficulty:` + fmt.Sprintf("%v", this.Difficulty) + `,`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`GasLimit:` + fmt.Sprintf("%v", this.GasLimit) + `,`,
		`GasUsed:` + fmt.Sprintf("%v", this.GasUsed) + `,`,
		`Time:` + fmt.Sprintf("%v", this.Time) + `,`,
		`ExtraData:` + fmt.Sprintf("%v", this.ExtraData) + `,`,
		`MixDigest:` + fmt.Sprintf("%v", this.MixDigest) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBlockHeader(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockHeader
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
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
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
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
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
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UncleHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
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
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UncleHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coinbase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
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
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coinbase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
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
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
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
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiptHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
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
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiptHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = append(m.Bloom[:0], dAtA[iNdEx:postIndex]...)
			if m.Bloom == nil {
				m.Bloom = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Difficulty", wireType)
			}
			m.Difficulty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Difficulty |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtraData = append(m.ExtraData[:0], dAtA[iNdEx:postIndex]...)
			if m.ExtraData == nil {
				m.ExtraData = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MixDigest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
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
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MixDigest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlockHeader
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
func skipBlockHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockHeader
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
					return 0, ErrIntOverflowBlockHeader
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
					return 0, ErrIntOverflowBlockHeader
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
				return 0, ErrInvalidLengthBlockHeader
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBlockHeader
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
				next, err := skipBlockHeader(dAtA[start:])
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
	ErrInvalidLengthBlockHeader = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockHeader   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/maichain/eth-indexer/indexer/pb/block_header.proto", fileDescriptorBlockHeader)
}

var fileDescriptorBlockHeader = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xeb, 0x34, 0x4d, 0x9b, 0x4d, 0x9e, 0xbe, 0xf8, 0x29, 0xc5, 0x02, 0xc9, 0x0e, 0xcb,
	0x25, 0x42, 0x6d, 0x03, 0x85, 0x0a, 0x11, 0xa8, 0x80, 0x28, 0x48, 0xad, 0x40, 0x02, 0x19, 0x7a,
	0x43, 0x0a, 0x6b, 0x7b, 0xeb, 0xac, 0x1a, 0xbf, 0x60, 0xaf, 0xa5, 0x84, 0x13, 0x1f, 0x83, 0x8f,
	0xc3, 0xb1, 0x47, 0x4e, 0x88, 0x93, 0x45, 0xd3, 0x9b, 0x8f, 0x7c, 0x02, 0xb4, 0xe3, 0xa4, 0xf6,
	0x0a, 0x29, 0xa7, 0xec, 0xfc, 0xe7, 0x3f, 0xbf, 0xc9, 0xce, 0x4e, 0x82, 0x8e, 0x5c, 0xc6, 0x87,
	0x89, 0xb5, 0x6f, 0x07, 0x5e, 0xc7, 0x23, 0xcc, 0x1e, 0x12, 0xe6, 0x77, 0x28, 0x1f, 0xee, 0x31,
	0xdf, 0xa1, 0x63, 0x1a, 0x75, 0xe6, 0x9f, 0xa1, 0xd5, 0xb1, 0x46, 0x81, 0x7d, 0x3e, 0x18, 0x52,
	0xe2, 0xd0, 0x68, 0x3f, 0x8c, 0x02, 0x1e, 0xa8, 0x95, 0xd0, 0xba, 0xb5, 0x57, 0x42, 0xb8, 0x81,
	0x1b, 0x74, 0x20, 0x65, 0x25, 0x67, 0x10, 0x41, 0x00, 0xa7, 0xbc, 0x04, 0xff, 0x44, 0xa8, 0xd1,
	0x13, 0xa4, 0x63, 0x00, 0xa9, 0x7d, 0x54, 0x39, 0xe9, 0x6b, 0x4a, 0x4b, 0x69, 0x57, 0x7b, 0x8f,
	0xb2, 0xd4, 0x68, 0x32, 0x67, 0x37, 0xf0, 0x18, 0xa7, 0x5e, 0xc8, 0x27, 0x7f, 0x52, 0xa3, 0xe5,
	0x06, 0x91, 0xd7, 0xc5, 0x61, 0xc4, 0x3c, 0x12, 0x4d, 0x06, 0xe7, 0x74, 0x82, 0x5b, 0x0e, 0xa5,
	0x21, 0xfd, 0x9c, 0x90, 0x51, 0x17, 0xef, 0x61, 0xb3, 0x72, 0xd2, 0x57, 0x5f, 0xa1, 0xea, 0x90,
	0xc4, 0x43, 0xad, 0xd2, 0x52, 0xda, 0xf5, 0xde, 0x83, 0x2c, 0x35, 0xd6, 0x45, 0x2c, 0x91, 0x6e,
	0xe7, 0x24, 0x3b, 0x18, 0x25, 0x9e, 0xdf, 0x15, 0xe9, 0xa7, 0x31, 0xfb, 0x42, 0xbb, 0x07, 0x87,
	0x87, 0xd8, 0x84, 0x72, 0xf5, 0x13, 0x6a, 0x84, 0x24, 0xa2, 0x3e, 0x1f, 0x00, 0x6d, 0x19, 0x68,
	0xcf, 0xb3, 0xd4, 0xb8, 0x51, 0x92, 0x25, 0x28, 0x96, 0xa0, 0x25, 0x57, 0x89, 0x8d, 0x72, 0xf9,
	0x58, 0x74, 0xf8, 0x88, 0x50, 0xe2, 0xdb, 0x23, 0x9a, 0x37, 0xa8, 0x42, 0x83, 0xa3, 0x2c, 0x35,
	0xb6, 0x0b, 0x55, 0xe2, 0xdf, 0x91, 0xf8, 0x85, 0xa9, 0x84, 0xaf, 0x83, 0x0a, 0xf4, 0x53, 0xb4,
	0x66, 0x07, 0xcc, 0xb7, 0x48, 0x4c, 0xb5, 0x15, 0x60, 0x3f, 0xc9, 0x52, 0x43, 0x9d, 0x6b, 0x12,
	0xd9, 0x90, 0xc8, 0x73, 0x4b, 0x89, 0x7b, 0x8d, 0x12, 0xd3, 0x8d, 0x82, 0x80, 0x6b, 0xb5, 0x62,
	0xba, 0x22, 0x5e, 0x30, 0x5d, 0x91, 0x2e, 0x4f, 0x57, 0xc4, 0xea, 0x3b, 0xb4, 0xca, 0xc7, 0xf9,
	0xc5, 0x57, 0x81, 0xf4, 0x38, 0x4b, 0x8d, 0xad, 0x99, 0x24, 0xc1, 0x74, 0x09, 0x36, 0x73, 0x94,
	0x78, 0x35, 0x3e, 0x86, 0xfb, 0x3a, 0xa8, 0x19, 0x51, 0x9b, 0xb2, 0x70, 0xf6, 0x60, 0x6b, 0x80,
	0x7d, 0x99, 0xa5, 0xc6, 0x4e, 0x59, 0x97, 0xd8, 0x77, 0xe5, 0x2f, 0x5a, 0xb2, 0x95, 0x1a, 0x34,
	0x66, 0x3a, 0x74, 0x79, 0x81, 0x56, 0xac, 0x51, 0x10, 0x78, 0x5a, 0xbd, 0xa5, 0xb4, 0x9b, 0xbd,
	0x7b, 0x59, 0x6a, 0x6c, 0x80, 0x20, 0x71, 0xff, 0x97, 0xb8, 0x90, 0xc7, 0x66, 0x5e, 0xa8, 0x7e,
	0x40, 0xc8, 0x61, 0x67, 0x67, 0xcc, 0x4e, 0x46, 0x7c, 0xa2, 0xa1, 0x96, 0xd2, 0x5e, 0x86, 0x65,
	0xdf, 0x2e, 0x54, 0x89, 0xa5, 0x49, 0xac, 0xc2, 0x84, 0xcd, 0x12, 0x47, 0xed, 0xa3, 0x9a, 0x9f,
	0x78, 0x16, 0x8d, 0xb4, 0x06, 0x10, 0x77, 0xb3, 0xd4, 0xd8, 0xcc, 0x15, 0x89, 0xb6, 0x2d, 0xd1,
	0x72, 0x03, 0x36, 0x67, 0xb5, 0xea, 0x5b, 0x54, 0x77, 0x49, 0x3c, 0x18, 0x31, 0x8f, 0x71, 0xad,
	0x09, 0xbf, 0xc3, 0x83, 0x4c, 0x5c, 0x67, 0x2e, 0x4a, 0xac, 0x9b, 0x12, 0xeb, 0xda, 0x83, 0xcd,
	0x35, 0x97, 0xc4, 0x6f, 0xc4, 0x51, 0x7d, 0x8d, 0xc4, 0x79, 0x90, 0xc4, 0xd4, 0xd1, 0xfe, 0x03,
	0xde, 0x7d, 0xb1, 0x84, 0x73, 0x4d, 0xc2, 0xed, 0xfc, 0x83, 0x13, 0x16, 0x6c, 0xae, 0xba, 0x24,
	0x3e, 0x8d, 0xa9, 0xa3, 0x3e, 0x43, 0x55, 0xce, 0x3c, 0xaa, 0xad, 0x03, 0xa8, 0x2d, 0x56, 0x4f,
	0xc4, 0x12, 0x44, 0x95, 0xb7, 0x85, 0x79, 0x14, 0x9b, 0x50, 0xa5, 0xbe, 0x47, 0x88, 0x8e, 0x79,
	0x44, 0x06, 0x0e, 0xe1, 0x44, 0xdb, 0x80, 0xe7, 0x83, 0xb9, 0x17, 0xea, 0x82, 0xb9, 0x17, 0x26,
	0x6c, 0xd6, 0x21, 0xe8, 0x13, 0x4e, 0x04, 0xd4, 0x63, 0xe3, 0x81, 0xc3, 0x5c, 0x1a, 0x73, 0x6d,
	0x13, 0x56, 0x0e, 0xa0, 0x85, 0xba, 0x00, 0x5a, 0x98, 0xb0, 0x59, 0xf7, 0xd8, 0xb8, 0x0f, 0x67,
	0xb1, 0x63, 0x7e, 0xe0, 0xdb, 0x54, 0xdb, 0x2a, 0x76, 0x0c, 0x84, 0x05, 0x3b, 0x06, 0x79, 0x6c,
	0xe6, 0x85, 0xbd, 0xd6, 0xc5, 0xa5, 0xbe, 0xf4, 0xeb, 0x52, 0x5f, 0xfa, 0x3a, 0xd5, 0x95, 0x8b,
	0xa9, 0xae, 0xfc, 0x98, 0xea, 0xca, 0xef, 0xa9, 0xae, 0x7c, 0xbb, 0xd2, 0x97, 0xbe, 0x5f, 0xe9,
	0x8a, 0x55, 0x83, 0x7f, 0xe0, 0x87, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x69, 0x03, 0x7f, 0x66,
	0xf5, 0x05, 0x00, 0x00,
}
