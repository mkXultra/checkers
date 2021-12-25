// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/leaderboard.proto

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

type Leaderboard struct {
	Winners []*WinningPlayer `protobuf:"bytes,1,rep,name=winners,proto3" json:"winners,omitempty"`
}

func (m *Leaderboard) Reset()         { *m = Leaderboard{} }
func (m *Leaderboard) String() string { return proto.CompactTextString(m) }
func (*Leaderboard) ProtoMessage()    {}
func (*Leaderboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9d046210f4aa4a, []int{0}
}
func (m *Leaderboard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Leaderboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Leaderboard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Leaderboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Leaderboard.Merge(m, src)
}
func (m *Leaderboard) XXX_Size() int {
	return m.Size()
}
func (m *Leaderboard) XXX_DiscardUnknown() {
	xxx_messageInfo_Leaderboard.DiscardUnknown(m)
}

var xxx_messageInfo_Leaderboard proto.InternalMessageInfo

func (m *Leaderboard) GetWinners() []*WinningPlayer {
	if m != nil {
		return m.Winners
	}
	return nil
}

func init() {
	proto.RegisterType((*Leaderboard)(nil), "mkXultra.checkers.checkers.Leaderboard")
}

func init() { proto.RegisterFile("checkers/leaderboard.proto", fileDescriptor_2f9d046210f4aa4a) }

var fileDescriptor_2f9d046210f4aa4a = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0x48, 0x4d,
	0xce, 0x4e, 0x2d, 0x2a, 0xd6, 0xcf, 0x49, 0x4d, 0x4c, 0x49, 0x2d, 0x4a, 0xca, 0x4f, 0x2c, 0x4a,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xca, 0xcd, 0x8e, 0x28, 0xcd, 0x29, 0x29, 0x4a,
	0xd4, 0x83, 0x29, 0x82, 0x33, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xca, 0xf4, 0x41, 0x2c,
	0x88, 0x0e, 0x29, 0x59, 0xb8, 0x69, 0xe5, 0x99, 0x79, 0x79, 0x99, 0x79, 0xe9, 0xf1, 0x05, 0x39,
	0x89, 0x95, 0xa9, 0x45, 0x10, 0x69, 0xa5, 0x20, 0x2e, 0x6e, 0x1f, 0x84, 0x2d, 0x42, 0xce, 0x5c,
	0xec, 0x20, 0x65, 0xa9, 0x45, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x9a, 0x7a, 0xb8,
	0x6d, 0xd4, 0x0b, 0x87, 0x98, 0x18, 0x00, 0x36, 0x30, 0x08, 0xa6, 0xd3, 0xc9, 0xfd, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0x61, 0xe6, 0xea, 0xc3, 0x1d, 0x58, 0x81, 0x60, 0x96, 0x54, 0x16, 0xa4,
	0x16, 0x27, 0xb1, 0x81, 0xdd, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xc2, 0x5a, 0xdc,
	0x12, 0x01, 0x00, 0x00,
}

func (m *Leaderboard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Leaderboard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Leaderboard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Winners) > 0 {
		for iNdEx := len(m.Winners) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Winners[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLeaderboard(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintLeaderboard(dAtA []byte, offset int, v uint64) int {
	offset -= sovLeaderboard(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Leaderboard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Winners) > 0 {
		for _, e := range m.Winners {
			l = e.Size()
			n += 1 + l + sovLeaderboard(uint64(l))
		}
	}
	return n
}

func sovLeaderboard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLeaderboard(x uint64) (n int) {
	return sovLeaderboard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Leaderboard) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLeaderboard
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
			return fmt.Errorf("proto: Leaderboard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Leaderboard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaderboard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLeaderboard
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLeaderboard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winners = append(m.Winners, &WinningPlayer{})
			if err := m.Winners[len(m.Winners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLeaderboard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLeaderboard
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
func skipLeaderboard(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLeaderboard
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
					return 0, ErrIntOverflowLeaderboard
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
					return 0, ErrIntOverflowLeaderboard
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
				return 0, ErrInvalidLengthLeaderboard
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLeaderboard
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLeaderboard
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLeaderboard        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLeaderboard          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLeaderboard = fmt.Errorf("proto: unexpected end of group")
)
