// Code generated by protoc-gen-go.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=maxMessageCount" json:"maxMessageCount,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absoluteMaxBytes" json:"absoluteMaxBytes,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *BatchTimeout) Reset()                    { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string            { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()               {}
func (*BatchTimeout) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// When submitting a new chain configuration transaction to create a new chain,
// the first configuration item must be of type Orderer with Key CreationPolicy
// and contents of a Marshaled CreationPolicy. The policy should be set to the
// policy which was supplied by the ordering service for the client's chain
// creation. The digest should be the hash of the concatenation of the remaining
// ConfigurationItem bytes. The signatures of the configuration item should
// satisfy the policy for chain creation.
type CreationPolicy struct {
	// The name of the policy which should be used to validate the creation of
	// this chain
	Policy string `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
	// The hash of the concatenation of remaining configuration item bytes
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *CreationPolicy) Reset()                    { *m = CreationPolicy{} }
func (m *CreationPolicy) String() string            { return proto.CompactTextString(m) }
func (*CreationPolicy) ProtoMessage()               {}
func (*CreationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// IngressPolicy is the name of the policy which incoming Broadcast messages are filtered against
type IngressPolicy struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *IngressPolicy) Reset()                    { *m = IngressPolicy{} }
func (m *IngressPolicy) String() string            { return proto.CompactTextString(m) }
func (*IngressPolicy) ProtoMessage()               {}
func (*IngressPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// EgressPolicy is the name of the policy which incoming Deliver messages are filtered against
type EgressPolicy struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *EgressPolicy) Reset()                    { *m = EgressPolicy{} }
func (m *EgressPolicy) String() string            { return proto.CompactTextString(m) }
func (*EgressPolicy) ProtoMessage()               {}
func (*EgressPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ChainCreators struct {
	// A list of policies, any of which may be specified as the chain creation
	// policy in a chain creation request
	Policies []string `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *ChainCreators) Reset()                    { *m = ChainCreators{} }
func (m *ChainCreators) String() string            { return proto.CompactTextString(m) }
func (*ChainCreators) ProtoMessage()               {}
func (*ChainCreators) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*CreationPolicy)(nil), "orderer.CreationPolicy")
	proto.RegisterType((*IngressPolicy)(nil), "orderer.IngressPolicy")
	proto.RegisterType((*EgressPolicy)(nil), "orderer.EgressPolicy")
	proto.RegisterType((*ChainCreators)(nil), "orderer.ChainCreators")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x41, 0x6b, 0xe3, 0x30,
	0x10, 0x85, 0xc9, 0x6e, 0x48, 0x36, 0x43, 0xb2, 0xbb, 0x68, 0x61, 0x31, 0xe9, 0x25, 0xb8, 0x17,
	0xd3, 0x96, 0xf8, 0xd0, 0x3f, 0x50, 0x6c, 0x7a, 0x28, 0x25, 0x50, 0xdc, 0x9c, 0x7a, 0x93, 0x9d,
	0x89, 0x2d, 0x12, 0x6b, 0x8c, 0x46, 0x86, 0xb8, 0xbf, 0xbe, 0x58, 0x56, 0x72, 0x68, 0x0f, 0x3d,
	0xe9, 0x7d, 0x33, 0x0f, 0xcd, 0x3c, 0x06, 0xae, 0xc8, 0xec, 0xd0, 0xa0, 0x89, 0x0b, 0xd2, 0x7b,
	0x55, 0xb6, 0x46, 0x5a, 0x45, 0x7a, 0xdd, 0x18, 0xb2, 0x24, 0xa6, 0xbe, 0xb9, 0xfc, 0x57, 0x50,
	0x5d, 0x93, 0x8e, 0x87, 0x67, 0xe8, 0x86, 0xd7, 0xb0, 0x48, 0x49, 0x33, 0x6a, 0x6e, 0x79, 0xdb,
	0x35, 0x28, 0x04, 0x8c, 0x6d, 0xd7, 0x60, 0x30, 0x5a, 0x8d, 0xa2, 0x59, 0xe6, 0x74, 0x28, 0x61,
	0x96, 0x48, 0x5b, 0x54, 0xaf, 0xea, 0x1d, 0x45, 0x04, 0x7f, 0x6a, 0x79, 0xda, 0x20, 0xb3, 0x2c,
	0x31, 0xa5, 0x56, 0x5b, 0xe7, 0x5d, 0x64, 0x9f, 0xcb, 0xe2, 0x06, 0xfe, 0xca, 0x9c, 0xe9, 0xd8,
	0x5a, 0xdc, 0xc8, 0x53, 0xd2, 0x59, 0xe4, 0xe0, 0x87, 0xb3, 0x7e, 0xa9, 0x87, 0x11, 0xcc, 0xdd,
	0x88, 0xad, 0xaa, 0x91, 0x5a, 0x2b, 0x02, 0x98, 0xda, 0x41, 0xfa, 0x4d, 0xce, 0x18, 0x3e, 0xc0,
	0xef, 0xd4, 0xa0, 0x4b, 0xf8, 0x42, 0x47, 0x55, 0x74, 0xe2, 0x3f, 0x4c, 0x1a, 0xa7, 0xbc, 0xd5,
	0x53, 0x5f, 0xdf, 0xa9, 0x12, 0xd9, 0xba, 0xa9, 0xf3, 0xcc, 0x53, 0x9f, 0xf9, 0x49, 0x97, 0x06,
	0x99, 0xfd, 0x07, 0x02, 0xc6, 0x5a, 0xd6, 0x97, 0xcc, 0xbd, 0x0e, 0x43, 0x98, 0x3f, 0x7e, 0xe7,
	0xb9, 0x85, 0x45, 0x5a, 0x49, 0xa5, 0xdd, 0x3e, 0x64, 0x58, 0x2c, 0xe1, 0x97, 0x9b, 0xad, 0x90,
	0x83, 0xd1, 0xea, 0x67, 0x34, 0xcb, 0x2e, 0xdc, 0x27, 0x7c, 0x96, 0xfb, 0x83, 0x4c, 0x0c, 0x1d,
	0xd0, 0x70, 0x9f, 0x30, 0x1f, 0xa4, 0xb7, 0x9e, 0x31, 0x59, 0xbf, 0xdd, 0x95, 0xca, 0x56, 0x6d,
	0xbe, 0x2e, 0xa8, 0x8e, 0xab, 0xae, 0x41, 0x73, 0xc4, 0x5d, 0x89, 0x26, 0xde, 0xcb, 0xdc, 0xa8,
	0x22, 0x76, 0xa7, 0xe3, 0xd8, 0x1f, 0x36, 0x9f, 0x38, 0xbe, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x6e, 0xaf, 0xb7, 0x25, 0x07, 0x02, 0x00, 0x00,
}