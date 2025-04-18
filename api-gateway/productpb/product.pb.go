// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: protos/product.proto

package productpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_protos_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProductUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age           int32                  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductUser) Reset() {
	*x = ProductUser{}
	mi := &file_protos_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUser) ProtoMessage() {}

func (x *ProductUser) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUser.ProtoReflect.Descriptor instead.
func (*ProductUser) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductUser) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductUser) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	UserId        int32                  `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User          *ProductUser           `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_protos_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Product) GetUser() *ProductUser {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAllProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllProductRequest) Reset() {
	*x = GetAllProductRequest{}
	mi := &file_protos_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductRequest) ProtoMessage() {}

func (x *GetAllProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductRequest.ProtoReflect.Descriptor instead.
func (*GetAllProductRequest) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{3}
}

type GetAllProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       []*Product             `protobuf:"bytes,1,rep,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllProductResponse) Reset() {
	*x = GetAllProductResponse{}
	mi := &file_protos_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductResponse) ProtoMessage() {}

func (x *GetAllProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductResponse) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProductResponse) GetProduct() []*Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_protos_product_proto protoreflect.FileDescriptor

const file_protos_product_proto_rawDesc = "" +
	"\n" +
	"\x14protos/product.proto\x12\x04main\"#\n" +
	"\x11GetProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"C\n" +
	"\vProductUser\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x10\n" +
	"\x03age\x18\x03 \x01(\x05R\x03age\"\x83\x01\n" +
	"\aProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x02R\x05price\x12\x17\n" +
	"\auser_id\x18\x05 \x01(\x05R\x06userId\x12%\n" +
	"\x04user\x18\x04 \x01(\v2\x11.main.ProductUserR\x04user\"\x16\n" +
	"\x14GetAllProductRequest\"@\n" +
	"\x15GetAllProductResponse\x12'\n" +
	"\aproduct\x18\x01 \x03(\v2\r.main.ProductR\aproduct2\x90\x01\n" +
	"\x0eProductService\x124\n" +
	"\n" +
	"GetProduct\x12\x17.main.GetProductRequest\x1a\r.main.Product\x12H\n" +
	"\rGetAllProduct\x12\x1a.main.GetAllProductRequest\x1a\x1b.main.GetAllProductResponseB\rZ\v./productpbb\x06proto3"

var (
	file_protos_product_proto_rawDescOnce sync.Once
	file_protos_product_proto_rawDescData []byte
)

func file_protos_product_proto_rawDescGZIP() []byte {
	file_protos_product_proto_rawDescOnce.Do(func() {
		file_protos_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_product_proto_rawDesc), len(file_protos_product_proto_rawDesc)))
	})
	return file_protos_product_proto_rawDescData
}

var file_protos_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_product_proto_goTypes = []any{
	(*GetProductRequest)(nil),     // 0: main.GetProductRequest
	(*ProductUser)(nil),           // 1: main.ProductUser
	(*Product)(nil),               // 2: main.Product
	(*GetAllProductRequest)(nil),  // 3: main.GetAllProductRequest
	(*GetAllProductResponse)(nil), // 4: main.GetAllProductResponse
}
var file_protos_product_proto_depIdxs = []int32{
	1, // 0: main.Product.user:type_name -> main.ProductUser
	2, // 1: main.GetAllProductResponse.product:type_name -> main.Product
	0, // 2: main.ProductService.GetProduct:input_type -> main.GetProductRequest
	3, // 3: main.ProductService.GetAllProduct:input_type -> main.GetAllProductRequest
	2, // 4: main.ProductService.GetProduct:output_type -> main.Product
	4, // 5: main.ProductService.GetAllProduct:output_type -> main.GetAllProductResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_product_proto_init() }
func file_protos_product_proto_init() {
	if File_protos_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_product_proto_rawDesc), len(file_protos_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_product_proto_goTypes,
		DependencyIndexes: file_protos_product_proto_depIdxs,
		MessageInfos:      file_protos_product_proto_msgTypes,
	}.Build()
	File_protos_product_proto = out.File
	file_protos_product_proto_goTypes = nil
	file_protos_product_proto_depIdxs = nil
}
