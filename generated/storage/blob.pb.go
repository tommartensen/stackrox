// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/blob.proto

package storage

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Next Tag: 7
type Blob struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" sql:"pk" search:"Blob Name"` // @gotags: sql:"pk" search:"Blob Name"
	Oid                  uint32           `protobuf:"varint,2,opt,name=oid,proto3" json:"oid,omitempty"`
	Checksum             string           `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Length               int64            `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty" search:"Blob Length"` // @gotags: search:"Blob Length"
	LastUpdated          *types.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	ModifiedTime         *types.Timestamp `protobuf:"bytes,6,opt,name=modified_time,json=modifiedTime,proto3" json:"modified_time,omitempty" search:"Blob Modified On"` // @gotags: search:"Blob Modified On"
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Blob) Reset()         { *m = Blob{} }
func (m *Blob) String() string { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()    {}
func (*Blob) Descriptor() ([]byte, []int) {
	return fileDescriptor_93b63e008eb8666f, []int{0}
}
func (m *Blob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Blob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Blob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Blob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blob.Merge(m, src)
}
func (m *Blob) XXX_Size() int {
	return m.Size()
}
func (m *Blob) XXX_DiscardUnknown() {
	xxx_messageInfo_Blob.DiscardUnknown(m)
}

var xxx_messageInfo_Blob proto.InternalMessageInfo

func (m *Blob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Blob) GetOid() uint32 {
	if m != nil {
		return m.Oid
	}
	return 0
}

func (m *Blob) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *Blob) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Blob) GetLastUpdated() *types.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func (m *Blob) GetModifiedTime() *types.Timestamp {
	if m != nil {
		return m.ModifiedTime
	}
	return nil
}

func (m *Blob) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Blob) Clone() *Blob {
	if m == nil {
		return nil
	}
	cloned := new(Blob)
	*cloned = *m

	cloned.LastUpdated = m.LastUpdated.Clone()
	cloned.ModifiedTime = m.ModifiedTime.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*Blob)(nil), "storage.Blob")
}

func init() { proto.RegisterFile("storage/blob.proto", fileDescriptor_93b63e008eb8666f) }

var fileDescriptor_93b63e008eb8666f = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x8d, 0xad, 0x55, 0x33, 0x33, 0x20, 0x59, 0x48, 0xec, 0xa2, 0x16, 0x57, 0x5d, 0xa5,
	0xa0, 0x6e, 0x45, 0x98, 0x47, 0x28, 0xba, 0x71, 0x33, 0x24, 0x4d, 0xa6, 0x13, 0x26, 0x99, 0x5b,
	0x9a, 0x14, 0x7c, 0x14, 0x1f, 0xc9, 0xa5, 0x8f, 0x30, 0xd4, 0x17, 0x91, 0xfe, 0xb9, 0x9d, 0xdd,
	0xbd, 0x87, 0xef, 0xc0, 0xc7, 0xc1, 0xc4, 0x79, 0x68, 0x78, 0xa5, 0x72, 0x61, 0x40, 0xb0, 0xba,
	0x01, 0x0f, 0xe4, 0x72, 0xca, 0xe2, 0xfb, 0x0a, 0xa0, 0x32, 0x2a, 0x1f, 0x62, 0xd1, 0x6e, 0x73,
	0xaf, 0xad, 0x72, 0x9e, 0xdb, 0x7a, 0x24, 0x1f, 0x8e, 0x08, 0x87, 0x6b, 0x03, 0x82, 0x10, 0x1c,
	0x1e, 0xb8, 0x55, 0x14, 0xa5, 0x28, 0xbb, 0x2e, 0x86, 0x9b, 0xdc, 0xe0, 0x00, 0xb4, 0xa4, 0xe7,
	0x29, 0xca, 0x56, 0x45, 0x7f, 0x92, 0x18, 0x5f, 0x95, 0x3b, 0x55, 0xee, 0x5d, 0x6b, 0x69, 0x30,
	0x90, 0xff, 0x3f, 0xb9, 0xc5, 0x91, 0x51, 0x87, 0xca, 0xef, 0x68, 0x98, 0xa2, 0x2c, 0x28, 0xa6,
	0x8f, 0xbc, 0xe0, 0xa5, 0xe1, 0xce, 0x6f, 0xda, 0x5a, 0x72, 0xaf, 0x24, 0xbd, 0x48, 0x51, 0xb6,
	0x78, 0x8c, 0xd9, 0xa8, 0xc6, 0x66, 0x35, 0xf6, 0x36, 0xab, 0x15, 0x8b, 0x9e, 0x7f, 0x1f, 0x71,
	0xf2, 0x8a, 0x57, 0x16, 0xa4, 0xde, 0x6a, 0x25, 0x37, 0xbd, 0x3d, 0x8d, 0x4e, 0xf6, 0x97, 0x73,
	0xa1, 0x8f, 0xd6, 0xcf, 0xdf, 0x5d, 0x82, 0x7e, 0xba, 0x04, 0x1d, 0xbb, 0x04, 0x7d, 0xfd, 0x26,
	0x67, 0xf8, 0x4e, 0x03, 0x73, 0x9e, 0x97, 0xfb, 0x06, 0x3e, 0xc7, 0x3e, 0x9b, 0x06, 0xfb, 0x98,
	0x97, 0x13, 0xd1, 0x90, 0x3f, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x8b, 0x60, 0x60, 0x5f,
	0x01, 0x00, 0x00,
}

func (m *Blob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Blob) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Blob) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ModifiedTime != nil {
		{
			size, err := m.ModifiedTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlob(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.LastUpdated != nil {
		{
			size, err := m.LastUpdated.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlob(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Length != 0 {
		i = encodeVarintBlob(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintBlob(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Oid != 0 {
		i = encodeVarintBlob(dAtA, i, uint64(m.Oid))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintBlob(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlob(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlob(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Blob) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBlob(uint64(l))
	}
	if m.Oid != 0 {
		n += 1 + sovBlob(uint64(m.Oid))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovBlob(uint64(l))
	}
	if m.Length != 0 {
		n += 1 + sovBlob(uint64(m.Length))
	}
	if m.LastUpdated != nil {
		l = m.LastUpdated.Size()
		n += 1 + l + sovBlob(uint64(l))
	}
	if m.ModifiedTime != nil {
		l = m.ModifiedTime.Size()
		n += 1 + l + sovBlob(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBlob(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlob(x uint64) (n int) {
	return sovBlob(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Blob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlob
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
			return fmt.Errorf("proto: Blob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Blob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
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
				return ErrInvalidLengthBlob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oid", wireType)
			}
			m.Oid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Oid |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
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
				return ErrInvalidLengthBlob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
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
				return ErrInvalidLengthBlob
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastUpdated == nil {
				m.LastUpdated = &types.Timestamp{}
			}
			if err := m.LastUpdated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModifiedTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
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
				return ErrInvalidLengthBlob
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ModifiedTime == nil {
				m.ModifiedTime = &types.Timestamp{}
			}
			if err := m.ModifiedTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlob
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlob(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlob
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
					return 0, ErrIntOverflowBlob
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
					return 0, ErrIntOverflowBlob
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
				return 0, ErrInvalidLengthBlob
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlob
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlob
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlob        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlob          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlob = fmt.Errorf("proto: unexpected end of group")
)
