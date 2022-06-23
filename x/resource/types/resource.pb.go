// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource/v1/resource.proto

package types

import (
	fmt "fmt"
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

type Resource struct {
	Header *ResourceHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []byte          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_cebae6241f1ea243, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetHeader() *ResourceHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Resource) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ResourceHeader struct {
	CollectionId      string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Id                string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ResourceType      string `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	MimeType          string `protobuf:"bytes,5,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Created           string `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Checksum          []byte `protobuf:"bytes,7,opt,name=checksum,proto3" json:"checksum,omitempty"`
	PreviousVersionId string `protobuf:"bytes,8,opt,name=previous_version_id,json=previousVersionId,proto3" json:"previous_version_id,omitempty"`
	NextVersionId     string `protobuf:"bytes,9,opt,name=next_version_id,json=nextVersionId,proto3" json:"next_version_id,omitempty"`
}

func (m *ResourceHeader) Reset()         { *m = ResourceHeader{} }
func (m *ResourceHeader) String() string { return proto.CompactTextString(m) }
func (*ResourceHeader) ProtoMessage()    {}
func (*ResourceHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_cebae6241f1ea243, []int{1}
}
func (m *ResourceHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHeader.Merge(m, src)
}
func (m *ResourceHeader) XXX_Size() int {
	return m.Size()
}
func (m *ResourceHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHeader proto.InternalMessageInfo

func (m *ResourceHeader) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *ResourceHeader) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourceHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHeader) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ResourceHeader) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *ResourceHeader) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *ResourceHeader) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func (m *ResourceHeader) GetPreviousVersionId() string {
	if m != nil {
		return m.PreviousVersionId
	}
	return ""
}

func (m *ResourceHeader) GetNextVersionId() string {
	if m != nil {
		return m.NextVersionId
	}
	return ""
}

func init() {
	proto.RegisterType((*Resource)(nil), "cheqdid.cheqdnode.resource.v1.Resource")
	proto.RegisterType((*ResourceHeader)(nil), "cheqdid.cheqdnode.resource.v1.ResourceHeader")
}

func init() { proto.RegisterFile("resource/v1/resource.proto", fileDescriptor_cebae6241f1ea243) }

var fileDescriptor_cebae6241f1ea243 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x4e, 0xea, 0x40,
	0x14, 0xa6, 0xbd, 0x5c, 0x68, 0xe7, 0x02, 0x37, 0x77, 0xee, 0xa6, 0xc1, 0xd8, 0x10, 0x4c, 0x0c,
	0x0b, 0x69, 0x83, 0xbe, 0x81, 0x89, 0x46, 0xb6, 0x8d, 0x71, 0xe1, 0x86, 0x94, 0x99, 0x13, 0x3b,
	0x91, 0x76, 0xea, 0x74, 0xda, 0xc0, 0x5b, 0xf8, 0x1a, 0xbe, 0x89, 0x4b, 0x96, 0x2e, 0x0d, 0xbc,
	0x88, 0xe9, 0xa1, 0xad, 0xba, 0x71, 0xd3, 0x9e, 0xf3, 0xfd, 0xcd, 0x9c, 0x9c, 0x21, 0x43, 0x05,
	0x99, 0xcc, 0x15, 0x03, 0xbf, 0x98, 0xf9, 0x75, 0xed, 0xa5, 0x4a, 0x6a, 0x49, 0x8f, 0x59, 0x04,
	0x4f, 0x5c, 0x70, 0x0f, 0xff, 0x89, 0xe4, 0xe0, 0x35, 0x8a, 0x62, 0x36, 0x06, 0x62, 0x05, 0x55,
	0x4b, 0xaf, 0x48, 0x27, 0x82, 0x90, 0x83, 0x72, 0x8c, 0x91, 0x31, 0xf9, 0x73, 0x3e, 0xf5, 0x7e,
	0xf4, 0x7a, 0xb5, 0xf1, 0x06, 0x4d, 0x41, 0x65, 0xa6, 0x94, 0xb4, 0x79, 0xa8, 0x43, 0xc7, 0x1c,
	0x19, 0x93, 0x5e, 0x80, 0xf5, 0xf8, 0xc5, 0x24, 0x83, 0xef, 0x72, 0x7a, 0x42, 0xfa, 0x4c, 0xae,
	0x56, 0xc0, 0xb4, 0x90, 0xc9, 0x42, 0x70, 0x3c, 0xd4, 0x0e, 0x7a, 0x9f, 0xe0, 0x9c, 0xd3, 0x01,
	0x31, 0x05, 0xc7, 0x24, 0x3b, 0x30, 0x05, 0x2f, 0xb3, 0x93, 0x30, 0x06, 0xe7, 0x17, 0x22, 0x58,
	0x97, 0x41, 0xf5, 0xad, 0x16, 0x7a, 0x93, 0x82, 0xd3, 0x3e, 0x04, 0xd5, 0xe0, 0xed, 0x26, 0x05,
	0x7a, 0x44, 0xec, 0x58, 0xc4, 0x95, 0xe0, 0x37, 0x0a, 0xac, 0x12, 0x40, 0xd2, 0x21, 0x5d, 0xa6,
	0x20, 0xd4, 0xc0, 0x9d, 0x0e, 0x52, 0x75, 0x4b, 0x87, 0xc4, 0x62, 0x11, 0xb0, 0xc7, 0x2c, 0x8f,
	0x9d, 0x2e, 0xce, 0xd3, 0xf4, 0xd4, 0x23, 0xff, 0x53, 0x05, 0x85, 0x90, 0x79, 0xb6, 0x28, 0x40,
	0x65, 0xd5, 0x18, 0x16, 0x26, 0xfc, 0xab, 0xa9, 0xbb, 0x03, 0x33, 0xe7, 0xf4, 0x94, 0xfc, 0x4d,
	0x60, 0xad, 0xbf, 0x6a, 0x6d, 0xd4, 0xf6, 0x4b, 0xb8, 0xd1, 0x5d, 0x5e, 0xbf, 0xee, 0x5c, 0x63,
	0xbb, 0x73, 0x8d, 0xf7, 0x9d, 0x6b, 0x3c, 0xef, 0xdd, 0xd6, 0x76, 0xef, 0xb6, 0xde, 0xf6, 0x6e,
	0xeb, 0xfe, 0xec, 0x41, 0xe8, 0x28, 0x5f, 0x7a, 0x4c, 0xc6, 0x3e, 0xae, 0xe4, 0xf0, 0x9d, 0x96,
	0x9b, 0xf1, 0xd7, 0xcd, 0xe6, 0xfd, 0x72, 0xc8, 0x6c, 0xd9, 0xc1, 0x07, 0x70, 0xf1, 0x11, 0x00,
	0x00, 0xff, 0xff, 0x96, 0xa7, 0xa9, 0x15, 0x1e, 0x02, 0x00, 0x00,
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResourceHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextVersionId) > 0 {
		i -= len(m.NextVersionId)
		copy(dAtA[i:], m.NextVersionId)
		i = encodeVarintResource(dAtA, i, uint64(len(m.NextVersionId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.PreviousVersionId) > 0 {
		i -= len(m.PreviousVersionId)
		copy(dAtA[i:], m.PreviousVersionId)
		i = encodeVarintResource(dAtA, i, uint64(len(m.PreviousVersionId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Created) > 0 {
		i -= len(m.Created)
		copy(dAtA[i:], m.Created)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Created)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MimeType) > 0 {
		i -= len(m.MimeType)
		copy(dAtA[i:], m.MimeType)
		i = encodeVarintResource(dAtA, i, uint64(len(m.MimeType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ResourceType) > 0 {
		i -= len(m.ResourceType)
		copy(dAtA[i:], m.ResourceType)
		i = encodeVarintResource(dAtA, i, uint64(len(m.ResourceType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CollectionId) > 0 {
		i -= len(m.CollectionId)
		copy(dAtA[i:], m.CollectionId)
		i = encodeVarintResource(dAtA, i, uint64(len(m.CollectionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func (m *ResourceHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollectionId)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.ResourceType)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.MimeType)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.PreviousVersionId)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.NextVersionId)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResourceHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
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
func (m *ResourceHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: ResourceHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MimeType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MimeType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Created = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousVersionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousVersionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextVersionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextVersionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)
