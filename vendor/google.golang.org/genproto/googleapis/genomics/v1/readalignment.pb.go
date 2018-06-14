// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/genomics/v1/readalignment.proto

package genomics // import "google.golang.org/genproto/googleapis/genomics/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A linear alignment can be represented by one CIGAR string. Describes the
// mapped position and local alignment of the read to the reference.
type LinearAlignment struct {
	// The position of this alignment.
	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// The mapping quality of this alignment. Represents how likely
	// the read maps to this position as opposed to other locations.
	//
	// Specifically, this is -10 log10 Pr(mapping position is wrong), rounded to
	// the nearest integer.
	MappingQuality int32 `protobuf:"varint,2,opt,name=mapping_quality,json=mappingQuality,proto3" json:"mapping_quality,omitempty"`
	// Represents the local alignment of this sequence (alignment matches, indels,
	// etc) against the reference.
	Cigar                []*CigarUnit `protobuf:"bytes,3,rep,name=cigar,proto3" json:"cigar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LinearAlignment) Reset()         { *m = LinearAlignment{} }
func (m *LinearAlignment) String() string { return proto.CompactTextString(m) }
func (*LinearAlignment) ProtoMessage()    {}
func (*LinearAlignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_readalignment_ff092b0bea51b5ff, []int{0}
}
func (m *LinearAlignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinearAlignment.Unmarshal(m, b)
}
func (m *LinearAlignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinearAlignment.Marshal(b, m, deterministic)
}
func (dst *LinearAlignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinearAlignment.Merge(dst, src)
}
func (m *LinearAlignment) XXX_Size() int {
	return xxx_messageInfo_LinearAlignment.Size(m)
}
func (m *LinearAlignment) XXX_DiscardUnknown() {
	xxx_messageInfo_LinearAlignment.DiscardUnknown(m)
}

var xxx_messageInfo_LinearAlignment proto.InternalMessageInfo

func (m *LinearAlignment) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *LinearAlignment) GetMappingQuality() int32 {
	if m != nil {
		return m.MappingQuality
	}
	return 0
}

func (m *LinearAlignment) GetCigar() []*CigarUnit {
	if m != nil {
		return m.Cigar
	}
	return nil
}

// A read alignment describes a linear alignment of a string of DNA to a
// [reference sequence][google.genomics.v1.Reference], in addition to metadata
// about the fragment (the molecule of DNA sequenced) and the read (the bases
// which were read by the sequencer). A read is equivalent to a line in a SAM
// file. A read belongs to exactly one read group and exactly one
// [read group set][google.genomics.v1.ReadGroupSet].
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
//
// ### Reverse-stranded reads
//
// Mapped reads (reads having a non-null `alignment`) can be aligned to either
// the forward or the reverse strand of their associated reference. Strandedness
// of a mapped read is encoded by `alignment.position.reverseStrand`.
//
// If we consider the reference to be a forward-stranded coordinate space of
// `[0, reference.length)` with `0` as the left-most position and
// `reference.length` as the right-most position, reads are always aligned left
// to right. That is, `alignment.position.position` always refers to the
// left-most reference coordinate and `alignment.cigar` describes the alignment
// of this read to the reference from left to right. All per-base fields such as
// `alignedSequence` and `alignedQuality` share this same left-to-right
// orientation; this is true of reads which are aligned to either strand. For
// reverse-stranded reads, this means that `alignedSequence` is the reverse
// complement of the bases that were originally reported by the sequencing
// machine.
//
// ### Generating a reference-aligned sequence string
//
// When interacting with mapped reads, it's often useful to produce a string
// representing the local alignment of the read to reference. The following
// pseudocode demonstrates one way of doing this:
//
//     out = ""
//     offset = 0
//     for c in read.alignment.cigar {
//       switch c.operation {
//       case "ALIGNMENT_MATCH", "SEQUENCE_MATCH", "SEQUENCE_MISMATCH":
//         out += read.alignedSequence[offset:offset+c.operationLength]
//         offset += c.operationLength
//         break
//       case "CLIP_SOFT", "INSERT":
//         offset += c.operationLength
//         break
//       case "PAD":
//         out += repeat("*", c.operationLength)
//         break
//       case "DELETE":
//         out += repeat("-", c.operationLength)
//         break
//       case "SKIP":
//         out += repeat(" ", c.operationLength)
//         break
//       case "CLIP_HARD":
//         break
//       }
//     }
//     return out
//
// ### Converting to SAM's CIGAR string
//
// The following pseudocode generates a SAM CIGAR string from the
// `cigar` field. Note that this is a lossy conversion
// (`cigar.referenceSequence` is lost).
//
//     cigarMap = {
//       "ALIGNMENT_MATCH": "M",
//       "INSERT": "I",
//       "DELETE": "D",
//       "SKIP": "N",
//       "CLIP_SOFT": "S",
//       "CLIP_HARD": "H",
//       "PAD": "P",
//       "SEQUENCE_MATCH": "=",
//       "SEQUENCE_MISMATCH": "X",
//     }
//     cigarStr = ""
//     for c in read.alignment.cigar {
//       cigarStr += c.operationLength + cigarMap[c.operation]
//     }
//     return cigarStr
type Read struct {
	// The server-generated read ID, unique across all reads. This is different
	// from the `fragmentName`.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the read group this read belongs to. A read belongs to exactly
	// one read group. This is a server-generated ID which is distinct from SAM's
	// RG tag (for that value, see
	// [ReadGroup.name][google.genomics.v1.ReadGroup.name]).
	ReadGroupId string `protobuf:"bytes,2,opt,name=read_group_id,json=readGroupId,proto3" json:"read_group_id,omitempty"`
	// The ID of the read group set this read belongs to. A read belongs to
	// exactly one read group set.
	ReadGroupSetId string `protobuf:"bytes,3,opt,name=read_group_set_id,json=readGroupSetId,proto3" json:"read_group_set_id,omitempty"`
	// The fragment name. Equivalent to QNAME (query template name) in SAM.
	FragmentName string `protobuf:"bytes,4,opt,name=fragment_name,json=fragmentName,proto3" json:"fragment_name,omitempty"`
	// The orientation and the distance between reads from the fragment are
	// consistent with the sequencing protocol (SAM flag 0x2).
	ProperPlacement bool `protobuf:"varint,5,opt,name=proper_placement,json=properPlacement,proto3" json:"proper_placement,omitempty"`
	// The fragment is a PCR or optical duplicate (SAM flag 0x400).
	DuplicateFragment bool `protobuf:"varint,6,opt,name=duplicate_fragment,json=duplicateFragment,proto3" json:"duplicate_fragment,omitempty"`
	// The observed length of the fragment, equivalent to TLEN in SAM.
	FragmentLength int32 `protobuf:"varint,7,opt,name=fragment_length,json=fragmentLength,proto3" json:"fragment_length,omitempty"`
	// The read number in sequencing. 0-based and less than numberReads. This
	// field replaces SAM flag 0x40 and 0x80.
	ReadNumber int32 `protobuf:"varint,8,opt,name=read_number,json=readNumber,proto3" json:"read_number,omitempty"`
	// The number of reads in the fragment (extension to SAM flag 0x1).
	NumberReads int32 `protobuf:"varint,9,opt,name=number_reads,json=numberReads,proto3" json:"number_reads,omitempty"`
	// Whether this read did not pass filters, such as platform or vendor quality
	// controls (SAM flag 0x200).
	FailedVendorQualityChecks bool `protobuf:"varint,10,opt,name=failed_vendor_quality_checks,json=failedVendorQualityChecks,proto3" json:"failed_vendor_quality_checks,omitempty"`
	// The linear alignment for this alignment record. This field is null for
	// unmapped reads.
	Alignment *LinearAlignment `protobuf:"bytes,11,opt,name=alignment,proto3" json:"alignment,omitempty"`
	// Whether this alignment is secondary. Equivalent to SAM flag 0x100.
	// A secondary alignment represents an alternative to the primary alignment
	// for this read. Aligners may return secondary alignments if a read can map
	// ambiguously to multiple coordinates in the genome. By convention, each read
	// has one and only one alignment where both `secondaryAlignment`
	// and `supplementaryAlignment` are false.
	SecondaryAlignment bool `protobuf:"varint,12,opt,name=secondary_alignment,json=secondaryAlignment,proto3" json:"secondary_alignment,omitempty"`
	// Whether this alignment is supplementary. Equivalent to SAM flag 0x800.
	// Supplementary alignments are used in the representation of a chimeric
	// alignment. In a chimeric alignment, a read is split into multiple
	// linear alignments that map to different reference contigs. The first
	// linear alignment in the read will be designated as the representative
	// alignment; the remaining linear alignments will be designated as
	// supplementary alignments. These alignments may have different mapping
	// quality scores. In each linear alignment in a chimeric alignment, the read
	// will be hard clipped. The `alignedSequence` and
	// `alignedQuality` fields in the alignment record will only
	// represent the bases for its respective linear alignment.
	SupplementaryAlignment bool `protobuf:"varint,13,opt,name=supplementary_alignment,json=supplementaryAlignment,proto3" json:"supplementary_alignment,omitempty"`
	// The bases of the read sequence contained in this alignment record,
	// **without CIGAR operations applied** (equivalent to SEQ in SAM).
	// `alignedSequence` and `alignedQuality` may be
	// shorter than the full read sequence and quality. This will occur if the
	// alignment is part of a chimeric alignment, or if the read was trimmed. When
	// this occurs, the CIGAR for this read will begin/end with a hard clip
	// operator that will indicate the length of the excised sequence.
	AlignedSequence string `protobuf:"bytes,14,opt,name=aligned_sequence,json=alignedSequence,proto3" json:"aligned_sequence,omitempty"`
	// The quality of the read sequence contained in this alignment record
	// (equivalent to QUAL in SAM).
	// `alignedSequence` and `alignedQuality` may be shorter than the full read
	// sequence and quality. This will occur if the alignment is part of a
	// chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR
	// for this read will begin/end with a hard clip operator that will indicate
	// the length of the excised sequence.
	AlignedQuality []int32 `protobuf:"varint,15,rep,packed,name=aligned_quality,json=alignedQuality,proto3" json:"aligned_quality,omitempty"`
	// The mapping of the primary alignment of the
	// `(readNumber+1)%numberReads` read in the fragment. It replaces
	// mate position and mate strand in SAM.
	NextMatePosition *Position `protobuf:"bytes,16,opt,name=next_mate_position,json=nextMatePosition,proto3" json:"next_mate_position,omitempty"`
	// A map of additional read alignment information. This must be of the form
	// map<string, string[]> (string key mapping to a list of string values).
	Info                 map[string]*_struct.ListValue `protobuf:"bytes,17,rep,name=info,proto3" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Read) Reset()         { *m = Read{} }
func (m *Read) String() string { return proto.CompactTextString(m) }
func (*Read) ProtoMessage()    {}
func (*Read) Descriptor() ([]byte, []int) {
	return fileDescriptor_readalignment_ff092b0bea51b5ff, []int{1}
}
func (m *Read) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Read.Unmarshal(m, b)
}
func (m *Read) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Read.Marshal(b, m, deterministic)
}
func (dst *Read) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Read.Merge(dst, src)
}
func (m *Read) XXX_Size() int {
	return xxx_messageInfo_Read.Size(m)
}
func (m *Read) XXX_DiscardUnknown() {
	xxx_messageInfo_Read.DiscardUnknown(m)
}

var xxx_messageInfo_Read proto.InternalMessageInfo

func (m *Read) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Read) GetReadGroupId() string {
	if m != nil {
		return m.ReadGroupId
	}
	return ""
}

func (m *Read) GetReadGroupSetId() string {
	if m != nil {
		return m.ReadGroupSetId
	}
	return ""
}

func (m *Read) GetFragmentName() string {
	if m != nil {
		return m.FragmentName
	}
	return ""
}

func (m *Read) GetProperPlacement() bool {
	if m != nil {
		return m.ProperPlacement
	}
	return false
}

func (m *Read) GetDuplicateFragment() bool {
	if m != nil {
		return m.DuplicateFragment
	}
	return false
}

func (m *Read) GetFragmentLength() int32 {
	if m != nil {
		return m.FragmentLength
	}
	return 0
}

func (m *Read) GetReadNumber() int32 {
	if m != nil {
		return m.ReadNumber
	}
	return 0
}

func (m *Read) GetNumberReads() int32 {
	if m != nil {
		return m.NumberReads
	}
	return 0
}

func (m *Read) GetFailedVendorQualityChecks() bool {
	if m != nil {
		return m.FailedVendorQualityChecks
	}
	return false
}

func (m *Read) GetAlignment() *LinearAlignment {
	if m != nil {
		return m.Alignment
	}
	return nil
}

func (m *Read) GetSecondaryAlignment() bool {
	if m != nil {
		return m.SecondaryAlignment
	}
	return false
}

func (m *Read) GetSupplementaryAlignment() bool {
	if m != nil {
		return m.SupplementaryAlignment
	}
	return false
}

func (m *Read) GetAlignedSequence() string {
	if m != nil {
		return m.AlignedSequence
	}
	return ""
}

func (m *Read) GetAlignedQuality() []int32 {
	if m != nil {
		return m.AlignedQuality
	}
	return nil
}

func (m *Read) GetNextMatePosition() *Position {
	if m != nil {
		return m.NextMatePosition
	}
	return nil
}

func (m *Read) GetInfo() map[string]*_struct.ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*LinearAlignment)(nil), "google.genomics.v1.LinearAlignment")
	proto.RegisterType((*Read)(nil), "google.genomics.v1.Read")
	proto.RegisterMapType((map[string]*_struct.ListValue)(nil), "google.genomics.v1.Read.InfoEntry")
}

func init() {
	proto.RegisterFile("google/genomics/v1/readalignment.proto", fileDescriptor_readalignment_ff092b0bea51b5ff)
}

var fileDescriptor_readalignment_ff092b0bea51b5ff = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc7, 0xe5, 0x84, 0x70, 0xc9, 0x09, 0x24, 0x61, 0xae, 0xc4, 0xf5, 0x8d, 0xb8, 0xb7, 0x21,
	0x48, 0x6d, 0x58, 0xd4, 0x2e, 0x20, 0xb5, 0x88, 0x2e, 0x2a, 0x40, 0x6d, 0x45, 0x45, 0x51, 0x6a,
	0x54, 0x16, 0xdd, 0x58, 0x83, 0x7d, 0x62, 0x46, 0xd8, 0x33, 0xc6, 0x1e, 0x47, 0xcd, 0x23, 0xf5,
	0xdd, 0xfa, 0x00, 0x5d, 0x56, 0x33, 0xf6, 0x38, 0xd0, 0x66, 0xd1, 0x5d, 0xf2, 0x3f, 0xbf, 0xf3,
	0xe1, 0xf3, 0x31, 0xf0, 0x34, 0x12, 0x22, 0x8a, 0xd1, 0x8d, 0x90, 0x8b, 0x84, 0x05, 0xb9, 0x3b,
	0xdb, 0x77, 0x33, 0xa4, 0x21, 0x8d, 0x59, 0xc4, 0x13, 0xe4, 0xd2, 0x49, 0x33, 0x21, 0x05, 0x21,
	0x25, 0xe7, 0x18, 0xce, 0x99, 0xed, 0x0f, 0xb6, 0x2b, 0x5f, 0x9a, 0x32, 0x97, 0x72, 0x2e, 0x24,
	0x95, 0x4c, 0xf0, 0xbc, 0xf4, 0x18, 0xfc, 0xbf, 0x24, 0x72, 0xc0, 0x22, 0x9a, 0x55, 0xf6, 0x9d,
	0x25, 0xf6, 0x54, 0xe4, 0x4c, 0xc5, 0xa8, 0x10, 0x93, 0x40, 0xff, 0xbb, 0x29, 0xa6, 0x6e, 0x2e,
	0xb3, 0x22, 0xa8, 0x4a, 0x1a, 0x7d, 0xb3, 0xa0, 0x77, 0xc1, 0x38, 0xd2, 0xec, 0xc4, 0x14, 0x4b,
	0x8e, 0x60, 0xcd, 0xc4, 0xb0, 0xad, 0xa1, 0x35, 0xee, 0x1c, 0x6c, 0x3b, 0xbf, 0x57, 0xee, 0x4c,
	0x2a, 0xc6, 0xab, 0x69, 0xf2, 0x0c, 0x7a, 0x09, 0x4d, 0x53, 0xc6, 0x23, 0xff, 0xbe, 0xa0, 0x31,
	0x93, 0x73, 0xbb, 0x31, 0xb4, 0xc6, 0x2d, 0xaf, 0x5b, 0xc9, 0x9f, 0x4a, 0x95, 0x1c, 0x42, 0x4b,
	0x7f, 0x86, 0xdd, 0x1c, 0x36, 0xc7, 0x9d, 0x83, 0xff, 0x96, 0xc5, 0x3f, 0x53, 0xc0, 0x67, 0xce,
	0xa4, 0x57, 0xb2, 0xa3, 0xef, 0xab, 0xb0, 0xe2, 0x21, 0x0d, 0x49, 0x17, 0x1a, 0x2c, 0xd4, 0xa5,
	0xb5, 0xbd, 0x06, 0x0b, 0xc9, 0x08, 0x36, 0x54, 0xbb, 0xfd, 0x28, 0x13, 0x45, 0xea, 0xb3, 0x50,
	0x27, 0x6d, 0x7b, 0x1d, 0x25, 0xbe, 0x57, 0xda, 0x79, 0x48, 0xf6, 0x60, 0xf3, 0x01, 0x93, 0xa3,
	0x54, 0x5c, 0x53, 0x73, 0xdd, 0x9a, 0xbb, 0x42, 0x79, 0x1e, 0x92, 0x5d, 0xd8, 0x98, 0x66, 0x34,
	0x52, 0xbd, 0xf0, 0x39, 0x4d, 0xd0, 0x5e, 0xd1, 0xd8, 0xba, 0x11, 0x2f, 0x69, 0x82, 0x64, 0x0f,
	0xfa, 0x69, 0x26, 0x52, 0xcc, 0xfc, 0x34, 0xa6, 0x01, 0x2a, 0xdd, 0x6e, 0x0d, 0xad, 0xf1, 0x9a,
	0xd7, 0x2b, 0xf5, 0x89, 0x91, 0xc9, 0x73, 0x20, 0x61, 0x91, 0xc6, 0x2c, 0xa0, 0x12, 0x7d, 0x13,
	0xc4, 0x5e, 0xd5, 0xf0, 0x66, 0x6d, 0x79, 0x57, 0x19, 0x54, 0x13, 0xeb, 0xf4, 0x31, 0xf2, 0x48,
	0xde, 0xda, 0x7f, 0x95, 0x4d, 0x34, 0xf2, 0x85, 0x56, 0xc9, 0x13, 0xd0, 0x5f, 0xe8, 0xf3, 0x22,
	0xb9, 0xc1, 0xcc, 0x5e, 0xd3, 0x10, 0x28, 0xe9, 0x52, 0x2b, 0x64, 0x07, 0xd6, 0x4b, 0x9b, 0xaf,
	0xc4, 0xdc, 0x6e, 0x6b, 0xa2, 0x53, 0x6a, 0xaa, 0x93, 0x39, 0x79, 0x03, 0xdb, 0x53, 0xca, 0x62,
	0x0c, 0xfd, 0x19, 0xf2, 0x50, 0x64, 0x66, 0x6e, 0x7e, 0x70, 0x8b, 0xc1, 0x5d, 0x6e, 0x83, 0xae,
	0xf2, 0xdf, 0x92, 0xb9, 0xd6, 0x48, 0x35, 0xc3, 0x33, 0x0d, 0x90, 0x13, 0x68, 0xd7, 0x6b, 0x6e,
	0x77, 0xf4, 0xb6, 0xec, 0x2e, 0x9b, 0xe6, 0x2f, 0x4b, 0xe6, 0x2d, 0xbc, 0x88, 0x0b, 0x7f, 0xe7,
	0x18, 0x08, 0x1e, 0xd2, 0x6c, 0xee, 0x2f, 0x82, 0xad, 0xeb, 0xd4, 0xa4, 0x36, 0x2d, 0x16, 0xf4,
	0x15, 0xfc, 0x93, 0x17, 0x69, 0x1a, 0xeb, 0xf6, 0x3e, 0x76, 0xda, 0xd0, 0x4e, 0x5b, 0x8f, 0xcc,
	0x0b, 0xc7, 0x3d, 0xe8, 0x6b, 0x14, 0x43, 0x3f, 0xc7, 0xfb, 0x02, 0x79, 0x80, 0x76, 0x57, 0x0f,
	0xb7, 0x57, 0xe9, 0x57, 0x95, 0xac, 0xa6, 0x60, 0x50, 0xb3, 0xca, 0xbd, 0x61, 0x53, 0x4d, 0xa1,
	0x92, 0xcd, 0x2a, 0x7f, 0x00, 0xc2, 0xf1, 0xab, 0xf4, 0x13, 0x35, 0xdd, 0xfa, 0x6e, 0xfa, 0x7f,
	0x70, 0x37, 0x7d, 0xe5, 0xf7, 0x91, 0x4a, 0x34, 0x0a, 0x79, 0x09, 0x2b, 0x8c, 0x4f, 0x85, 0xbd,
	0xa9, 0xaf, 0x62, 0xb4, 0xcc, 0x5b, 0x8d, 0xcd, 0x39, 0xe7, 0x53, 0xf1, 0x96, 0xcb, 0x6c, 0xee,
	0x69, 0x7e, 0x70, 0x05, 0xed, 0x5a, 0x22, 0x7d, 0x68, 0xde, 0xe1, 0xbc, 0x3a, 0x0f, 0xf5, 0x93,
	0xbc, 0x80, 0xd6, 0x8c, 0xc6, 0x05, 0xea, 0xbb, 0xe8, 0x1c, 0x0c, 0x4c, 0x5c, 0xf3, 0x24, 0x38,
	0x17, 0x2c, 0x97, 0xd7, 0x8a, 0xf0, 0x4a, 0xf0, 0xb8, 0x71, 0x64, 0x9d, 0x26, 0xb0, 0x15, 0x88,
	0x64, 0x49, 0x0d, 0xa7, 0x44, 0x15, 0x51, 0x77, 0x75, 0xa2, 0xa2, 0x4c, 0xac, 0x2f, 0xc7, 0x86,
	0x14, 0x31, 0xe5, 0x91, 0x23, 0xb2, 0x48, 0x3d, 0x4b, 0x3a, 0x87, 0x5b, 0x9a, 0x68, 0xca, 0xf2,
	0x87, 0x4f, 0xd5, 0x6b, 0xf3, 0xfb, 0x87, 0x65, 0xdd, 0xac, 0x6a, 0xf2, 0xf0, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd0, 0xe1, 0xf6, 0x57, 0x4d, 0x05, 0x00, 0x00,
}
