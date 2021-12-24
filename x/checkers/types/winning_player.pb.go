// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/winning_player.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type WinningPlayer struct {
	PlayerAddress string `protobuf:"bytes,1,opt,name=playerAddress,proto3" json:"playerAddress,omitempty"`
	WonCount      uint64 `protobuf:"varint,2,opt,name=wonCount,proto3" json:"wonCount,omitempty"`
	DateAdded     string `protobuf:"bytes,3,opt,name=dateAdded,proto3" json:"dateAdded,omitempty"`
}

func (m *WinningPlayer) Reset()         { *m = WinningPlayer{} }
func (m *WinningPlayer) String() string { return proto.CompactTextString(m) }
func (*WinningPlayer) ProtoMessage()    {}
func (*WinningPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f70b93a04992d2f, []int{0}
}
func (m *WinningPlayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WinningPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WinningPlayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WinningPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WinningPlayer.Merge(m, src)
}
func (m *WinningPlayer) XXX_Size() int {
	return m.Size()
}
func (m *WinningPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_WinningPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_WinningPlayer proto.InternalMessageInfo

func (m *WinningPlayer) GetPlayerAddress() string {
	if m != nil {
		return m.PlayerAddress
	}
	return ""
}

func (m *WinningPlayer) GetWonCount() uint64 {
	if m != nil {
		return m.WonCount
	}
	return 0
}

func (m *WinningPlayer) GetDateAdded() string {
	if m != nil {
		return m.DateAdded
	}
	return ""
}

func init() {
	proto.RegisterType((*WinningPlayer)(nil), "mkXultra.checkers.checkers.WinningPlayer")
}

func init() { proto.RegisterFile("checkers/winning_player.proto", fileDescriptor_9f70b93a04992d2f) }

var fileDescriptor_9f70b93a04992d2f = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0x48, 0x4d,
	0xce, 0x4e, 0x2d, 0x2a, 0xd6, 0x2f, 0xcf, 0xcc, 0xcb, 0xcb, 0xcc, 0x4b, 0x8f, 0x2f, 0xc8, 0x49,
	0xac, 0x4c, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xca, 0xcd, 0x8e, 0x28, 0xcd,
	0x29, 0x29, 0x4a, 0xd4, 0x83, 0xa9, 0x83, 0x33, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xca,
	0xf4, 0x41, 0x2c, 0x88, 0x0e, 0xa5, 0x7c, 0x2e, 0xde, 0x70, 0x88, 0x49, 0x01, 0x60, 0x83, 0x84,
	0x54, 0xb8, 0x78, 0x21, 0x46, 0x3a, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x0a, 0x49, 0x71, 0x71, 0x94, 0xe7, 0xe7, 0x39, 0xe7, 0x97, 0xe6,
	0x95, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0xc1, 0xf9, 0x42, 0x32, 0x5c, 0x9c, 0x29, 0x89,
	0x25, 0xa9, 0x8e, 0x29, 0x29, 0xa9, 0x29, 0x12, 0xcc, 0x60, 0xdd, 0x08, 0x01, 0x27, 0xf7, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xf9, 0x43, 0x1f, 0xee, 0xdf, 0x0a, 0x04, 0xb3, 0xa4, 0xb2,
	0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x01, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x2a,
	0x78, 0x25, 0x13, 0x01, 0x00, 0x00,
}

func (m *WinningPlayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WinningPlayer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WinningPlayer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DateAdded) > 0 {
		i -= len(m.DateAdded)
		copy(dAtA[i:], m.DateAdded)
		i = encodeVarintWinningPlayer(dAtA, i, uint64(len(m.DateAdded)))
		i--
		dAtA[i] = 0x1a
	}
	if m.WonCount != 0 {
		i = encodeVarintWinningPlayer(dAtA, i, uint64(m.WonCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PlayerAddress) > 0 {
		i -= len(m.PlayerAddress)
		copy(dAtA[i:], m.PlayerAddress)
		i = encodeVarintWinningPlayer(dAtA, i, uint64(len(m.PlayerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWinningPlayer(dAtA []byte, offset int, v uint64) int {
	offset -= sovWinningPlayer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WinningPlayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PlayerAddress)
	if l > 0 {
		n += 1 + l + sovWinningPlayer(uint64(l))
	}
	if m.WonCount != 0 {
		n += 1 + sovWinningPlayer(uint64(m.WonCount))
	}
	l = len(m.DateAdded)
	if l > 0 {
		n += 1 + l + sovWinningPlayer(uint64(l))
	}
	return n
}

func sovWinningPlayer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWinningPlayer(x uint64) (n int) {
	return sovWinningPlayer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WinningPlayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWinningPlayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WinningPlayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WinningPlayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWinningPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWinningPlayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWinningPlayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WonCount", wireType)
			}
			m.WonCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWinningPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WonCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateAdded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWinningPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWinningPlayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWinningPlayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DateAdded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWinningPlayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWinningPlayer
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
func skipWinningPlayer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWinningPlayer
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
					return 0, ErrIntOverflowWinningPlayer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWinningPlayer
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
			if length < 0 {
				return 0, ErrInvalidLengthWinningPlayer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWinningPlayer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWinningPlayer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWinningPlayer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWinningPlayer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWinningPlayer = fmt.Errorf("proto: unexpected end of group")
)
