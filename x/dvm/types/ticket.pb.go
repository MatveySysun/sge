// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/dvm/ticket.proto

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

// PubkeysChangeProposalPayload indicates data of public keys changes proposal
// ticket.
type PubkeysChangeProposalPayload struct {
	// additions contain new pub keys to be added to public keys.
	Additions []string `protobuf:"bytes,1,rep,name=additions,proto3" json:"additions,omitempty"`
	// deletions contain old pub keys to be removed from public keys.
	Deletions []string `protobuf:"bytes,2,rep,name=deletions,proto3" json:"deletions,omitempty"`
}

func (m *PubkeysChangeProposalPayload) Reset()         { *m = PubkeysChangeProposalPayload{} }
func (m *PubkeysChangeProposalPayload) String() string { return proto.CompactTextString(m) }
func (*PubkeysChangeProposalPayload) ProtoMessage()    {}
func (*PubkeysChangeProposalPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7121f0f264448d06, []int{0}
}
func (m *PubkeysChangeProposalPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubkeysChangeProposalPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubkeysChangeProposalPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubkeysChangeProposalPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubkeysChangeProposalPayload.Merge(m, src)
}
func (m *PubkeysChangeProposalPayload) XXX_Size() int {
	return m.Size()
}
func (m *PubkeysChangeProposalPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PubkeysChangeProposalPayload.DiscardUnknown(m)
}

var xxx_messageInfo_PubkeysChangeProposalPayload proto.InternalMessageInfo

func (m *PubkeysChangeProposalPayload) GetAdditions() []string {
	if m != nil {
		return m.Additions
	}
	return nil
}

func (m *PubkeysChangeProposalPayload) GetDeletions() []string {
	if m != nil {
		return m.Deletions
	}
	return nil
}

// ProposalVotePayload indicates vote data ticket.
type ProposalVotePayload struct {
	// proposal_id is the id of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// vote is the chosen option for the vote.
	Vote ProposalVote `protobuf:"varint,2,opt,name=vote,proto3,enum=sgenetwork.sge.dvm.ProposalVote" json:"vote,omitempty"`
}

func (m *ProposalVotePayload) Reset()         { *m = ProposalVotePayload{} }
func (m *ProposalVotePayload) String() string { return proto.CompactTextString(m) }
func (*ProposalVotePayload) ProtoMessage()    {}
func (*ProposalVotePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7121f0f264448d06, []int{1}
}
func (m *ProposalVotePayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalVotePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalVotePayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalVotePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalVotePayload.Merge(m, src)
}
func (m *ProposalVotePayload) XXX_Size() int {
	return m.Size()
}
func (m *ProposalVotePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalVotePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalVotePayload proto.InternalMessageInfo

func (m *ProposalVotePayload) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *ProposalVotePayload) GetVote() ProposalVote {
	if m != nil {
		return m.Vote
	}
	return ProposalVote_PROPOSAL_VOTE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*PubkeysChangeProposalPayload)(nil), "sgenetwork.sge.dvm.PubkeysChangeProposalPayload")
	proto.RegisterType((*ProposalVotePayload)(nil), "sgenetwork.sge.dvm.ProposalVotePayload")
}

func init() { proto.RegisterFile("sge/dvm/ticket.proto", fileDescriptor_7121f0f264448d06) }

var fileDescriptor_7121f0f264448d06 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4e, 0x4f, 0xd5,
	0x4f, 0x29, 0xcb, 0xd5, 0x2f, 0xc9, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x2a, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x4e, 0x4f,
	0xd5, 0x4b, 0x29, 0xcb, 0x95, 0x12, 0x82, 0xa9, 0x2c, 0xcb, 0x2f, 0x49, 0x85, 0xa8, 0x53, 0x8a,
	0xe2, 0x92, 0x09, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0x2c, 0x76, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x0d,
	0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0x09, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x11, 0x92,
	0xe1, 0xe2, 0x4c, 0x4c, 0x49, 0xc9, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x96, 0x60, 0x54, 0x60, 0xd6,
	0xe0, 0x0c, 0x42, 0x08, 0x80, 0x64, 0x53, 0x52, 0x73, 0x52, 0x21, 0xb2, 0x4c, 0x10, 0x59, 0xb8,
	0x80, 0x52, 0x0e, 0x97, 0x30, 0xcc, 0xb8, 0xb0, 0xfc, 0x92, 0x54, 0x98, 0x91, 0xf2, 0x5c, 0xdc,
	0x05, 0x50, 0xe1, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x2e, 0x98, 0x90,
	0x67, 0x8a, 0x90, 0x09, 0x17, 0x0b, 0xc8, 0x85, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x0a,
	0x7a, 0x98, 0x5e, 0xd1, 0x43, 0x36, 0x37, 0x08, 0xac, 0xda, 0xc9, 0xe1, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xd4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x8b, 0xd3, 0x53, 0x75, 0xa1, 0x86, 0x81, 0xd8, 0xfa, 0x15, 0x90, 0xa0, 0xab, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0x07, 0x89, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x8d, 0x8d,
	0x94, 0x52, 0x01, 0x00, 0x00,
}

func (m *PubkeysChangeProposalPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubkeysChangeProposalPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubkeysChangeProposalPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deletions) > 0 {
		for iNdEx := len(m.Deletions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Deletions[iNdEx])
			copy(dAtA[i:], m.Deletions[iNdEx])
			i = encodeVarintTicket(dAtA, i, uint64(len(m.Deletions[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Additions) > 0 {
		for iNdEx := len(m.Additions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Additions[iNdEx])
			copy(dAtA[i:], m.Additions[iNdEx])
			i = encodeVarintTicket(dAtA, i, uint64(len(m.Additions[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProposalVotePayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalVotePayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalVotePayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x10
	}
	if m.ProposalId != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PubkeysChangeProposalPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Additions) > 0 {
		for _, s := range m.Additions {
			l = len(s)
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if len(m.Deletions) > 0 {
		for _, s := range m.Deletions {
			l = len(s)
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	return n
}

func (m *ProposalVotePayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovTicket(uint64(m.ProposalId))
	}
	if m.Vote != 0 {
		n += 1 + sovTicket(uint64(m.Vote))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PubkeysChangeProposalPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: PubkeysChangeProposalPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubkeysChangeProposalPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Additions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Additions = append(m.Additions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deletions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deletions = append(m.Deletions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *ProposalVotePayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: ProposalVotePayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalVotePayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= ProposalVote(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
