// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: api/shop/avito-shop.proto

package shop

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BuyItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          string                 `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BuyItemRequest) Reset() {
	*x = BuyItemRequest{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuyItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyItemRequest) ProtoMessage() {}

func (x *BuyItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyItemRequest.ProtoReflect.Descriptor instead.
func (*BuyItemRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{0}
}

func (x *BuyItemRequest) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

type InfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{1}
}

func (x *InfoRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type InfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Coins         int32                  `protobuf:"varint,1,opt,name=coins,proto3" json:"coins,omitempty"`
	Inventory     []*InventoryItem       `protobuf:"bytes,2,rep,name=inventory,proto3" json:"inventory,omitempty"`
	CoinHistory   *CoinHistory           `protobuf:"bytes,3,opt,name=coinHistory,proto3" json:"coinHistory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{2}
}

func (x *InfoResponse) GetCoins() int32 {
	if x != nil {
		return x.Coins
	}
	return 0
}

func (x *InfoResponse) GetInventory() []*InventoryItem {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *InfoResponse) GetCoinHistory() *CoinHistory {
	if x != nil {
		return x.CoinHistory
	}
	return nil
}

type InventoryItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InventoryItem) Reset() {
	*x = InventoryItem{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryItem) ProtoMessage() {}

func (x *InventoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryItem.ProtoReflect.Descriptor instead.
func (*InventoryItem) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{3}
}

func (x *InventoryItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InventoryItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CoinHistory struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Received      []*Received            `protobuf:"bytes,1,rep,name=received,proto3" json:"received,omitempty"`
	Sent          []*Sent                `protobuf:"bytes,2,rep,name=sent,proto3" json:"sent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CoinHistory) Reset() {
	*x = CoinHistory{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoinHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinHistory) ProtoMessage() {}

func (x *CoinHistory) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinHistory.ProtoReflect.Descriptor instead.
func (*CoinHistory) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{4}
}

func (x *CoinHistory) GetReceived() []*Received {
	if x != nil {
		return x.Received
	}
	return nil
}

func (x *CoinHistory) GetSent() []*Sent {
	if x != nil {
		return x.Sent
	}
	return nil
}

type Received struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromUser      string                 `protobuf:"bytes,1,opt,name=fromUser,proto3" json:"fromUser,omitempty"`
	Amount        int32                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Received) Reset() {
	*x = Received{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Received) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Received) ProtoMessage() {}

func (x *Received) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Received.ProtoReflect.Descriptor instead.
func (*Received) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{5}
}

func (x *Received) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *Received) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Sent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ToUser        string                 `protobuf:"bytes,1,opt,name=toUser,proto3" json:"toUser,omitempty"`
	Amount        int32                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sent) Reset() {
	*x = Sent{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sent) ProtoMessage() {}

func (x *Sent) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sent.ProtoReflect.Descriptor instead.
func (*Sent) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{6}
}

func (x *Sent) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *Sent) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AuthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{7}
}

func (x *AuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{8}
}

func (x *AuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SendCoinRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromUser      string                 `protobuf:"bytes,1,opt,name=fromUser,proto3" json:"fromUser,omitempty"`
	ToUser        string                 `protobuf:"bytes,2,opt,name=toUser,proto3" json:"toUser,omitempty"`
	Amount        int32                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendCoinRequest) Reset() {
	*x = SendCoinRequest{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCoinRequest) ProtoMessage() {}

func (x *SendCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCoinRequest.ProtoReflect.Descriptor instead.
func (*SendCoinRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{9}
}

func (x *SendCoinRequest) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *SendCoinRequest) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *SendCoinRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SendCoinResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	RemainingCoins int32                  `protobuf:"varint,1,opt,name=remaining_coins,json=remainingCoins,proto3" json:"remaining_coins,omitempty"`
	TransactionId  string                 `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SendCoinResponse) Reset() {
	*x = SendCoinResponse{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCoinResponse) ProtoMessage() {}

func (x *SendCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCoinResponse.ProtoReflect.Descriptor instead.
func (*SendCoinResponse) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{10}
}

func (x *SendCoinResponse) GetRemainingCoins() int32 {
	if x != nil {
		return x.RemainingCoins
	}
	return 0
}

func (x *SendCoinResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type BuyItemResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	RemainingCoins int32                  `protobuf:"varint,1,opt,name=remaining_coins,json=remainingCoins,proto3" json:"remaining_coins,omitempty"`
	Items          []*InventoryItem       `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BuyItemResponse) Reset() {
	*x = BuyItemResponse{}
	mi := &file_api_shop_avito_shop_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuyItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyItemResponse) ProtoMessage() {}

func (x *BuyItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_avito_shop_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyItemResponse.ProtoReflect.Descriptor instead.
func (*BuyItemResponse) Descriptor() ([]byte, []int) {
	return file_api_shop_avito_shop_proto_rawDescGZIP(), []int{11}
}

func (x *BuyItemResponse) GetRemainingCoins() int32 {
	if x != nil {
		return x.RemainingCoins
	}
	return 0
}

func (x *BuyItemResponse) GetItems() []*InventoryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var file_api_shop_avito_shop_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50000,
		Name:          "avito.shop.v1.requires_auth",
		Tag:           "varint,50000,opt,name=requires_auth",
		Filename:      "api/shop/avito-shop.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional bool requires_auth = 50000;
	E_RequiresAuth = &file_api_shop_avito_shop_proto_extTypes[0]
)

var File_api_shop_avito_shop_proto protoreflect.FileDescriptor

var file_api_shop_avito_shop_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x76, 0x69, 0x74, 0x6f,
	0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x76, 0x69,
	0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x42, 0x75, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x29, 0x0a,
	0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x69,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x12,
	0x3a, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3c, 0x0a, 0x0b, 0x63,
	0x6f, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f,
	0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x3f, 0x0a, 0x0d, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x6b, 0x0a, 0x0b, 0x43, 0x6f,
	0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x76,
	0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x27,
	0x0a, 0x04, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61,
	0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x74, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x45, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5d, 0x0a, 0x0f,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x10, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x69,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x6e, 0x0a, 0x0f, 0x42, 0x75, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x6d,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x76, 0x69,
	0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32,
	0x93, 0x03, 0x0a, 0x09, 0x41, 0x76, 0x69, 0x74, 0x6f, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x59, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x69, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x6f, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x69, 0x6e, 0x12, 0x65, 0x0a, 0x07, 0x42, 0x75, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d,
	0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x75, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x80,
	0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x62, 0x75, 0x79, 0x2f, 0x7b, 0x69, 0x74, 0x65, 0x6d, 0x7d, 0x12, 0x59, 0x0a, 0x04, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x1a, 0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x61, 0x76, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x80, 0xb5, 0x18,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x45, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x41, 0x75, 0x74, 0x68, 0x42, 0x11, 0x5a, 0x0f,
	0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x3b, 0x73, 0x68, 0x6f, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_shop_avito_shop_proto_rawDescOnce sync.Once
	file_api_shop_avito_shop_proto_rawDescData = file_api_shop_avito_shop_proto_rawDesc
)

func file_api_shop_avito_shop_proto_rawDescGZIP() []byte {
	file_api_shop_avito_shop_proto_rawDescOnce.Do(func() {
		file_api_shop_avito_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_shop_avito_shop_proto_rawDescData)
	})
	return file_api_shop_avito_shop_proto_rawDescData
}

var file_api_shop_avito_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_shop_avito_shop_proto_goTypes = []any{
	(*BuyItemRequest)(nil),             // 0: avito.shop.v1.BuyItemRequest
	(*InfoRequest)(nil),                // 1: avito.shop.v1.InfoRequest
	(*InfoResponse)(nil),               // 2: avito.shop.v1.InfoResponse
	(*InventoryItem)(nil),              // 3: avito.shop.v1.InventoryItem
	(*CoinHistory)(nil),                // 4: avito.shop.v1.CoinHistory
	(*Received)(nil),                   // 5: avito.shop.v1.Received
	(*Sent)(nil),                       // 6: avito.shop.v1.Sent
	(*AuthRequest)(nil),                // 7: avito.shop.v1.AuthRequest
	(*AuthResponse)(nil),               // 8: avito.shop.v1.AuthResponse
	(*SendCoinRequest)(nil),            // 9: avito.shop.v1.SendCoinRequest
	(*SendCoinResponse)(nil),           // 10: avito.shop.v1.SendCoinResponse
	(*BuyItemResponse)(nil),            // 11: avito.shop.v1.BuyItemResponse
	(*descriptorpb.MethodOptions)(nil), // 12: google.protobuf.MethodOptions
}
var file_api_shop_avito_shop_proto_depIdxs = []int32{
	3,  // 0: avito.shop.v1.InfoResponse.inventory:type_name -> avito.shop.v1.InventoryItem
	4,  // 1: avito.shop.v1.InfoResponse.coinHistory:type_name -> avito.shop.v1.CoinHistory
	5,  // 2: avito.shop.v1.CoinHistory.received:type_name -> avito.shop.v1.Received
	6,  // 3: avito.shop.v1.CoinHistory.sent:type_name -> avito.shop.v1.Sent
	3,  // 4: avito.shop.v1.BuyItemResponse.items:type_name -> avito.shop.v1.InventoryItem
	12, // 5: avito.shop.v1.requires_auth:extendee -> google.protobuf.MethodOptions
	1,  // 6: avito.shop.v1.AvitoShop.GetInfo:input_type -> avito.shop.v1.InfoRequest
	9,  // 7: avito.shop.v1.AvitoShop.SendCoin:input_type -> avito.shop.v1.SendCoinRequest
	0,  // 8: avito.shop.v1.AvitoShop.BuyItem:input_type -> avito.shop.v1.BuyItemRequest
	7,  // 9: avito.shop.v1.AvitoShop.Auth:input_type -> avito.shop.v1.AuthRequest
	2,  // 10: avito.shop.v1.AvitoShop.GetInfo:output_type -> avito.shop.v1.InfoResponse
	10, // 11: avito.shop.v1.AvitoShop.SendCoin:output_type -> avito.shop.v1.SendCoinResponse
	11, // 12: avito.shop.v1.AvitoShop.BuyItem:output_type -> avito.shop.v1.BuyItemResponse
	8,  // 13: avito.shop.v1.AvitoShop.Auth:output_type -> avito.shop.v1.AuthResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	5,  // [5:6] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_shop_avito_shop_proto_init() }
func file_api_shop_avito_shop_proto_init() {
	if File_api_shop_avito_shop_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_shop_avito_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_api_shop_avito_shop_proto_goTypes,
		DependencyIndexes: file_api_shop_avito_shop_proto_depIdxs,
		MessageInfos:      file_api_shop_avito_shop_proto_msgTypes,
		ExtensionInfos:    file_api_shop_avito_shop_proto_extTypes,
	}.Build()
	File_api_shop_avito_shop_proto = out.File
	file_api_shop_avito_shop_proto_rawDesc = nil
	file_api_shop_avito_shop_proto_goTypes = nil
	file_api_shop_avito_shop_proto_depIdxs = nil
}
