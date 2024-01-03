// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayer/chains/ethereum/config/config.proto

package ethereum

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type ChainConfig struct {
	ChainId                string     `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	EthChainId             uint64     `protobuf:"varint,2,opt,name=eth_chain_id,json=ethChainId,proto3" json:"eth_chain_id,omitempty"`
	RpcAddr                string     `protobuf:"bytes,3,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
	Signer                 *types.Any `protobuf:"bytes,4,opt,name=signer,proto3" json:"signer,omitempty"`
	IbcAddress             string     `protobuf:"bytes,5,opt,name=ibc_address,json=ibcAddress,proto3" json:"ibc_address,omitempty"`
	InitialSendCheckpoint  uint64     `protobuf:"varint,6,opt,name=initial_send_checkpoint,json=initialSendCheckpoint,proto3" json:"initial_send_checkpoint,omitempty"`
	InitialRecvCheckpoint  uint64     `protobuf:"varint,7,opt,name=initial_recv_checkpoint,json=initialRecvCheckpoint,proto3" json:"initial_recv_checkpoint,omitempty"`
	EnableDebugTrace       bool       `protobuf:"varint,8,opt,name=enable_debug_trace,json=enableDebugTrace,proto3" json:"enable_debug_trace,omitempty"`
	AverageBlockTimeMsec   uint64     `protobuf:"varint,9,opt,name=average_block_time_msec,json=averageBlockTimeMsec,proto3" json:"average_block_time_msec,omitempty"`
	MaxRetryForInclusion   uint64     `protobuf:"varint,10,opt,name=max_retry_for_inclusion,json=maxRetryForInclusion,proto3" json:"max_retry_for_inclusion,omitempty"`
	EnableLegacyTx         bool       `protobuf:"varint,11,opt,name=enable_legacy_tx,json=enableLegacyTx,proto3" json:"enable_legacy_tx,omitempty"`
	LimitPriorityFeePerGas string     `protobuf:"bytes,12,opt,name=limit_priority_fee_per_gas,json=limitPriorityFeePerGas,proto3" json:"limit_priority_fee_per_gas,omitempty"`
	PriorityFeeRate        *Fraction  `protobuf:"bytes,13,opt,name=priority_fee_rate,json=priorityFeeRate,proto3" json:"priority_fee_rate,omitempty"`
	LimitFeePerGas         string     `protobuf:"bytes,14,opt,name=limit_fee_per_gas,json=limitFeePerGas,proto3" json:"limit_fee_per_gas,omitempty"`
	BaseFeeRate            *Fraction  `protobuf:"bytes,15,opt,name=base_fee_rate,json=baseFeeRate,proto3" json:"base_fee_rate,omitempty"`
	RewardPercentile       float32    `protobuf:"fixed32,16,opt,name=rewardPercentile,proto3" json:"rewardPercentile,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8a57ab2f9f14837, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

type Fraction struct {
	Numerator   uint64 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator uint64 `protobuf:"varint,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
}

func (m *Fraction) Reset()         { *m = Fraction{} }
func (m *Fraction) String() string { return proto.CompactTextString(m) }
func (*Fraction) ProtoMessage()    {}
func (*Fraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8a57ab2f9f14837, []int{1}
}
func (m *Fraction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fraction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fraction.Merge(m, src)
}
func (m *Fraction) XXX_Size() int {
	return m.Size()
}
func (m *Fraction) XXX_DiscardUnknown() {
	xxx_messageInfo_Fraction.DiscardUnknown(m)
}

var xxx_messageInfo_Fraction proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainConfig)(nil), "relayer.chains.ethereum.config.ChainConfig")
	proto.RegisterType((*Fraction)(nil), "relayer.chains.ethereum.config.Fraction")
}

func init() {
	proto.RegisterFile("relayer/chains/ethereum/config/config.proto", fileDescriptor_a8a57ab2f9f14837)
}

var fileDescriptor_a8a57ab2f9f14837 = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x6e, 0xd3, 0x4a,
	0x18, 0x8e, 0x7b, 0x7a, 0xd2, 0x74, 0xd2, 0xeb, 0xa8, 0xe7, 0xd4, 0xad, 0x90, 0x89, 0xba, 0x32,
	0xd0, 0xda, 0x12, 0x08, 0x16, 0xec, 0xda, 0xa0, 0xa2, 0xa2, 0x22, 0x55, 0x26, 0x2b, 0x36, 0xa3,
	0xf1, 0xf8, 0x8f, 0x33, 0xaa, 0x3d, 0x63, 0xfd, 0x9e, 0x94, 0xe4, 0x2d, 0x78, 0x05, 0xde, 0xa6,
	0xcb, 0x2e, 0x59, 0x42, 0xfb, 0x22, 0xc8, 0x63, 0xe7, 0x52, 0x21, 0x21, 0xb1, 0xb2, 0xe7, 0xbb,
	0xfd, 0xdf, 0xd8, 0x1e, 0x93, 0x17, 0x08, 0x19, 0x9f, 0x02, 0x86, 0x62, 0xc4, 0xa5, 0x2a, 0x43,
	0x30, 0x23, 0x40, 0x18, 0xe7, 0xa1, 0xd0, 0x6a, 0x28, 0xd3, 0xe6, 0x12, 0x14, 0xa8, 0x8d, 0xa6,
	0x5e, 0x23, 0x0e, 0x6a, 0x71, 0x30, 0x13, 0x07, 0xb5, 0xea, 0x70, 0x2f, 0xd5, 0xa9, 0xb6, 0xd2,
	0xb0, 0xba, 0xab, 0x5d, 0x87, 0x07, 0xa9, 0xd6, 0x69, 0x06, 0xa1, 0x5d, 0xc5, 0xe3, 0x61, 0xc8,
	0xd5, 0xb4, 0xa6, 0x8e, 0xbe, 0xb5, 0x49, 0xb7, 0x5f, 0x65, 0xf5, 0x6d, 0x00, 0x3d, 0x20, 0x1d,
	0x1b, 0xcd, 0x64, 0xe2, 0x3a, 0x3d, 0xc7, 0x5f, 0x8f, 0xd6, 0xec, 0xfa, 0x22, 0xa1, 0x3d, 0xb2,
	0x01, 0x66, 0xc4, 0xe6, 0xf4, 0x4a, 0xcf, 0xf1, 0x57, 0x23, 0x02, 0x66, 0xd4, 0x6f, 0x14, 0x07,
	0xa4, 0x83, 0x85, 0x60, 0x3c, 0x49, 0xd0, 0xfd, 0xa7, 0x36, 0x63, 0x21, 0x4e, 0x93, 0x04, 0xe9,
	0x31, 0x69, 0x97, 0x32, 0x55, 0x80, 0xee, 0x6a, 0xcf, 0xf1, 0xbb, 0x2f, 0xf7, 0x82, 0xba, 0x53,
	0x30, 0xeb, 0x14, 0x9c, 0xaa, 0x69, 0xd4, 0x68, 0xe8, 0x53, 0xd2, 0x95, 0x71, 0x1d, 0x04, 0x65,
	0xe9, 0xfe, 0x6b, 0xb3, 0x88, 0x8c, 0x6d, 0x16, 0x94, 0x25, 0x7d, 0x43, 0xf6, 0xa5, 0x92, 0x46,
	0xf2, 0x8c, 0x95, 0xa0, 0x12, 0x26, 0x46, 0x20, 0xae, 0x0b, 0x2d, 0x95, 0x71, 0xdb, 0xb6, 0xd6,
	0x7f, 0x0d, 0xfd, 0x09, 0x54, 0xd2, 0x9f, 0x93, 0xcb, 0x3e, 0x04, 0x71, 0xb3, 0xec, 0x5b, 0x7b,
	0xe4, 0x8b, 0x40, 0xdc, 0x2c, 0xf9, 0x8e, 0x09, 0x05, 0xc5, 0xe3, 0x0c, 0x58, 0x02, 0xf1, 0x38,
	0x65, 0x06, 0xb9, 0x00, 0xb7, 0xd3, 0x73, 0xfc, 0x4e, 0xb4, 0x53, 0x33, 0xef, 0x2a, 0x62, 0x50,
	0xe1, 0xf4, 0x35, 0xd9, 0xe7, 0x37, 0x80, 0x3c, 0x05, 0x16, 0x67, 0x5a, 0x5c, 0x33, 0x23, 0x73,
	0x60, 0x79, 0x09, 0xc2, 0x5d, 0xb7, 0x53, 0xf6, 0x1a, 0xfa, 0xac, 0x62, 0x07, 0x32, 0x87, 0x8f,
	0x25, 0x88, 0xca, 0x96, 0xf3, 0x09, 0x43, 0x30, 0x38, 0x65, 0x43, 0x8d, 0x4c, 0x2a, 0x91, 0x8d,
	0x4b, 0xa9, 0x95, 0x4b, 0x6a, 0x5b, 0xce, 0x27, 0x51, 0xc5, 0x9e, 0x6b, 0xbc, 0x98, 0x71, 0xd4,
	0x27, 0x4d, 0x03, 0x96, 0x41, 0xca, 0xc5, 0x94, 0x99, 0x89, 0xdb, 0xb5, 0xcd, 0xb6, 0x6a, 0xfc,
	0xd2, 0xc2, 0x83, 0x09, 0x7d, 0x4b, 0x0e, 0x33, 0x99, 0x4b, 0xc3, 0x0a, 0x94, 0x1a, 0xa5, 0x99,
	0xb2, 0x21, 0x00, 0x2b, 0x00, 0x59, 0xca, 0x4b, 0x77, 0xc3, 0x3e, 0xe5, 0xff, 0xad, 0xe2, 0xaa,
	0x11, 0x9c, 0x03, 0x5c, 0x01, 0xbe, 0xe7, 0x25, 0x1d, 0x90, 0xdd, 0x47, 0x2e, 0xe4, 0x06, 0xdc,
	0x4d, 0xfb, 0x2e, 0xfd, 0xe0, 0xcf, 0x5f, 0x65, 0x70, 0x8e, 0x5c, 0x18, 0xa9, 0x55, 0xb4, 0x5d,
	0x2c, 0x72, 0x23, 0x6e, 0x80, 0x3e, 0x23, 0xbb, 0x75, 0xa3, 0xe5, 0x22, 0x5b, 0xb6, 0xc8, 0x96,
	0x25, 0x16, 0x05, 0x2e, 0xc9, 0x66, 0xcc, 0x4b, 0x58, 0x0c, 0xdf, 0xfe, 0xcb, 0xe1, 0xdd, 0xca,
	0x3e, 0x1b, 0xfc, 0x9c, 0xec, 0x20, 0x7c, 0xe1, 0x98, 0x5c, 0x01, 0x0a, 0x50, 0x46, 0x66, 0xe0,
	0xee, 0xf4, 0x1c, 0x7f, 0x25, 0xfa, 0x0d, 0x3f, 0xfa, 0x40, 0x3a, 0xb3, 0x10, 0xfa, 0x84, 0xac,
	0xab, 0x71, 0x0e, 0xc8, 0x8d, 0x46, 0x7b, 0x40, 0x56, 0xa3, 0x05, 0x40, 0x7b, 0xa4, 0x9b, 0x80,
	0xd2, 0xb9, 0x54, 0x96, 0xaf, 0x4f, 0xc8, 0x32, 0x74, 0xc6, 0x6f, 0x7f, 0x7a, 0xad, 0xdb, 0x7b,
	0xcf, 0xb9, 0xbb, 0xf7, 0x9c, 0x1f, 0xf7, 0x9e, 0xf3, 0xf5, 0xc1, 0x6b, 0xdd, 0x3d, 0x78, 0xad,
	0xef, 0x0f, 0x5e, 0xeb, 0x73, 0x3f, 0x95, 0x66, 0x34, 0x8e, 0x03, 0xa1, 0xf3, 0x30, 0xe1, 0x86,
	0xdb, 0x2d, 0x65, 0x3c, 0x9e, 0xff, 0x14, 0x4e, 0x64, 0x2c, 0x4e, 0xec, 0x86, 0x4f, 0x2c, 0x17,
	0x16, 0xd7, 0x69, 0x68, 0xd7, 0x73, 0x49, 0xdc, 0xb6, 0x47, 0xea, 0xd5, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0xa7, 0x92, 0xab, 0x59, 0x04, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RewardPercentile != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.RewardPercentile))))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x85
	}
	if m.BaseFeeRate != nil {
		{
			size, err := m.BaseFeeRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x7a
	}
	if len(m.LimitFeePerGas) > 0 {
		i -= len(m.LimitFeePerGas)
		copy(dAtA[i:], m.LimitFeePerGas)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.LimitFeePerGas)))
		i--
		dAtA[i] = 0x72
	}
	if m.PriorityFeeRate != nil {
		{
			size, err := m.PriorityFeeRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if len(m.LimitPriorityFeePerGas) > 0 {
		i -= len(m.LimitPriorityFeePerGas)
		copy(dAtA[i:], m.LimitPriorityFeePerGas)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.LimitPriorityFeePerGas)))
		i--
		dAtA[i] = 0x62
	}
	if m.EnableLegacyTx {
		i--
		if m.EnableLegacyTx {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.MaxRetryForInclusion != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxRetryForInclusion))
		i--
		dAtA[i] = 0x50
	}
	if m.AverageBlockTimeMsec != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.AverageBlockTimeMsec))
		i--
		dAtA[i] = 0x48
	}
	if m.EnableDebugTrace {
		i--
		if m.EnableDebugTrace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.InitialRecvCheckpoint != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.InitialRecvCheckpoint))
		i--
		dAtA[i] = 0x38
	}
	if m.InitialSendCheckpoint != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.InitialSendCheckpoint))
		i--
		dAtA[i] = 0x30
	}
	if len(m.IbcAddress) > 0 {
		i -= len(m.IbcAddress)
		copy(dAtA[i:], m.IbcAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Signer != nil {
		{
			size, err := m.Signer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.RpcAddr) > 0 {
		i -= len(m.RpcAddr)
		copy(dAtA[i:], m.RpcAddr)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RpcAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.EthChainId != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.EthChainId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Fraction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fraction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fraction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Denominator != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Denominator))
		i--
		dAtA[i] = 0x10
	}
	if m.Numerator != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Numerator))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.EthChainId != 0 {
		n += 1 + sovConfig(uint64(m.EthChainId))
	}
	l = len(m.RpcAddr)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Signer != nil {
		l = m.Signer.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IbcAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.InitialSendCheckpoint != 0 {
		n += 1 + sovConfig(uint64(m.InitialSendCheckpoint))
	}
	if m.InitialRecvCheckpoint != 0 {
		n += 1 + sovConfig(uint64(m.InitialRecvCheckpoint))
	}
	if m.EnableDebugTrace {
		n += 2
	}
	if m.AverageBlockTimeMsec != 0 {
		n += 1 + sovConfig(uint64(m.AverageBlockTimeMsec))
	}
	if m.MaxRetryForInclusion != 0 {
		n += 1 + sovConfig(uint64(m.MaxRetryForInclusion))
	}
	if m.EnableLegacyTx {
		n += 2
	}
	l = len(m.LimitPriorityFeePerGas)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.PriorityFeeRate != nil {
		l = m.PriorityFeeRate.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.LimitFeePerGas)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.BaseFeeRate != nil {
		l = m.BaseFeeRate.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.RewardPercentile != 0 {
		n += 6
	}
	return n
}

func (m *Fraction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Numerator != 0 {
		n += 1 + sovConfig(uint64(m.Numerator))
	}
	if m.Denominator != 0 {
		n += 1 + sovConfig(uint64(m.Denominator))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthChainId", wireType)
			}
			m.EthChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signer == nil {
				m.Signer = &types.Any{}
			}
			if err := m.Signer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSendCheckpoint", wireType)
			}
			m.InitialSendCheckpoint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialSendCheckpoint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialRecvCheckpoint", wireType)
			}
			m.InitialRecvCheckpoint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialRecvCheckpoint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableDebugTrace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableDebugTrace = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageBlockTimeMsec", wireType)
			}
			m.AverageBlockTimeMsec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AverageBlockTimeMsec |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetryForInclusion", wireType)
			}
			m.MaxRetryForInclusion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRetryForInclusion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableLegacyTx", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableLegacyTx = bool(v != 0)
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitPriorityFeePerGas", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LimitPriorityFeePerGas = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriorityFeeRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PriorityFeeRate == nil {
				m.PriorityFeeRate = &Fraction{}
			}
			if err := m.PriorityFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitFeePerGas", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LimitFeePerGas = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFeeRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseFeeRate == nil {
				m.BaseFeeRate = &Fraction{}
			}
			if err := m.BaseFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPercentile", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.RewardPercentile = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func (m *Fraction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Fraction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fraction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Numerator", wireType)
			}
			m.Numerator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Numerator |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denominator", wireType)
			}
			m.Denominator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Denominator |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
