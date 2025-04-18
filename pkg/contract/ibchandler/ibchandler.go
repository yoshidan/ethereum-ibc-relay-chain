// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibchandler

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ChannelCounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type ChannelCounterpartyData struct {
	PortId    string
	ChannelId string
}

// ChannelData is an auto generated low-level Go binding around an user-defined struct.
type ChannelData struct {
	State          uint8
	Ordering       uint8
	Counterparty   ChannelCounterpartyData
	ConnectionHops []string
	Version        string
}

// ConnectionEndData is an auto generated low-level Go binding around an user-defined struct.
type ConnectionEndData struct {
	ClientId     string
	Versions     []VersionData
	State        uint8
	Counterparty CounterpartyData
	DelayPeriod  uint64
}

// CounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type CounterpartyData struct {
	ClientId     string
	ConnectionId string
	Prefix       MerklePrefixData
}

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IIBCChannelAcknowledgePacketMsgPacketAcknowledgement is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelAcknowledgePacketMsgPacketAcknowledgement struct {
	Packet          PacketData
	Acknowledgement []byte
	Proof           []byte
	ProofHeight     HeightData
}

// IIBCChannelHandshakeMsgChannelCloseConfirm is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelHandshakeMsgChannelCloseConfirm struct {
	PortId      string
	ChannelId   string
	ProofInit   []byte
	ProofHeight HeightData
}

// IIBCChannelHandshakeMsgChannelCloseInit is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelHandshakeMsgChannelCloseInit struct {
	PortId    string
	ChannelId string
}

// IIBCChannelHandshakeMsgChannelOpenAck is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelHandshakeMsgChannelOpenAck struct {
	PortId                string
	ChannelId             string
	CounterpartyVersion   string
	CounterpartyChannelId string
	ProofTry              []byte
	ProofHeight           HeightData
}

// IIBCChannelHandshakeMsgChannelOpenConfirm is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelHandshakeMsgChannelOpenConfirm struct {
	PortId      string
	ChannelId   string
	ProofAck    []byte
	ProofHeight HeightData
}

// IIBCChannelHandshakeMsgChannelOpenInit is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelHandshakeMsgChannelOpenInit struct {
	PortId  string
	Channel ChannelData
}

// IIBCChannelHandshakeMsgChannelOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelHandshakeMsgChannelOpenTry struct {
	PortId              string
	Channel             ChannelData
	CounterpartyVersion string
	ProofInit           []byte
	ProofHeight         HeightData
}

// IIBCChannelPacketTimeoutMsgTimeoutOnClose is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelPacketTimeoutMsgTimeoutOnClose struct {
	Packet           PacketData
	ProofUnreceived  []byte
	ProofClose       []byte
	ProofHeight      HeightData
	NextSequenceRecv uint64
}

// IIBCChannelPacketTimeoutMsgTimeoutPacket is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelPacketTimeoutMsgTimeoutPacket struct {
	Packet           PacketData
	Proof            []byte
	ProofHeight      HeightData
	NextSequenceRecv uint64
}

// IIBCChannelRecvPacketMsgPacketRecv is an auto generated low-level Go binding around an user-defined struct.
type IIBCChannelRecvPacketMsgPacketRecv struct {
	Packet      PacketData
	Proof       []byte
	ProofHeight HeightData
}

// IIBCClientMsgCreateClient is an auto generated low-level Go binding around an user-defined struct.
type IIBCClientMsgCreateClient struct {
	ClientType          string
	ClientStateBytes    []byte
	ConsensusStateBytes []byte
}

// IIBCClientMsgUpdateClient is an auto generated low-level Go binding around an user-defined struct.
type IIBCClientMsgUpdateClient struct {
	ClientId      string
	ClientMessage []byte
}

// IIBCConnectionMsgConnectionOpenAck is an auto generated low-level Go binding around an user-defined struct.
type IIBCConnectionMsgConnectionOpenAck struct {
	ConnectionId             string
	ClientStateBytes         []byte
	Version                  VersionData
	CounterpartyConnectionId string
	ProofTry                 []byte
	ProofClient              []byte
	ProofConsensus           []byte
	ProofHeight              HeightData
	ConsensusHeight          HeightData
}

// IIBCConnectionMsgConnectionOpenConfirm is an auto generated low-level Go binding around an user-defined struct.
type IIBCConnectionMsgConnectionOpenConfirm struct {
	ConnectionId string
	ProofAck     []byte
	ProofHeight  HeightData
}

// IIBCConnectionMsgConnectionOpenInit is an auto generated low-level Go binding around an user-defined struct.
type IIBCConnectionMsgConnectionOpenInit struct {
	ClientId     string
	Counterparty CounterpartyData
	Version      VersionData
	DelayPeriod  uint64
}

// IIBCConnectionMsgConnectionOpenTry is an auto generated low-level Go binding around an user-defined struct.
type IIBCConnectionMsgConnectionOpenTry struct {
	Counterparty         CounterpartyData
	DelayPeriod          uint64
	ClientId             string
	ClientStateBytes     []byte
	CounterpartyVersions []VersionData
	ProofInit            []byte
	ProofClient          []byte
	ProofConsensus       []byte
	ProofHeight          HeightData
	ConsensusHeight      HeightData
}

// MerklePrefixData is an auto generated low-level Go binding around an user-defined struct.
type MerklePrefixData struct {
	KeyPrefix []byte
}

// PacketData is an auto generated low-level Go binding around an user-defined struct.
type PacketData struct {
	Sequence           uint64
	SourcePort         string
	SourceChannel      string
	DestinationPort    string
	DestinationChannel string
	Data               []byte
	TimeoutHeight      HeightData
	TimeoutTimestamp   uint64
}

// VersionData is an auto generated low-level Go binding around an user-defined struct.
type VersionData struct {
	Identifier string
	Features   []string
}

// IbchandlerMetaData contains all meta data concerning the Ibchandler contract.
var IbchandlerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acknowledgePacket\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelAcknowledgePacket.MsgPacketAcknowledgement\",\"components\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bindPort\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"moduleAddress\",\"type\":\"address\",\"internalType\":\"contractIIBCModule\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelCloseConfirm\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelHandshake.MsgChannelCloseConfirm\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proofInit\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelCloseInit\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelHandshake.MsgChannelCloseInit\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenAck\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelHandshake.MsgChannelOpenAck\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proofTry\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenConfirm\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelHandshake.MsgChannelOpenConfirm\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proofAck\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenInit\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelHandshake.MsgChannelOpenInit\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel\",\"type\":\"tuple\",\"internalType\":\"structChannel.Data\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumChannel.State\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannel.Order\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelCounterparty.Data\",\"components\":[{\"name\":\"port_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel_id\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"connection_hops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenTry\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelHandshake.MsgChannelOpenTry\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel\",\"type\":\"tuple\",\"internalType\":\"structChannel.Data\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumChannel.State\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannel.Order\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelCounterparty.Data\",\"components\":[{\"name\":\"port_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel_id\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"connection_hops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proofInit\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectionOpenAck\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCConnection.MsgConnectionOpenAck\",\"components\":[{\"name\":\"connectionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"clientStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"version\",\"type\":\"tuple\",\"internalType\":\"structVersion.Data\",\"components\":[{\"name\":\"identifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"features\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"counterpartyConnectionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proofTry\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofClient\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofConsensus\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"consensusHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectionOpenConfirm\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCConnection.MsgConnectionOpenConfirm\",\"components\":[{\"name\":\"connectionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proofAck\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectionOpenInit\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCConnection.MsgConnectionOpenInit\",\"components\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structCounterparty.Data\",\"components\":[{\"name\":\"client_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"connection_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"prefix\",\"type\":\"tuple\",\"internalType\":\"structMerklePrefix.Data\",\"components\":[{\"name\":\"key_prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"version\",\"type\":\"tuple\",\"internalType\":\"structVersion.Data\",\"components\":[{\"name\":\"identifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"features\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"delayPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectionOpenTry\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCConnection.MsgConnectionOpenTry\",\"components\":[{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structCounterparty.Data\",\"components\":[{\"name\":\"client_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"connection_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"prefix\",\"type\":\"tuple\",\"internalType\":\"structMerklePrefix.Data\",\"components\":[{\"name\":\"key_prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"delayPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"clientStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"counterpartyVersions\",\"type\":\"tuple[]\",\"internalType\":\"structVersion.Data[]\",\"components\":[{\"name\":\"identifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"features\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"proofInit\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofClient\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofConsensus\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"consensusHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createClient\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCClient.MsgCreateClient\",\"components\":[{\"name\":\"clientType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"clientStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"consensusStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getChannel\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structChannel.Data\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumChannel.State\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannel.Order\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelCounterparty.Data\",\"components\":[{\"name\":\"port_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channel_id\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"connection_hops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClientByType\",\"inputs\":[{\"name\":\"clientType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClientState\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClientType\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCommitment\",\"inputs\":[{\"name\":\"hashedPath\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCommitmentPrefix\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConnection\",\"inputs\":[{\"name\":\"connectionId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structConnectionEnd.Data\",\"components\":[{\"name\":\"client_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"versions\",\"type\":\"tuple[]\",\"internalType\":\"structVersion.Data[]\",\"components\":[{\"name\":\"identifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"features\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumConnectionEnd.State\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structCounterparty.Data\",\"components\":[{\"name\":\"client_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"connection_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"prefix\",\"type\":\"tuple\",\"internalType\":\"structMerklePrefix.Data\",\"components\":[{\"name\":\"key_prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"delay_period\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsensusState\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"consensusStateBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedTimePerBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextSequenceAck\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextSequenceRecv\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextSequenceSend\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPacketReceipt\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIBCChannelLib.PacketReceipt\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recvPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelRecvPacket.MsgPacketRecv\",\"components\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerClient\",\"inputs\":[{\"name\":\"clientType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"contractILightClient\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedTimePerBlock\",\"inputs\":[{\"name\":\"expectedTimePerBlock_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"timeoutOnClose\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelPacketTimeout.MsgTimeoutOnClose\",\"components\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proofUnreceived\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofClose\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"nextSequenceRecv\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"timeoutPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCChannelPacketTimeout.MsgTimeoutPacket\",\"components\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"nextSequenceRecv\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIBCClient.MsgUpdateClient\",\"components\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"clientMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeAcknowledgement\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AcknowledgePacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneratedChannelIdentifier\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneratedClientIdentifier\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneratedConnectionIdentifier\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPacket\",\"inputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structPacket.Data\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_port\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destination_channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeout_height\",\"type\":\"tuple\",\"internalType\":\"structHeight.Data\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeout_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WriteAcknowledgement\",\"inputs\":[{\"name\":\"destinationPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destinationChannel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// IbchandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IbchandlerMetaData.ABI instead.
var IbchandlerABI = IbchandlerMetaData.ABI

// Ibchandler is an auto generated Go binding around an Ethereum contract.
type Ibchandler struct {
	IbchandlerCaller     // Read-only binding to the contract
	IbchandlerTransactor // Write-only binding to the contract
	IbchandlerFilterer   // Log filterer for contract events
}

// IbchandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbchandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbchandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbchandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbchandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbchandlerSession struct {
	Contract     *Ibchandler       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbchandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbchandlerCallerSession struct {
	Contract *IbchandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IbchandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbchandlerTransactorSession struct {
	Contract     *IbchandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IbchandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbchandlerRaw struct {
	Contract *Ibchandler // Generic contract binding to access the raw methods on
}

// IbchandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbchandlerCallerRaw struct {
	Contract *IbchandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IbchandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbchandlerTransactorRaw struct {
	Contract *IbchandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbchandler creates a new instance of Ibchandler, bound to a specific deployed contract.
func NewIbchandler(address common.Address, backend bind.ContractBackend) (*Ibchandler, error) {
	contract, err := bindIbchandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibchandler{IbchandlerCaller: IbchandlerCaller{contract: contract}, IbchandlerTransactor: IbchandlerTransactor{contract: contract}, IbchandlerFilterer: IbchandlerFilterer{contract: contract}}, nil
}

// NewIbchandlerCaller creates a new read-only instance of Ibchandler, bound to a specific deployed contract.
func NewIbchandlerCaller(address common.Address, caller bind.ContractCaller) (*IbchandlerCaller, error) {
	contract, err := bindIbchandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbchandlerCaller{contract: contract}, nil
}

// NewIbchandlerTransactor creates a new write-only instance of Ibchandler, bound to a specific deployed contract.
func NewIbchandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IbchandlerTransactor, error) {
	contract, err := bindIbchandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbchandlerTransactor{contract: contract}, nil
}

// NewIbchandlerFilterer creates a new log filterer instance of Ibchandler, bound to a specific deployed contract.
func NewIbchandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IbchandlerFilterer, error) {
	contract, err := bindIbchandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbchandlerFilterer{contract: contract}, nil
}

// bindIbchandler binds a generic wrapper to an already deployed contract.
func bindIbchandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbchandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibchandler *IbchandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibchandler.Contract.IbchandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibchandler *IbchandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchandler.Contract.IbchandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibchandler *IbchandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibchandler.Contract.IbchandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibchandler *IbchandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibchandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibchandler *IbchandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibchandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibchandler *IbchandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibchandler.Contract.contract.Transact(opts, method, params...)
}

// GetChannel is a free data retrieval call binding the contract method 0x3000217a.
//
// Solidity: function getChannel(string portId, string channelId) view returns((uint8,uint8,(string,string),string[],string), bool)
func (_Ibchandler *IbchandlerCaller) GetChannel(opts *bind.CallOpts, portId string, channelId string) (ChannelData, bool, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getChannel", portId, channelId)

	if err != nil {
		return *new(ChannelData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ChannelData)).(*ChannelData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetChannel is a free data retrieval call binding the contract method 0x3000217a.
//
// Solidity: function getChannel(string portId, string channelId) view returns((uint8,uint8,(string,string),string[],string), bool)
func (_Ibchandler *IbchandlerSession) GetChannel(portId string, channelId string) (ChannelData, bool, error) {
	return _Ibchandler.Contract.GetChannel(&_Ibchandler.CallOpts, portId, channelId)
}

// GetChannel is a free data retrieval call binding the contract method 0x3000217a.
//
// Solidity: function getChannel(string portId, string channelId) view returns((uint8,uint8,(string,string),string[],string), bool)
func (_Ibchandler *IbchandlerCallerSession) GetChannel(portId string, channelId string) (ChannelData, bool, error) {
	return _Ibchandler.Contract.GetChannel(&_Ibchandler.CallOpts, portId, channelId)
}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_Ibchandler *IbchandlerCaller) GetClient(opts *bind.CallOpts, clientId string) (common.Address, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getClient", clientId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_Ibchandler *IbchandlerSession) GetClient(clientId string) (common.Address, error) {
	return _Ibchandler.Contract.GetClient(&_Ibchandler.CallOpts, clientId)
}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_Ibchandler *IbchandlerCallerSession) GetClient(clientId string) (common.Address, error) {
	return _Ibchandler.Contract.GetClient(&_Ibchandler.CallOpts, clientId)
}

// GetClientByType is a free data retrieval call binding the contract method 0x40d20d13.
//
// Solidity: function getClientByType(string clientType) view returns(address)
func (_Ibchandler *IbchandlerCaller) GetClientByType(opts *bind.CallOpts, clientType string) (common.Address, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getClientByType", clientType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClientByType is a free data retrieval call binding the contract method 0x40d20d13.
//
// Solidity: function getClientByType(string clientType) view returns(address)
func (_Ibchandler *IbchandlerSession) GetClientByType(clientType string) (common.Address, error) {
	return _Ibchandler.Contract.GetClientByType(&_Ibchandler.CallOpts, clientType)
}

// GetClientByType is a free data retrieval call binding the contract method 0x40d20d13.
//
// Solidity: function getClientByType(string clientType) view returns(address)
func (_Ibchandler *IbchandlerCallerSession) GetClientByType(clientType string) (common.Address, error) {
	return _Ibchandler.Contract.GetClientByType(&_Ibchandler.CallOpts, clientType)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes, bool)
func (_Ibchandler *IbchandlerCaller) GetClientState(opts *bind.CallOpts, clientId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getClientState", clientId)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes, bool)
func (_Ibchandler *IbchandlerSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Ibchandler.Contract.GetClientState(&_Ibchandler.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns(bytes, bool)
func (_Ibchandler *IbchandlerCallerSession) GetClientState(clientId string) ([]byte, bool, error) {
	return _Ibchandler.Contract.GetClientState(&_Ibchandler.CallOpts, clientId)
}

// GetClientType is a free data retrieval call binding the contract method 0x84515f5d.
//
// Solidity: function getClientType(string clientId) view returns(string)
func (_Ibchandler *IbchandlerCaller) GetClientType(opts *bind.CallOpts, clientId string) (string, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getClientType", clientId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetClientType is a free data retrieval call binding the contract method 0x84515f5d.
//
// Solidity: function getClientType(string clientId) view returns(string)
func (_Ibchandler *IbchandlerSession) GetClientType(clientId string) (string, error) {
	return _Ibchandler.Contract.GetClientType(&_Ibchandler.CallOpts, clientId)
}

// GetClientType is a free data retrieval call binding the contract method 0x84515f5d.
//
// Solidity: function getClientType(string clientId) view returns(string)
func (_Ibchandler *IbchandlerCallerSession) GetClientType(clientId string) (string, error) {
	return _Ibchandler.Contract.GetClientType(&_Ibchandler.CallOpts, clientId)
}

// GetCommitment is a free data retrieval call binding the contract method 0x7795820c.
//
// Solidity: function getCommitment(bytes32 hashedPath) view returns(bytes32)
func (_Ibchandler *IbchandlerCaller) GetCommitment(opts *bind.CallOpts, hashedPath [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getCommitment", hashedPath)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0x7795820c.
//
// Solidity: function getCommitment(bytes32 hashedPath) view returns(bytes32)
func (_Ibchandler *IbchandlerSession) GetCommitment(hashedPath [32]byte) ([32]byte, error) {
	return _Ibchandler.Contract.GetCommitment(&_Ibchandler.CallOpts, hashedPath)
}

// GetCommitment is a free data retrieval call binding the contract method 0x7795820c.
//
// Solidity: function getCommitment(bytes32 hashedPath) view returns(bytes32)
func (_Ibchandler *IbchandlerCallerSession) GetCommitment(hashedPath [32]byte) ([32]byte, error) {
	return _Ibchandler.Contract.GetCommitment(&_Ibchandler.CallOpts, hashedPath)
}

// GetCommitmentPrefix is a free data retrieval call binding the contract method 0xb5ad7134.
//
// Solidity: function getCommitmentPrefix() view returns(bytes)
func (_Ibchandler *IbchandlerCaller) GetCommitmentPrefix(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getCommitmentPrefix")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCommitmentPrefix is a free data retrieval call binding the contract method 0xb5ad7134.
//
// Solidity: function getCommitmentPrefix() view returns(bytes)
func (_Ibchandler *IbchandlerSession) GetCommitmentPrefix() ([]byte, error) {
	return _Ibchandler.Contract.GetCommitmentPrefix(&_Ibchandler.CallOpts)
}

// GetCommitmentPrefix is a free data retrieval call binding the contract method 0xb5ad7134.
//
// Solidity: function getCommitmentPrefix() view returns(bytes)
func (_Ibchandler *IbchandlerCallerSession) GetCommitmentPrefix() ([]byte, error) {
	return _Ibchandler.Contract.GetCommitmentPrefix(&_Ibchandler.CallOpts)
}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,(string,string,(bytes)),uint64), bool)
func (_Ibchandler *IbchandlerCaller) GetConnection(opts *bind.CallOpts, connectionId string) (ConnectionEndData, bool, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getConnection", connectionId)

	if err != nil {
		return *new(ConnectionEndData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ConnectionEndData)).(*ConnectionEndData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,(string,string,(bytes)),uint64), bool)
func (_Ibchandler *IbchandlerSession) GetConnection(connectionId string) (ConnectionEndData, bool, error) {
	return _Ibchandler.Contract.GetConnection(&_Ibchandler.CallOpts, connectionId)
}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,(string,string,(bytes)),uint64), bool)
func (_Ibchandler *IbchandlerCallerSession) GetConnection(connectionId string) (ConnectionEndData, bool, error) {
	return _Ibchandler.Contract.GetConnection(&_Ibchandler.CallOpts, connectionId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Ibchandler *IbchandlerCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height HeightData) ([]byte, bool, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getConsensusState", clientId, height)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Ibchandler *IbchandlerSession) GetConsensusState(clientId string, height HeightData) ([]byte, bool, error) {
	return _Ibchandler.Contract.GetConsensusState(&_Ibchandler.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0x6cf44bf4.
//
// Solidity: function getConsensusState(string clientId, (uint64,uint64) height) view returns(bytes consensusStateBytes, bool)
func (_Ibchandler *IbchandlerCallerSession) GetConsensusState(clientId string, height HeightData) ([]byte, bool, error) {
	return _Ibchandler.Contract.GetConsensusState(&_Ibchandler.CallOpts, clientId, height)
}

// GetExpectedTimePerBlock is a free data retrieval call binding the contract method 0xec75d829.
//
// Solidity: function getExpectedTimePerBlock() view returns(uint64)
func (_Ibchandler *IbchandlerCaller) GetExpectedTimePerBlock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getExpectedTimePerBlock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetExpectedTimePerBlock is a free data retrieval call binding the contract method 0xec75d829.
//
// Solidity: function getExpectedTimePerBlock() view returns(uint64)
func (_Ibchandler *IbchandlerSession) GetExpectedTimePerBlock() (uint64, error) {
	return _Ibchandler.Contract.GetExpectedTimePerBlock(&_Ibchandler.CallOpts)
}

// GetExpectedTimePerBlock is a free data retrieval call binding the contract method 0xec75d829.
//
// Solidity: function getExpectedTimePerBlock() view returns(uint64)
func (_Ibchandler *IbchandlerCallerSession) GetExpectedTimePerBlock() (uint64, error) {
	return _Ibchandler.Contract.GetExpectedTimePerBlock(&_Ibchandler.CallOpts)
}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerCaller) GetNextSequenceAck(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getNextSequenceAck", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerSession) GetNextSequenceAck(portId string, channelId string) (uint64, error) {
	return _Ibchandler.Contract.GetNextSequenceAck(&_Ibchandler.CallOpts, portId, channelId)
}

// GetNextSequenceAck is a free data retrieval call binding the contract method 0x4e08c6f3.
//
// Solidity: function getNextSequenceAck(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerCallerSession) GetNextSequenceAck(portId string, channelId string) (uint64, error) {
	return _Ibchandler.Contract.GetNextSequenceAck(&_Ibchandler.CallOpts, portId, channelId)
}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerCaller) GetNextSequenceRecv(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getNextSequenceRecv", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerSession) GetNextSequenceRecv(portId string, channelId string) (uint64, error) {
	return _Ibchandler.Contract.GetNextSequenceRecv(&_Ibchandler.CallOpts, portId, channelId)
}

// GetNextSequenceRecv is a free data retrieval call binding the contract method 0xe211bb06.
//
// Solidity: function getNextSequenceRecv(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerCallerSession) GetNextSequenceRecv(portId string, channelId string) (uint64, error) {
	return _Ibchandler.Contract.GetNextSequenceRecv(&_Ibchandler.CallOpts, portId, channelId)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerCaller) GetNextSequenceSend(opts *bind.CallOpts, portId string, channelId string) (uint64, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getNextSequenceSend", portId, channelId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerSession) GetNextSequenceSend(portId string, channelId string) (uint64, error) {
	return _Ibchandler.Contract.GetNextSequenceSend(&_Ibchandler.CallOpts, portId, channelId)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string portId, string channelId) view returns(uint64)
func (_Ibchandler *IbchandlerCallerSession) GetNextSequenceSend(portId string, channelId string) (uint64, error) {
	return _Ibchandler.Contract.GetNextSequenceSend(&_Ibchandler.CallOpts, portId, channelId)
}

// GetPacketReceipt is a free data retrieval call binding the contract method 0xb5226815.
//
// Solidity: function getPacketReceipt(string portId, string channelId, uint64 sequence) view returns(uint8)
func (_Ibchandler *IbchandlerCaller) GetPacketReceipt(opts *bind.CallOpts, portId string, channelId string, sequence uint64) (uint8, error) {
	var out []interface{}
	err := _Ibchandler.contract.Call(opts, &out, "getPacketReceipt", portId, channelId, sequence)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPacketReceipt is a free data retrieval call binding the contract method 0xb5226815.
//
// Solidity: function getPacketReceipt(string portId, string channelId, uint64 sequence) view returns(uint8)
func (_Ibchandler *IbchandlerSession) GetPacketReceipt(portId string, channelId string, sequence uint64) (uint8, error) {
	return _Ibchandler.Contract.GetPacketReceipt(&_Ibchandler.CallOpts, portId, channelId, sequence)
}

// GetPacketReceipt is a free data retrieval call binding the contract method 0xb5226815.
//
// Solidity: function getPacketReceipt(string portId, string channelId, uint64 sequence) view returns(uint8)
func (_Ibchandler *IbchandlerCallerSession) GetPacketReceipt(portId string, channelId string, sequence uint64) (uint8, error) {
	return _Ibchandler.Contract.GetPacketReceipt(&_Ibchandler.CallOpts, portId, channelId, sequence)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0x59f37976.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactor) AcknowledgePacket(opts *bind.TransactOpts, arg0 IIBCChannelAcknowledgePacketMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "acknowledgePacket", arg0)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0x59f37976.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerSession) AcknowledgePacket(arg0 IIBCChannelAcknowledgePacketMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.Contract.AcknowledgePacket(&_Ibchandler.TransactOpts, arg0)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0x59f37976.
//
// Solidity: function acknowledgePacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) AcknowledgePacket(arg0 IIBCChannelAcknowledgePacketMsgPacketAcknowledgement) (*types.Transaction, error) {
	return _Ibchandler.Contract.AcknowledgePacket(&_Ibchandler.TransactOpts, arg0)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibchandler *IbchandlerTransactor) BindPort(opts *bind.TransactOpts, portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "bindPort", portId, moduleAddress)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibchandler *IbchandlerSession) BindPort(portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibchandler.Contract.BindPort(&_Ibchandler.TransactOpts, portId, moduleAddress)
}

// BindPort is a paid mutator transaction binding the contract method 0x117e886a.
//
// Solidity: function bindPort(string portId, address moduleAddress) returns()
func (_Ibchandler *IbchandlerTransactorSession) BindPort(portId string, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ibchandler.Contract.BindPort(&_Ibchandler.TransactOpts, portId, moduleAddress)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0x25cbc3a6.
//
// Solidity: function channelCloseConfirm((string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactor) ChannelCloseConfirm(opts *bind.TransactOpts, arg0 IIBCChannelHandshakeMsgChannelCloseConfirm) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelCloseConfirm", arg0)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0x25cbc3a6.
//
// Solidity: function channelCloseConfirm((string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerSession) ChannelCloseConfirm(arg0 IIBCChannelHandshakeMsgChannelCloseConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelCloseConfirm(&_Ibchandler.TransactOpts, arg0)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0x25cbc3a6.
//
// Solidity: function channelCloseConfirm((string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) ChannelCloseConfirm(arg0 IIBCChannelHandshakeMsgChannelCloseConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelCloseConfirm(&_Ibchandler.TransactOpts, arg0)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0xa06cb3a2.
//
// Solidity: function channelCloseInit((string,string) ) returns()
func (_Ibchandler *IbchandlerTransactor) ChannelCloseInit(opts *bind.TransactOpts, arg0 IIBCChannelHandshakeMsgChannelCloseInit) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelCloseInit", arg0)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0xa06cb3a2.
//
// Solidity: function channelCloseInit((string,string) ) returns()
func (_Ibchandler *IbchandlerSession) ChannelCloseInit(arg0 IIBCChannelHandshakeMsgChannelCloseInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelCloseInit(&_Ibchandler.TransactOpts, arg0)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0xa06cb3a2.
//
// Solidity: function channelCloseInit((string,string) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) ChannelCloseInit(arg0 IIBCChannelHandshakeMsgChannelCloseInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelCloseInit(&_Ibchandler.TransactOpts, arg0)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x256c4199.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactor) ChannelOpenAck(opts *bind.TransactOpts, arg0 IIBCChannelHandshakeMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelOpenAck", arg0)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x256c4199.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerSession) ChannelOpenAck(arg0 IIBCChannelHandshakeMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenAck(&_Ibchandler.TransactOpts, arg0)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x256c4199.
//
// Solidity: function channelOpenAck((string,string,string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) ChannelOpenAck(arg0 IIBCChannelHandshakeMsgChannelOpenAck) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenAck(&_Ibchandler.TransactOpts, arg0)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x5bd51b62.
//
// Solidity: function channelOpenConfirm((string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactor) ChannelOpenConfirm(opts *bind.TransactOpts, arg0 IIBCChannelHandshakeMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelOpenConfirm", arg0)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x5bd51b62.
//
// Solidity: function channelOpenConfirm((string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerSession) ChannelOpenConfirm(arg0 IIBCChannelHandshakeMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenConfirm(&_Ibchandler.TransactOpts, arg0)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x5bd51b62.
//
// Solidity: function channelOpenConfirm((string,string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) ChannelOpenConfirm(arg0 IIBCChannelHandshakeMsgChannelOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenConfirm(&_Ibchandler.TransactOpts, arg0)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xdd3469fc.
//
// Solidity: function channelOpenInit((string,(uint8,uint8,(string,string),string[],string)) ) returns(string, string)
func (_Ibchandler *IbchandlerTransactor) ChannelOpenInit(opts *bind.TransactOpts, arg0 IIBCChannelHandshakeMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelOpenInit", arg0)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xdd3469fc.
//
// Solidity: function channelOpenInit((string,(uint8,uint8,(string,string),string[],string)) ) returns(string, string)
func (_Ibchandler *IbchandlerSession) ChannelOpenInit(arg0 IIBCChannelHandshakeMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenInit(&_Ibchandler.TransactOpts, arg0)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xdd3469fc.
//
// Solidity: function channelOpenInit((string,(uint8,uint8,(string,string),string[],string)) ) returns(string, string)
func (_Ibchandler *IbchandlerTransactorSession) ChannelOpenInit(arg0 IIBCChannelHandshakeMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenInit(&_Ibchandler.TransactOpts, arg0)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x11b88a15.
//
// Solidity: function channelOpenTry((string,(uint8,uint8,(string,string),string[],string),string,bytes,(uint64,uint64)) ) returns(string, string)
func (_Ibchandler *IbchandlerTransactor) ChannelOpenTry(opts *bind.TransactOpts, arg0 IIBCChannelHandshakeMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "channelOpenTry", arg0)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x11b88a15.
//
// Solidity: function channelOpenTry((string,(uint8,uint8,(string,string),string[],string),string,bytes,(uint64,uint64)) ) returns(string, string)
func (_Ibchandler *IbchandlerSession) ChannelOpenTry(arg0 IIBCChannelHandshakeMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenTry(&_Ibchandler.TransactOpts, arg0)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x11b88a15.
//
// Solidity: function channelOpenTry((string,(uint8,uint8,(string,string),string[],string),string,bytes,(uint64,uint64)) ) returns(string, string)
func (_Ibchandler *IbchandlerTransactorSession) ChannelOpenTry(arg0 IIBCChannelHandshakeMsgChannelOpenTry) (*types.Transaction, error) {
	return _Ibchandler.Contract.ChannelOpenTry(&_Ibchandler.TransactOpts, arg0)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0xb531861f.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,(uint64,uint64),(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactor) ConnectionOpenAck(opts *bind.TransactOpts, arg0 IIBCConnectionMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "connectionOpenAck", arg0)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0xb531861f.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,(uint64,uint64),(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerSession) ConnectionOpenAck(arg0 IIBCConnectionMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenAck(&_Ibchandler.TransactOpts, arg0)
}

// ConnectionOpenAck is a paid mutator transaction binding the contract method 0xb531861f.
//
// Solidity: function connectionOpenAck((string,bytes,(string,string[]),string,bytes,bytes,bytes,(uint64,uint64),(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) ConnectionOpenAck(arg0 IIBCConnectionMsgConnectionOpenAck) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenAck(&_Ibchandler.TransactOpts, arg0)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0x6a728f2c.
//
// Solidity: function connectionOpenConfirm((string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactor) ConnectionOpenConfirm(opts *bind.TransactOpts, arg0 IIBCConnectionMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "connectionOpenConfirm", arg0)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0x6a728f2c.
//
// Solidity: function connectionOpenConfirm((string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerSession) ConnectionOpenConfirm(arg0 IIBCConnectionMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenConfirm(&_Ibchandler.TransactOpts, arg0)
}

// ConnectionOpenConfirm is a paid mutator transaction binding the contract method 0x6a728f2c.
//
// Solidity: function connectionOpenConfirm((string,bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) ConnectionOpenConfirm(arg0 IIBCConnectionMsgConnectionOpenConfirm) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenConfirm(&_Ibchandler.TransactOpts, arg0)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xd3c68ba0.
//
// Solidity: function connectionOpenInit((string,(string,string,(bytes)),(string,string[]),uint64) ) returns(string)
func (_Ibchandler *IbchandlerTransactor) ConnectionOpenInit(opts *bind.TransactOpts, arg0 IIBCConnectionMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "connectionOpenInit", arg0)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xd3c68ba0.
//
// Solidity: function connectionOpenInit((string,(string,string,(bytes)),(string,string[]),uint64) ) returns(string)
func (_Ibchandler *IbchandlerSession) ConnectionOpenInit(arg0 IIBCConnectionMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenInit(&_Ibchandler.TransactOpts, arg0)
}

// ConnectionOpenInit is a paid mutator transaction binding the contract method 0xd3c68ba0.
//
// Solidity: function connectionOpenInit((string,(string,string,(bytes)),(string,string[]),uint64) ) returns(string)
func (_Ibchandler *IbchandlerTransactorSession) ConnectionOpenInit(arg0 IIBCConnectionMsgConnectionOpenInit) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenInit(&_Ibchandler.TransactOpts, arg0)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x04f68e5c.
//
// Solidity: function connectionOpenTry(((string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,(uint64,uint64),(uint64,uint64)) ) returns(string)
func (_Ibchandler *IbchandlerTransactor) ConnectionOpenTry(opts *bind.TransactOpts, arg0 IIBCConnectionMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "connectionOpenTry", arg0)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x04f68e5c.
//
// Solidity: function connectionOpenTry(((string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,(uint64,uint64),(uint64,uint64)) ) returns(string)
func (_Ibchandler *IbchandlerSession) ConnectionOpenTry(arg0 IIBCConnectionMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenTry(&_Ibchandler.TransactOpts, arg0)
}

// ConnectionOpenTry is a paid mutator transaction binding the contract method 0x04f68e5c.
//
// Solidity: function connectionOpenTry(((string,string,(bytes)),uint64,string,bytes,(string,string[])[],bytes,bytes,bytes,(uint64,uint64),(uint64,uint64)) ) returns(string)
func (_Ibchandler *IbchandlerTransactorSession) ConnectionOpenTry(arg0 IIBCConnectionMsgConnectionOpenTry) (*types.Transaction, error) {
	return _Ibchandler.Contract.ConnectionOpenTry(&_Ibchandler.TransactOpts, arg0)
}

// CreateClient is a paid mutator transaction binding the contract method 0xd5a24481.
//
// Solidity: function createClient((string,bytes,bytes) ) returns(string)
func (_Ibchandler *IbchandlerTransactor) CreateClient(opts *bind.TransactOpts, arg0 IIBCClientMsgCreateClient) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "createClient", arg0)
}

// CreateClient is a paid mutator transaction binding the contract method 0xd5a24481.
//
// Solidity: function createClient((string,bytes,bytes) ) returns(string)
func (_Ibchandler *IbchandlerSession) CreateClient(arg0 IIBCClientMsgCreateClient) (*types.Transaction, error) {
	return _Ibchandler.Contract.CreateClient(&_Ibchandler.TransactOpts, arg0)
}

// CreateClient is a paid mutator transaction binding the contract method 0xd5a24481.
//
// Solidity: function createClient((string,bytes,bytes) ) returns(string)
func (_Ibchandler *IbchandlerTransactorSession) CreateClient(arg0 IIBCClientMsgCreateClient) (*types.Transaction, error) {
	return _Ibchandler.Contract.CreateClient(&_Ibchandler.TransactOpts, arg0)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x236ebd70.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactor) RecvPacket(opts *bind.TransactOpts, arg0 IIBCChannelRecvPacketMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "recvPacket", arg0)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x236ebd70.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerSession) RecvPacket(arg0 IIBCChannelRecvPacketMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.Contract.RecvPacket(&_Ibchandler.TransactOpts, arg0)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x236ebd70.
//
// Solidity: function recvPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,(uint64,uint64)) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) RecvPacket(arg0 IIBCChannelRecvPacketMsgPacketRecv) (*types.Transaction, error) {
	return _Ibchandler.Contract.RecvPacket(&_Ibchandler.TransactOpts, arg0)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibchandler *IbchandlerTransactor) RegisterClient(opts *bind.TransactOpts, clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "registerClient", clientType, client)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibchandler *IbchandlerSession) RegisterClient(clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibchandler.Contract.RegisterClient(&_Ibchandler.TransactOpts, clientType, client)
}

// RegisterClient is a paid mutator transaction binding the contract method 0x18c19870.
//
// Solidity: function registerClient(string clientType, address client) returns()
func (_Ibchandler *IbchandlerTransactorSession) RegisterClient(clientType string, client common.Address) (*types.Transaction, error) {
	return _Ibchandler.Contract.RegisterClient(&_Ibchandler.TransactOpts, clientType, client)
}

// SendPacket is a paid mutator transaction binding the contract method 0xae4cd201.
//
// Solidity: function sendPacket(string , string , (uint64,uint64) , uint64 , bytes ) returns(uint64)
func (_Ibchandler *IbchandlerTransactor) SendPacket(opts *bind.TransactOpts, arg0 string, arg1 string, arg2 HeightData, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "sendPacket", arg0, arg1, arg2, arg3, arg4)
}

// SendPacket is a paid mutator transaction binding the contract method 0xae4cd201.
//
// Solidity: function sendPacket(string , string , (uint64,uint64) , uint64 , bytes ) returns(uint64)
func (_Ibchandler *IbchandlerSession) SendPacket(arg0 string, arg1 string, arg2 HeightData, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _Ibchandler.Contract.SendPacket(&_Ibchandler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendPacket is a paid mutator transaction binding the contract method 0xae4cd201.
//
// Solidity: function sendPacket(string , string , (uint64,uint64) , uint64 , bytes ) returns(uint64)
func (_Ibchandler *IbchandlerTransactorSession) SendPacket(arg0 string, arg1 string, arg2 HeightData, arg3 uint64, arg4 []byte) (*types.Transaction, error) {
	return _Ibchandler.Contract.SendPacket(&_Ibchandler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SetExpectedTimePerBlock is a paid mutator transaction binding the contract method 0x27184c13.
//
// Solidity: function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) returns()
func (_Ibchandler *IbchandlerTransactor) SetExpectedTimePerBlock(opts *bind.TransactOpts, expectedTimePerBlock_ uint64) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "setExpectedTimePerBlock", expectedTimePerBlock_)
}

// SetExpectedTimePerBlock is a paid mutator transaction binding the contract method 0x27184c13.
//
// Solidity: function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) returns()
func (_Ibchandler *IbchandlerSession) SetExpectedTimePerBlock(expectedTimePerBlock_ uint64) (*types.Transaction, error) {
	return _Ibchandler.Contract.SetExpectedTimePerBlock(&_Ibchandler.TransactOpts, expectedTimePerBlock_)
}

// SetExpectedTimePerBlock is a paid mutator transaction binding the contract method 0x27184c13.
//
// Solidity: function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) returns()
func (_Ibchandler *IbchandlerTransactorSession) SetExpectedTimePerBlock(expectedTimePerBlock_ uint64) (*types.Transaction, error) {
	return _Ibchandler.Contract.SetExpectedTimePerBlock(&_Ibchandler.TransactOpts, expectedTimePerBlock_)
}

// TimeoutOnClose is a paid mutator transaction binding the contract method 0x9ebb2107.
//
// Solidity: function timeoutOnClose(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,(uint64,uint64),uint64) ) returns()
func (_Ibchandler *IbchandlerTransactor) TimeoutOnClose(opts *bind.TransactOpts, arg0 IIBCChannelPacketTimeoutMsgTimeoutOnClose) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "timeoutOnClose", arg0)
}

// TimeoutOnClose is a paid mutator transaction binding the contract method 0x9ebb2107.
//
// Solidity: function timeoutOnClose(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,(uint64,uint64),uint64) ) returns()
func (_Ibchandler *IbchandlerSession) TimeoutOnClose(arg0 IIBCChannelPacketTimeoutMsgTimeoutOnClose) (*types.Transaction, error) {
	return _Ibchandler.Contract.TimeoutOnClose(&_Ibchandler.TransactOpts, arg0)
}

// TimeoutOnClose is a paid mutator transaction binding the contract method 0x9ebb2107.
//
// Solidity: function timeoutOnClose(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,bytes,(uint64,uint64),uint64) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) TimeoutOnClose(arg0 IIBCChannelPacketTimeoutMsgTimeoutOnClose) (*types.Transaction, error) {
	return _Ibchandler.Contract.TimeoutOnClose(&_Ibchandler.TransactOpts, arg0)
}

// TimeoutPacket is a paid mutator transaction binding the contract method 0xaa18c8b1.
//
// Solidity: function timeoutPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,(uint64,uint64),uint64) ) returns()
func (_Ibchandler *IbchandlerTransactor) TimeoutPacket(opts *bind.TransactOpts, arg0 IIBCChannelPacketTimeoutMsgTimeoutPacket) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "timeoutPacket", arg0)
}

// TimeoutPacket is a paid mutator transaction binding the contract method 0xaa18c8b1.
//
// Solidity: function timeoutPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,(uint64,uint64),uint64) ) returns()
func (_Ibchandler *IbchandlerSession) TimeoutPacket(arg0 IIBCChannelPacketTimeoutMsgTimeoutPacket) (*types.Transaction, error) {
	return _Ibchandler.Contract.TimeoutPacket(&_Ibchandler.TransactOpts, arg0)
}

// TimeoutPacket is a paid mutator transaction binding the contract method 0xaa18c8b1.
//
// Solidity: function timeoutPacket(((uint64,string,string,string,string,bytes,(uint64,uint64),uint64),bytes,(uint64,uint64),uint64) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) TimeoutPacket(arg0 IIBCChannelPacketTimeoutMsgTimeoutPacket) (*types.Transaction, error) {
	return _Ibchandler.Contract.TimeoutPacket(&_Ibchandler.TransactOpts, arg0)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) ) returns()
func (_Ibchandler *IbchandlerTransactor) UpdateClient(opts *bind.TransactOpts, arg0 IIBCClientMsgUpdateClient) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "updateClient", arg0)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) ) returns()
func (_Ibchandler *IbchandlerSession) UpdateClient(arg0 IIBCClientMsgUpdateClient) (*types.Transaction, error) {
	return _Ibchandler.Contract.UpdateClient(&_Ibchandler.TransactOpts, arg0)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xda6cea55.
//
// Solidity: function updateClient((string,bytes) ) returns()
func (_Ibchandler *IbchandlerTransactorSession) UpdateClient(arg0 IIBCClientMsgUpdateClient) (*types.Transaction, error) {
	return _Ibchandler.Contract.UpdateClient(&_Ibchandler.TransactOpts, arg0)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0xb56e79de.
//
// Solidity: function writeAcknowledgement(string , string , uint64 , bytes ) returns()
func (_Ibchandler *IbchandlerTransactor) WriteAcknowledgement(opts *bind.TransactOpts, arg0 string, arg1 string, arg2 uint64, arg3 []byte) (*types.Transaction, error) {
	return _Ibchandler.contract.Transact(opts, "writeAcknowledgement", arg0, arg1, arg2, arg3)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0xb56e79de.
//
// Solidity: function writeAcknowledgement(string , string , uint64 , bytes ) returns()
func (_Ibchandler *IbchandlerSession) WriteAcknowledgement(arg0 string, arg1 string, arg2 uint64, arg3 []byte) (*types.Transaction, error) {
	return _Ibchandler.Contract.WriteAcknowledgement(&_Ibchandler.TransactOpts, arg0, arg1, arg2, arg3)
}

// WriteAcknowledgement is a paid mutator transaction binding the contract method 0xb56e79de.
//
// Solidity: function writeAcknowledgement(string , string , uint64 , bytes ) returns()
func (_Ibchandler *IbchandlerTransactorSession) WriteAcknowledgement(arg0 string, arg1 string, arg2 uint64, arg3 []byte) (*types.Transaction, error) {
	return _Ibchandler.Contract.WriteAcknowledgement(&_Ibchandler.TransactOpts, arg0, arg1, arg2, arg3)
}

// IbchandlerAcknowledgePacketIterator is returned from FilterAcknowledgePacket and is used to iterate over the raw logs and unpacked data for AcknowledgePacket events raised by the Ibchandler contract.
type IbchandlerAcknowledgePacketIterator struct {
	Event *IbchandlerAcknowledgePacket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbchandlerAcknowledgePacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerAcknowledgePacket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbchandlerAcknowledgePacket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbchandlerAcknowledgePacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerAcknowledgePacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerAcknowledgePacket represents a AcknowledgePacket event raised by the Ibchandler contract.
type IbchandlerAcknowledgePacket struct {
	Packet          PacketData
	Acknowledgement []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgePacket is a free log retrieval operation binding the contract event 0x47471450765e6e1b0b055ba2a1de04d4ce71f778c92b306e725083eb120dfd89.
//
// Solidity: event AcknowledgePacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) FilterAcknowledgePacket(opts *bind.FilterOpts) (*IbchandlerAcknowledgePacketIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "AcknowledgePacket")
	if err != nil {
		return nil, err
	}
	return &IbchandlerAcknowledgePacketIterator{contract: _Ibchandler.contract, event: "AcknowledgePacket", logs: logs, sub: sub}, nil
}

// WatchAcknowledgePacket is a free log subscription operation binding the contract event 0x47471450765e6e1b0b055ba2a1de04d4ce71f778c92b306e725083eb120dfd89.
//
// Solidity: event AcknowledgePacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) WatchAcknowledgePacket(opts *bind.WatchOpts, sink chan<- *IbchandlerAcknowledgePacket) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "AcknowledgePacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerAcknowledgePacket)
				if err := _Ibchandler.contract.UnpackLog(event, "AcknowledgePacket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAcknowledgePacket is a log parse operation binding the contract event 0x47471450765e6e1b0b055ba2a1de04d4ce71f778c92b306e725083eb120dfd89.
//
// Solidity: event AcknowledgePacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) ParseAcknowledgePacket(log types.Log) (*IbchandlerAcknowledgePacket, error) {
	event := new(IbchandlerAcknowledgePacket)
	if err := _Ibchandler.contract.UnpackLog(event, "AcknowledgePacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerGeneratedChannelIdentifierIterator is returned from FilterGeneratedChannelIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedChannelIdentifier events raised by the Ibchandler contract.
type IbchandlerGeneratedChannelIdentifierIterator struct {
	Event *IbchandlerGeneratedChannelIdentifier // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbchandlerGeneratedChannelIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerGeneratedChannelIdentifier)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbchandlerGeneratedChannelIdentifier)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbchandlerGeneratedChannelIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerGeneratedChannelIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerGeneratedChannelIdentifier represents a GeneratedChannelIdentifier event raised by the Ibchandler contract.
type IbchandlerGeneratedChannelIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedChannelIdentifier is a free log retrieval operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) FilterGeneratedChannelIdentifier(opts *bind.FilterOpts) (*IbchandlerGeneratedChannelIdentifierIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "GeneratedChannelIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchandlerGeneratedChannelIdentifierIterator{contract: _Ibchandler.contract, event: "GeneratedChannelIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedChannelIdentifier is a free log subscription operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) WatchGeneratedChannelIdentifier(opts *bind.WatchOpts, sink chan<- *IbchandlerGeneratedChannelIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "GeneratedChannelIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerGeneratedChannelIdentifier)
				if err := _Ibchandler.contract.UnpackLog(event, "GeneratedChannelIdentifier", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGeneratedChannelIdentifier is a log parse operation binding the contract event 0x01fb9b8778b6fb840b058bb971dea3ba81c167b010a0216afe600826884f9ba7.
//
// Solidity: event GeneratedChannelIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) ParseGeneratedChannelIdentifier(log types.Log) (*IbchandlerGeneratedChannelIdentifier, error) {
	event := new(IbchandlerGeneratedChannelIdentifier)
	if err := _Ibchandler.contract.UnpackLog(event, "GeneratedChannelIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerGeneratedClientIdentifierIterator is returned from FilterGeneratedClientIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedClientIdentifier events raised by the Ibchandler contract.
type IbchandlerGeneratedClientIdentifierIterator struct {
	Event *IbchandlerGeneratedClientIdentifier // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbchandlerGeneratedClientIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerGeneratedClientIdentifier)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbchandlerGeneratedClientIdentifier)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbchandlerGeneratedClientIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerGeneratedClientIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerGeneratedClientIdentifier represents a GeneratedClientIdentifier event raised by the Ibchandler contract.
type IbchandlerGeneratedClientIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedClientIdentifier is a free log retrieval operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) FilterGeneratedClientIdentifier(opts *bind.FilterOpts) (*IbchandlerGeneratedClientIdentifierIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "GeneratedClientIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchandlerGeneratedClientIdentifierIterator{contract: _Ibchandler.contract, event: "GeneratedClientIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedClientIdentifier is a free log subscription operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) WatchGeneratedClientIdentifier(opts *bind.WatchOpts, sink chan<- *IbchandlerGeneratedClientIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "GeneratedClientIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerGeneratedClientIdentifier)
				if err := _Ibchandler.contract.UnpackLog(event, "GeneratedClientIdentifier", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGeneratedClientIdentifier is a log parse operation binding the contract event 0x601bfcc455d5d4d7738f8c6ac232e0d7cc9c31dab811f1d87c100af0b7fc3a20.
//
// Solidity: event GeneratedClientIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) ParseGeneratedClientIdentifier(log types.Log) (*IbchandlerGeneratedClientIdentifier, error) {
	event := new(IbchandlerGeneratedClientIdentifier)
	if err := _Ibchandler.contract.UnpackLog(event, "GeneratedClientIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerGeneratedConnectionIdentifierIterator is returned from FilterGeneratedConnectionIdentifier and is used to iterate over the raw logs and unpacked data for GeneratedConnectionIdentifier events raised by the Ibchandler contract.
type IbchandlerGeneratedConnectionIdentifierIterator struct {
	Event *IbchandlerGeneratedConnectionIdentifier // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbchandlerGeneratedConnectionIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerGeneratedConnectionIdentifier)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbchandlerGeneratedConnectionIdentifier)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbchandlerGeneratedConnectionIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerGeneratedConnectionIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerGeneratedConnectionIdentifier represents a GeneratedConnectionIdentifier event raised by the Ibchandler contract.
type IbchandlerGeneratedConnectionIdentifier struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGeneratedConnectionIdentifier is a free log retrieval operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) FilterGeneratedConnectionIdentifier(opts *bind.FilterOpts) (*IbchandlerGeneratedConnectionIdentifierIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "GeneratedConnectionIdentifier")
	if err != nil {
		return nil, err
	}
	return &IbchandlerGeneratedConnectionIdentifierIterator{contract: _Ibchandler.contract, event: "GeneratedConnectionIdentifier", logs: logs, sub: sub}, nil
}

// WatchGeneratedConnectionIdentifier is a free log subscription operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) WatchGeneratedConnectionIdentifier(opts *bind.WatchOpts, sink chan<- *IbchandlerGeneratedConnectionIdentifier) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "GeneratedConnectionIdentifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerGeneratedConnectionIdentifier)
				if err := _Ibchandler.contract.UnpackLog(event, "GeneratedConnectionIdentifier", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGeneratedConnectionIdentifier is a log parse operation binding the contract event 0xbcf8ae1e9272e040280c9adfc8033bb831043a9959e37ef4af1f7e8ded16321b.
//
// Solidity: event GeneratedConnectionIdentifier(string arg0)
func (_Ibchandler *IbchandlerFilterer) ParseGeneratedConnectionIdentifier(log types.Log) (*IbchandlerGeneratedConnectionIdentifier, error) {
	event := new(IbchandlerGeneratedConnectionIdentifier)
	if err := _Ibchandler.contract.UnpackLog(event, "GeneratedConnectionIdentifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerRecvPacketIterator is returned from FilterRecvPacket and is used to iterate over the raw logs and unpacked data for RecvPacket events raised by the Ibchandler contract.
type IbchandlerRecvPacketIterator struct {
	Event *IbchandlerRecvPacket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbchandlerRecvPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerRecvPacket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbchandlerRecvPacket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbchandlerRecvPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerRecvPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerRecvPacket represents a RecvPacket event raised by the Ibchandler contract.
type IbchandlerRecvPacket struct {
	Packet PacketData
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecvPacket is a free log retrieval operation binding the contract event 0x346f4351ee865d86a679d00f3995f0520f803d3a227604af08430e26e9345a7a.
//
// Solidity: event RecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) FilterRecvPacket(opts *bind.FilterOpts) (*IbchandlerRecvPacketIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "RecvPacket")
	if err != nil {
		return nil, err
	}
	return &IbchandlerRecvPacketIterator{contract: _Ibchandler.contract, event: "RecvPacket", logs: logs, sub: sub}, nil
}

// WatchRecvPacket is a free log subscription operation binding the contract event 0x346f4351ee865d86a679d00f3995f0520f803d3a227604af08430e26e9345a7a.
//
// Solidity: event RecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) WatchRecvPacket(opts *bind.WatchOpts, sink chan<- *IbchandlerRecvPacket) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "RecvPacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerRecvPacket)
				if err := _Ibchandler.contract.UnpackLog(event, "RecvPacket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRecvPacket is a log parse operation binding the contract event 0x346f4351ee865d86a679d00f3995f0520f803d3a227604af08430e26e9345a7a.
//
// Solidity: event RecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) ParseRecvPacket(log types.Log) (*IbchandlerRecvPacket, error) {
	event := new(IbchandlerRecvPacket)
	if err := _Ibchandler.contract.UnpackLog(event, "RecvPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerSendPacketIterator is returned from FilterSendPacket and is used to iterate over the raw logs and unpacked data for SendPacket events raised by the Ibchandler contract.
type IbchandlerSendPacketIterator struct {
	Event *IbchandlerSendPacket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbchandlerSendPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerSendPacket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbchandlerSendPacket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbchandlerSendPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerSendPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerSendPacket represents a SendPacket event raised by the Ibchandler contract.
type IbchandlerSendPacket struct {
	Sequence         uint64
	SourcePort       string
	SourceChannel    string
	TimeoutHeight    HeightData
	TimeoutTimestamp uint64
	Data             []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSendPacket is a free log retrieval operation binding the contract event 0x2a89ca0e962a61b8115575da63f54bb249cf0137947fc9ab016ac9df88aa347e.
//
// Solidity: event SendPacket(uint64 sequence, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp, bytes data)
func (_Ibchandler *IbchandlerFilterer) FilterSendPacket(opts *bind.FilterOpts) (*IbchandlerSendPacketIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "SendPacket")
	if err != nil {
		return nil, err
	}
	return &IbchandlerSendPacketIterator{contract: _Ibchandler.contract, event: "SendPacket", logs: logs, sub: sub}, nil
}

// WatchSendPacket is a free log subscription operation binding the contract event 0x2a89ca0e962a61b8115575da63f54bb249cf0137947fc9ab016ac9df88aa347e.
//
// Solidity: event SendPacket(uint64 sequence, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp, bytes data)
func (_Ibchandler *IbchandlerFilterer) WatchSendPacket(opts *bind.WatchOpts, sink chan<- *IbchandlerSendPacket) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "SendPacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerSendPacket)
				if err := _Ibchandler.contract.UnpackLog(event, "SendPacket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendPacket is a log parse operation binding the contract event 0x2a89ca0e962a61b8115575da63f54bb249cf0137947fc9ab016ac9df88aa347e.
//
// Solidity: event SendPacket(uint64 sequence, string sourcePort, string sourceChannel, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp, bytes data)
func (_Ibchandler *IbchandlerFilterer) ParseSendPacket(log types.Log) (*IbchandlerSendPacket, error) {
	event := new(IbchandlerSendPacket)
	if err := _Ibchandler.contract.UnpackLog(event, "SendPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerTimeoutPacketIterator is returned from FilterTimeoutPacket and is used to iterate over the raw logs and unpacked data for TimeoutPacket events raised by the Ibchandler contract.
type IbchandlerTimeoutPacketIterator struct {
	Event *IbchandlerTimeoutPacket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbchandlerTimeoutPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerTimeoutPacket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbchandlerTimeoutPacket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbchandlerTimeoutPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerTimeoutPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerTimeoutPacket represents a TimeoutPacket event raised by the Ibchandler contract.
type IbchandlerTimeoutPacket struct {
	Packet PacketData
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTimeoutPacket is a free log retrieval operation binding the contract event 0xa6ccdfd06294bbb481b7b08ab170c1377cccdcaa9e35b2e346a36ee32a1f8f06.
//
// Solidity: event TimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) FilterTimeoutPacket(opts *bind.FilterOpts) (*IbchandlerTimeoutPacketIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "TimeoutPacket")
	if err != nil {
		return nil, err
	}
	return &IbchandlerTimeoutPacketIterator{contract: _Ibchandler.contract, event: "TimeoutPacket", logs: logs, sub: sub}, nil
}

// WatchTimeoutPacket is a free log subscription operation binding the contract event 0xa6ccdfd06294bbb481b7b08ab170c1377cccdcaa9e35b2e346a36ee32a1f8f06.
//
// Solidity: event TimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) WatchTimeoutPacket(opts *bind.WatchOpts, sink chan<- *IbchandlerTimeoutPacket) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "TimeoutPacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerTimeoutPacket)
				if err := _Ibchandler.contract.UnpackLog(event, "TimeoutPacket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimeoutPacket is a log parse operation binding the contract event 0xa6ccdfd06294bbb481b7b08ab170c1377cccdcaa9e35b2e346a36ee32a1f8f06.
//
// Solidity: event TimeoutPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet)
func (_Ibchandler *IbchandlerFilterer) ParseTimeoutPacket(log types.Log) (*IbchandlerTimeoutPacket, error) {
	event := new(IbchandlerTimeoutPacket)
	if err := _Ibchandler.contract.UnpackLog(event, "TimeoutPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbchandlerWriteAcknowledgementIterator is returned from FilterWriteAcknowledgement and is used to iterate over the raw logs and unpacked data for WriteAcknowledgement events raised by the Ibchandler contract.
type IbchandlerWriteAcknowledgementIterator struct {
	Event *IbchandlerWriteAcknowledgement // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IbchandlerWriteAcknowledgementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbchandlerWriteAcknowledgement)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IbchandlerWriteAcknowledgement)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IbchandlerWriteAcknowledgementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbchandlerWriteAcknowledgementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbchandlerWriteAcknowledgement represents a WriteAcknowledgement event raised by the Ibchandler contract.
type IbchandlerWriteAcknowledgement struct {
	DestinationPortId  string
	DestinationChannel string
	Sequence           uint64
	Acknowledgement    []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWriteAcknowledgement is a free log retrieval operation binding the contract event 0x39b14668930c816f244f4073c0fdf459d3dd73ae571b57b3efe8205919472d2a.
//
// Solidity: event WriteAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) FilterWriteAcknowledgement(opts *bind.FilterOpts) (*IbchandlerWriteAcknowledgementIterator, error) {

	logs, sub, err := _Ibchandler.contract.FilterLogs(opts, "WriteAcknowledgement")
	if err != nil {
		return nil, err
	}
	return &IbchandlerWriteAcknowledgementIterator{contract: _Ibchandler.contract, event: "WriteAcknowledgement", logs: logs, sub: sub}, nil
}

// WatchWriteAcknowledgement is a free log subscription operation binding the contract event 0x39b14668930c816f244f4073c0fdf459d3dd73ae571b57b3efe8205919472d2a.
//
// Solidity: event WriteAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) WatchWriteAcknowledgement(opts *bind.WatchOpts, sink chan<- *IbchandlerWriteAcknowledgement) (event.Subscription, error) {

	logs, sub, err := _Ibchandler.contract.WatchLogs(opts, "WriteAcknowledgement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbchandlerWriteAcknowledgement)
				if err := _Ibchandler.contract.UnpackLog(event, "WriteAcknowledgement", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWriteAcknowledgement is a log parse operation binding the contract event 0x39b14668930c816f244f4073c0fdf459d3dd73ae571b57b3efe8205919472d2a.
//
// Solidity: event WriteAcknowledgement(string destinationPortId, string destinationChannel, uint64 sequence, bytes acknowledgement)
func (_Ibchandler *IbchandlerFilterer) ParseWriteAcknowledgement(log types.Log) (*IbchandlerWriteAcknowledgement, error) {
	event := new(IbchandlerWriteAcknowledgement)
	if err := _Ibchandler.contract.UnpackLog(event, "WriteAcknowledgement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
