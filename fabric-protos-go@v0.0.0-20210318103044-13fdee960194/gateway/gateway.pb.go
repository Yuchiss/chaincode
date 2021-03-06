// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway/gateway.proto

package gateway

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/hyperledger/fabric-protos-go/common"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// EndorseRequest contains the details required to obtain sufficient endorsements for a
// transaction to be committed to the ledger.
type EndorseRequest struct {
	// The unique identifier for the transaction.
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Identifier of the channel this request is bound for.
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The signed proposal ready for endorsement.
	ProposedTransaction  *peer.SignedProposal `protobuf:"bytes,3,opt,name=proposed_transaction,json=proposedTransaction,proto3" json:"proposed_transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EndorseRequest) Reset()         { *m = EndorseRequest{} }
func (m *EndorseRequest) String() string { return proto.CompactTextString(m) }
func (*EndorseRequest) ProtoMessage()    {}
func (*EndorseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{0}
}

func (m *EndorseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndorseRequest.Unmarshal(m, b)
}
func (m *EndorseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndorseRequest.Marshal(b, m, deterministic)
}
func (m *EndorseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndorseRequest.Merge(m, src)
}
func (m *EndorseRequest) XXX_Size() int {
	return xxx_messageInfo_EndorseRequest.Size(m)
}
func (m *EndorseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndorseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndorseRequest proto.InternalMessageInfo

func (m *EndorseRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *EndorseRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *EndorseRequest) GetProposedTransaction() *peer.SignedProposal {
	if m != nil {
		return m.ProposedTransaction
	}
	return nil
}

// EndorseResponse returns the result of endorsing a transaction.
type EndorseResponse struct {
	// The response that is returned by the transaction function, as defined
	// in peer/proposal_response.proto
	Result *peer.Response `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	// The unsigned set of transaction responses from the endorsing peers for signing by the client
	// before submitting to ordering service (via gateway).
	PreparedTransaction  *common.Envelope `protobuf:"bytes,2,opt,name=prepared_transaction,json=preparedTransaction,proto3" json:"prepared_transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EndorseResponse) Reset()         { *m = EndorseResponse{} }
func (m *EndorseResponse) String() string { return proto.CompactTextString(m) }
func (*EndorseResponse) ProtoMessage()    {}
func (*EndorseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{1}
}

func (m *EndorseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndorseResponse.Unmarshal(m, b)
}
func (m *EndorseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndorseResponse.Marshal(b, m, deterministic)
}
func (m *EndorseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndorseResponse.Merge(m, src)
}
func (m *EndorseResponse) XXX_Size() int {
	return xxx_messageInfo_EndorseResponse.Size(m)
}
func (m *EndorseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndorseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndorseResponse proto.InternalMessageInfo

func (m *EndorseResponse) GetResult() *peer.Response {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *EndorseResponse) GetPreparedTransaction() *common.Envelope {
	if m != nil {
		return m.PreparedTransaction
	}
	return nil
}

// SubmitRequest contains the details required to submit a transaction (update the ledger).
type SubmitRequest struct {
	// Identifier of the transaction to submit.
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Identifier of the channel this request is bound for.
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The signed set of endorsed transaction responses to submit.
	PreparedTransaction  *common.Envelope `protobuf:"bytes,3,opt,name=prepared_transaction,json=preparedTransaction,proto3" json:"prepared_transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SubmitRequest) Reset()         { *m = SubmitRequest{} }
func (m *SubmitRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitRequest) ProtoMessage()    {}
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{2}
}

func (m *SubmitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitRequest.Unmarshal(m, b)
}
func (m *SubmitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitRequest.Marshal(b, m, deterministic)
}
func (m *SubmitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRequest.Merge(m, src)
}
func (m *SubmitRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitRequest.Size(m)
}
func (m *SubmitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRequest proto.InternalMessageInfo

func (m *SubmitRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *SubmitRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SubmitRequest) GetPreparedTransaction() *common.Envelope {
	if m != nil {
		return m.PreparedTransaction
	}
	return nil
}

// SubmitResponse returns the result of submitting a transaction.
type SubmitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitResponse) Reset()         { *m = SubmitResponse{} }
func (m *SubmitResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitResponse) ProtoMessage()    {}
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{3}
}

func (m *SubmitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitResponse.Unmarshal(m, b)
}
func (m *SubmitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitResponse.Marshal(b, m, deterministic)
}
func (m *SubmitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitResponse.Merge(m, src)
}
func (m *SubmitResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitResponse.Size(m)
}
func (m *SubmitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitResponse proto.InternalMessageInfo

// CommitStatusRequest contains the details required to check whether a transaction has been
// successfully committed.
type CommitStatusRequest struct {
	// Identifier of the transaction to check.
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Identifier of the channel this request is bound for.
	ChannelId            string   `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitStatusRequest) Reset()         { *m = CommitStatusRequest{} }
func (m *CommitStatusRequest) String() string { return proto.CompactTextString(m) }
func (*CommitStatusRequest) ProtoMessage()    {}
func (*CommitStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{4}
}

func (m *CommitStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitStatusRequest.Unmarshal(m, b)
}
func (m *CommitStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitStatusRequest.Marshal(b, m, deterministic)
}
func (m *CommitStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitStatusRequest.Merge(m, src)
}
func (m *CommitStatusRequest) XXX_Size() int {
	return xxx_messageInfo_CommitStatusRequest.Size(m)
}
func (m *CommitStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommitStatusRequest proto.InternalMessageInfo

func (m *CommitStatusRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *CommitStatusRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// CommitStatusResponse returns the result of committing a transaction.
type CommitStatusResponse struct {
	// The result of the transaction commit, as defined in peer/transaction.proto
	Result               peer.TxValidationCode `protobuf:"varint,1,opt,name=result,proto3,enum=protos.TxValidationCode" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CommitStatusResponse) Reset()         { *m = CommitStatusResponse{} }
func (m *CommitStatusResponse) String() string { return proto.CompactTextString(m) }
func (*CommitStatusResponse) ProtoMessage()    {}
func (*CommitStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{5}
}

func (m *CommitStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitStatusResponse.Unmarshal(m, b)
}
func (m *CommitStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitStatusResponse.Marshal(b, m, deterministic)
}
func (m *CommitStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitStatusResponse.Merge(m, src)
}
func (m *CommitStatusResponse) XXX_Size() int {
	return xxx_messageInfo_CommitStatusResponse.Size(m)
}
func (m *CommitStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommitStatusResponse proto.InternalMessageInfo

func (m *CommitStatusResponse) GetResult() peer.TxValidationCode {
	if m != nil {
		return m.Result
	}
	return peer.TxValidationCode_VALID
}

// EvaluateRequest contains the details required to evaluate a transaction (query the ledger).
type EvaluateRequest struct {
	// Identifier of the transaction to evaluate.
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Identifier of the channel this request is bound for.
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The signed proposal ready for evaluation.
	ProposedTransaction  *peer.SignedProposal `protobuf:"bytes,3,opt,name=proposed_transaction,json=proposedTransaction,proto3" json:"proposed_transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EvaluateRequest) Reset()         { *m = EvaluateRequest{} }
func (m *EvaluateRequest) String() string { return proto.CompactTextString(m) }
func (*EvaluateRequest) ProtoMessage()    {}
func (*EvaluateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{6}
}

func (m *EvaluateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateRequest.Unmarshal(m, b)
}
func (m *EvaluateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateRequest.Marshal(b, m, deterministic)
}
func (m *EvaluateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateRequest.Merge(m, src)
}
func (m *EvaluateRequest) XXX_Size() int {
	return xxx_messageInfo_EvaluateRequest.Size(m)
}
func (m *EvaluateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateRequest proto.InternalMessageInfo

func (m *EvaluateRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *EvaluateRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *EvaluateRequest) GetProposedTransaction() *peer.SignedProposal {
	if m != nil {
		return m.ProposedTransaction
	}
	return nil
}

// EvaluateResponse returns the result of evaluating a transaction.
type EvaluateResponse struct {
	// The response that is returned by the transaction function, as defined
	// in peer/proposal_response.proto
	Result               *peer.Response `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EvaluateResponse) Reset()         { *m = EvaluateResponse{} }
func (m *EvaluateResponse) String() string { return proto.CompactTextString(m) }
func (*EvaluateResponse) ProtoMessage()    {}
func (*EvaluateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{7}
}

func (m *EvaluateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateResponse.Unmarshal(m, b)
}
func (m *EvaluateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateResponse.Marshal(b, m, deterministic)
}
func (m *EvaluateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateResponse.Merge(m, src)
}
func (m *EvaluateResponse) XXX_Size() int {
	return xxx_messageInfo_EvaluateResponse.Size(m)
}
func (m *EvaluateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateResponse proto.InternalMessageInfo

func (m *EvaluateResponse) GetResult() *peer.Response {
	if m != nil {
		return m.Result
	}
	return nil
}

// If any of the functions in the Gateway service returns an error, then it will be in the format of
// a google.rpc.Status message. The 'details' field of this message will be populated with extra
// information if the error is a result of one or more failed requests to remote peers or orderer nodes.
// EndpointError contains details of errors that are received by any of the endorsing peers
// as a result of processing the Evaluate or Endorse services, or from the ordering node(s) as a result of
// processing the Submit service.
type EndpointError struct {
	// The address of the endorsing peer or ordering node that returned an error.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The MSP Identifier of this endpoint.
	MspId string `protobuf:"bytes,2,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	// The error message returned by this endpoint.
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointError) Reset()         { *m = EndpointError{} }
func (m *EndpointError) String() string { return proto.CompactTextString(m) }
func (*EndpointError) ProtoMessage()    {}
func (*EndpointError) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{8}
}

func (m *EndpointError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointError.Unmarshal(m, b)
}
func (m *EndpointError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointError.Marshal(b, m, deterministic)
}
func (m *EndpointError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointError.Merge(m, src)
}
func (m *EndpointError) XXX_Size() int {
	return xxx_messageInfo_EndpointError.Size(m)
}
func (m *EndpointError) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointError.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointError proto.InternalMessageInfo

func (m *EndpointError) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EndpointError) GetMspId() string {
	if m != nil {
		return m.MspId
	}
	return ""
}

func (m *EndpointError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// ProposedTransaction contains the details required for offline signing prior to evaluating or endorsing
// a transaction.
type ProposedTransaction struct {
	// Identifier of the proposed transaction.
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// The signed proposal.
	Proposal             *peer.SignedProposal `protobuf:"bytes,2,opt,name=proposal,proto3" json:"proposal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProposedTransaction) Reset()         { *m = ProposedTransaction{} }
func (m *ProposedTransaction) String() string { return proto.CompactTextString(m) }
func (*ProposedTransaction) ProtoMessage()    {}
func (*ProposedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{9}
}

func (m *ProposedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposedTransaction.Unmarshal(m, b)
}
func (m *ProposedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposedTransaction.Marshal(b, m, deterministic)
}
func (m *ProposedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedTransaction.Merge(m, src)
}
func (m *ProposedTransaction) XXX_Size() int {
	return xxx_messageInfo_ProposedTransaction.Size(m)
}
func (m *ProposedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedTransaction proto.InternalMessageInfo

func (m *ProposedTransaction) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *ProposedTransaction) GetProposal() *peer.SignedProposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

// PreparedTransaction contains the details required for offline signing prior to submitting a transaction.
type PreparedTransaction struct {
	// Identifier of the prepared transaction.
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// The transaction envelope.
	Envelope *common.Envelope `protobuf:"bytes,2,opt,name=envelope,proto3" json:"envelope,omitempty"`
	// The response that is returned by the transaction function during endorsement, as defined
	// in peer/proposal_response.proto
	Result               *peer.Response `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PreparedTransaction) Reset()         { *m = PreparedTransaction{} }
func (m *PreparedTransaction) String() string { return proto.CompactTextString(m) }
func (*PreparedTransaction) ProtoMessage()    {}
func (*PreparedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_285396c8df15061f, []int{10}
}

func (m *PreparedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreparedTransaction.Unmarshal(m, b)
}
func (m *PreparedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreparedTransaction.Marshal(b, m, deterministic)
}
func (m *PreparedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreparedTransaction.Merge(m, src)
}
func (m *PreparedTransaction) XXX_Size() int {
	return xxx_messageInfo_PreparedTransaction.Size(m)
}
func (m *PreparedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_PreparedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_PreparedTransaction proto.InternalMessageInfo

func (m *PreparedTransaction) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *PreparedTransaction) GetEnvelope() *common.Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

func (m *PreparedTransaction) GetResult() *peer.Response {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*EndorseRequest)(nil), "gateway.EndorseRequest")
	proto.RegisterType((*EndorseResponse)(nil), "gateway.EndorseResponse")
	proto.RegisterType((*SubmitRequest)(nil), "gateway.SubmitRequest")
	proto.RegisterType((*SubmitResponse)(nil), "gateway.SubmitResponse")
	proto.RegisterType((*CommitStatusRequest)(nil), "gateway.CommitStatusRequest")
	proto.RegisterType((*CommitStatusResponse)(nil), "gateway.CommitStatusResponse")
	proto.RegisterType((*EvaluateRequest)(nil), "gateway.EvaluateRequest")
	proto.RegisterType((*EvaluateResponse)(nil), "gateway.EvaluateResponse")
	proto.RegisterType((*EndpointError)(nil), "gateway.EndpointError")
	proto.RegisterType((*ProposedTransaction)(nil), "gateway.ProposedTransaction")
	proto.RegisterType((*PreparedTransaction)(nil), "gateway.PreparedTransaction")
}

func init() { proto.RegisterFile("gateway/gateway.proto", fileDescriptor_285396c8df15061f) }

var fileDescriptor_285396c8df15061f = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6b, 0xdb, 0x3e,
	0x14, 0xc7, 0x2d, 0xdf, 0xa4, 0x79, 0x6d, 0xf2, 0x0d, 0x4a, 0x9b, 0x66, 0xa6, 0x85, 0x62, 0x28,
	0xe4, 0xb0, 0xc6, 0x23, 0x3b, 0x0d, 0x0a, 0x83, 0x86, 0xb0, 0x85, 0x5d, 0x42, 0x52, 0x76, 0xe8,
	0x0e, 0x41, 0x89, 0x34, 0xc7, 0x60, 0x5b, 0x9a, 0x24, 0x77, 0xeb, 0x6d, 0x7f, 0xc2, 0x0e, 0x3b,
	0xed, 0xb6, 0xfd, 0xa5, 0xc3, 0x91, 0xe4, 0xd8, 0x4b, 0x33, 0x5a, 0xe8, 0x61, 0x27, 0xa3, 0xf7,
	0xde, 0x47, 0xfa, 0xbc, 0xcf, 0xfb, 0x61, 0x38, 0x0a, 0xb0, 0xa2, 0x9f, 0xf1, 0x9d, 0x6f, 0xbe,
	0x3d, 0x2e, 0x98, 0x62, 0xa8, 0x6a, 0x8e, 0x6e, 0x8b, 0x53, 0x2a, 0x7c, 0x2e, 0x18, 0x67, 0x12,
	0x47, 0xda, 0xeb, 0x9e, 0x94, 0x8c, 0x33, 0x41, 0x25, 0x67, 0x89, 0xa4, 0xc6, 0xdb, 0x5e, 0x79,
	0x95, 0xc0, 0x89, 0xc4, 0x0b, 0x15, 0xb2, 0xc4, 0xd8, 0x5b, 0x0b, 0x16, 0xc7, 0x2c, 0xf1, 0xf5,
	0x47, 0x1b, 0xbd, 0x9f, 0x0e, 0x34, 0x86, 0x09, 0x61, 0x42, 0xd2, 0x09, 0xfd, 0x94, 0x52, 0xa9,
	0xd0, 0x39, 0x34, 0x0a, 0xe0, 0x59, 0x48, 0x3a, 0xce, 0x99, 0xd3, 0xad, 0x4d, 0xea, 0x05, 0xeb,
	0x88, 0xa0, 0x53, 0x80, 0xc5, 0x12, 0x27, 0x09, 0x8d, 0xb2, 0x90, 0x9d, 0x55, 0x48, 0xcd, 0x58,
	0x46, 0x04, 0x8d, 0xe0, 0x50, 0x13, 0xa4, 0x64, 0x56, 0x00, 0x76, 0x76, 0xcf, 0x9c, 0xee, 0x7e,
	0xbf, 0xad, 0x9f, 0x97, 0xbd, 0x69, 0x18, 0x24, 0x94, 0x8c, 0x4d, 0x2a, 0x93, 0x96, 0xc5, 0x5c,
	0xaf, 0x21, 0xde, 0x57, 0x07, 0xfe, 0xcf, 0x39, 0xea, 0x54, 0x51, 0x17, 0x2a, 0x82, 0xca, 0x34,
	0x52, 0x2b, 0x72, 0xfb, 0xfd, 0xa6, 0xbd, 0xd0, 0x46, 0x4c, 0x8c, 0x1f, 0x0d, 0x32, 0x22, 0x94,
	0x63, 0xf1, 0x07, 0x91, 0x1d, 0x83, 0x33, 0x72, 0x0c, 0x93, 0x5b, 0x1a, 0x31, 0x4e, 0x33, 0x0a,
	0x3a, 0xba, 0x48, 0xe1, 0x87, 0x03, 0xf5, 0x69, 0x3a, 0x8f, 0x43, 0xf5, 0xb4, 0x2a, 0x6d, 0x23,
	0xb7, 0xfb, 0x18, 0x72, 0x4d, 0x68, 0x58, 0x6e, 0x3a, 0x77, 0xef, 0x03, 0xb4, 0x06, 0x2c, 0x8e,
	0x43, 0x35, 0x55, 0x58, 0xa5, 0xf2, 0x49, 0x39, 0x7b, 0x6f, 0xe1, 0xb0, 0x7c, 0xb9, 0x29, 0xc9,
	0x8b, 0x52, 0x49, 0x1a, 0xfd, 0x8e, 0x2d, 0xc9, 0xf5, 0x97, 0xf7, 0x38, 0x0a, 0x09, 0xce, 0xae,
	0x1f, 0x30, 0x92, 0x97, 0xc6, 0xfb, 0x95, 0x15, 0xf6, 0x16, 0x47, 0x29, 0x56, 0xff, 0x6e, 0xf7,
	0x5d, 0x42, 0x73, 0xcd, 0xf1, 0xb1, 0xdd, 0xe7, 0xdd, 0x40, 0x7d, 0x98, 0x10, 0xce, 0xc2, 0x44,
	0x0d, 0x85, 0x60, 0x02, 0x75, 0xa0, 0x8a, 0x09, 0x11, 0x54, 0x4a, 0x93, 0x98, 0x3d, 0xa2, 0x23,
	0xa8, 0xc4, 0x92, 0xaf, 0xd3, 0xf9, 0x2f, 0x96, 0x7c, 0x44, 0x32, 0x40, 0x4c, 0xa5, 0xc4, 0x01,
	0x5d, 0xb1, 0xaf, 0x4d, 0xec, 0xd1, 0xe3, 0xd0, 0x1a, 0x6f, 0x12, 0x7e, 0xa8, 0x82, 0x7d, 0xd8,
	0xb3, 0x1b, 0xc4, 0xcc, 0xc2, 0x36, 0x59, 0xf2, 0x38, 0xef, 0xbb, 0x93, 0x3d, 0xb9, 0xd1, 0x81,
	0x0f, 0x7d, 0xf2, 0x39, 0xec, 0x51, 0xd3, 0xc9, 0x5b, 0xc7, 0x2f, 0x8f, 0x28, 0x88, 0xbc, 0xfb,
	0x77, 0x91, 0xfb, 0xdf, 0x76, 0xa0, 0xfa, 0x46, 0x2f, 0x4c, 0x74, 0x09, 0x55, 0xb3, 0x2b, 0xd0,
	0x71, 0xcf, 0x2e, 0xd5, 0xf2, 0x86, 0x73, 0x3b, 0x9b, 0x0e, 0x53, 0xd8, 0x57, 0x50, 0xd1, 0xa3,
	0x84, 0xda, 0x79, 0x4c, 0x69, 0xee, 0xdd, 0xe3, 0x0d, 0xbb, 0x81, 0xbe, 0x83, 0x83, 0xe2, 0x58,
	0xa0, 0x93, 0x3c, 0xf0, 0x9e, 0x51, 0x74, 0x4f, 0xb7, 0x78, 0xcd, 0x65, 0xaf, 0x61, 0xcf, 0x36,
	0x1d, 0x2a, 0xb0, 0x2d, 0xcf, 0x8a, 0xfb, 0xec, 0x1e, 0x8f, 0xbe, 0xe0, 0x6a, 0x09, 0xe7, 0x4c,
	0x04, 0xbd, 0xe5, 0x1d, 0xa7, 0x22, 0xa2, 0x24, 0xa0, 0xa2, 0xf7, 0x11, 0xcf, 0x45, 0xb8, 0xb0,
	0x22, 0x1a, 0xe4, 0xd5, 0x81, 0x11, 0x6e, 0x9c, 0x99, 0xc7, 0xce, 0x8d, 0x1f, 0x84, 0x6a, 0x99,
	0xce, 0xb3, 0xba, 0xf8, 0x05, 0xb4, 0xaf, 0xd1, 0x17, 0x1a, 0x7d, 0x11, 0x30, 0xfb, 0xc3, 0x9a,
	0x57, 0x56, 0xa6, 0x97, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x4a, 0xf8, 0x90, 0xca, 0x06,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	// The Endorse service passes a proposed transaction to the gateway in order to
	// obtain sufficient endorsement.
	// The gateway will determine the endorsement plan for the requested chaincode and
	// forward to the appropriate peers for endorsement. It will return to the client a
	// prepared transaction in the form of an Envelope message as defined
	// in common/common.proto. The client must sign the contents of this envelope
	// before invoking the Submit service
	Endorse(ctx context.Context, in *EndorseRequest, opts ...grpc.CallOption) (*EndorseResponse, error)
	// The Submit service will process the prepared transaction returned from Endorse service
	// once it has been signed by the client. It will wait for the transaction to be submitted to the
	// ordering service but the client must invoke the CommitStatus service to wait for the transaction
	// to be committed.
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	// The CommitStatus service will indicate whether a prepared transaction previously submitted to
	// the Submit sevice has been committed. It will wait for the commit to occur if it hasn???t already
	// committed.
	CommitStatus(ctx context.Context, in *CommitStatusRequest, opts ...grpc.CallOption) (*CommitStatusResponse, error)
	// The Evaluate service passes a proposed transaction to the gateway in order to invoke the
	// transaction function and return the result to the client. No ledger updates are made.
	// The gateway will select an appropriate peer to query based on block height and load.
	Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Endorse(ctx context.Context, in *EndorseRequest, opts ...grpc.CallOption) (*EndorseResponse, error) {
	out := new(EndorseResponse)
	err := c.cc.Invoke(ctx, "/gateway.Gateway/Endorse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/gateway.Gateway/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CommitStatus(ctx context.Context, in *CommitStatusRequest, opts ...grpc.CallOption) (*CommitStatusResponse, error) {
	out := new(CommitStatusResponse)
	err := c.cc.Invoke(ctx, "/gateway.Gateway/CommitStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error) {
	out := new(EvaluateResponse)
	err := c.cc.Invoke(ctx, "/gateway.Gateway/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	// The Endorse service passes a proposed transaction to the gateway in order to
	// obtain sufficient endorsement.
	// The gateway will determine the endorsement plan for the requested chaincode and
	// forward to the appropriate peers for endorsement. It will return to the client a
	// prepared transaction in the form of an Envelope message as defined
	// in common/common.proto. The client must sign the contents of this envelope
	// before invoking the Submit service
	Endorse(context.Context, *EndorseRequest) (*EndorseResponse, error)
	// The Submit service will process the prepared transaction returned from Endorse service
	// once it has been signed by the client. It will wait for the transaction to be submitted to the
	// ordering service but the client must invoke the CommitStatus service to wait for the transaction
	// to be committed.
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	// The CommitStatus service will indicate whether a prepared transaction previously submitted to
	// the Submit sevice has been committed. It will wait for the commit to occur if it hasn???t already
	// committed.
	CommitStatus(context.Context, *CommitStatusRequest) (*CommitStatusResponse, error)
	// The Evaluate service passes a proposed transaction to the gateway in order to invoke the
	// transaction function and return the result to the client. No ledger updates are made.
	// The gateway will select an appropriate peer to query based on block height and load.
	Evaluate(context.Context, *EvaluateRequest) (*EvaluateResponse, error)
}

// UnimplementedGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (*UnimplementedGatewayServer) Endorse(ctx context.Context, req *EndorseRequest) (*EndorseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Endorse not implemented")
}
func (*UnimplementedGatewayServer) Submit(ctx context.Context, req *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (*UnimplementedGatewayServer) CommitStatus(ctx context.Context, req *CommitStatusRequest) (*CommitStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitStatus not implemented")
}
func (*UnimplementedGatewayServer) Evaluate(ctx context.Context, req *EvaluateRequest) (*EvaluateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_Endorse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndorseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Endorse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/Endorse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Endorse(ctx, req.(*EndorseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CommitStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CommitStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/CommitStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CommitStatus(ctx, req.(*CommitStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Evaluate(ctx, req.(*EvaluateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Endorse",
			Handler:    _Gateway_Endorse_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _Gateway_Submit_Handler,
		},
		{
			MethodName: "CommitStatus",
			Handler:    _Gateway_CommitStatus_Handler,
		},
		{
			MethodName: "Evaluate",
			Handler:    _Gateway_Evaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/gateway.proto",
}
