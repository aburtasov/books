// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: books.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName string `protobuf:"bytes,2,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Author) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Books struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title []string `protobuf:"bytes,1,rep,name=title,proto3" json:"title,omitempty"`
}

func (x *Books) Reset() {
	*x = Books{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Books) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Books) ProtoMessage() {}

func (x *Books) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Books.ProtoReflect.Descriptor instead.
func (*Books) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{2}
}

func (x *Books) GetTitle() []string {
	if x != nil {
		return x.Title
	}
	return nil
}

type ResponceID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResponceID) Reset() {
	*x = ResponceID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceID) ProtoMessage() {}

func (x *ResponceID) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceID.ProtoReflect.Descriptor instead.
func (*ResponceID) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{3}
}

func (x *ResponceID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AddBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AuthorId    int64  `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *AddBook) Reset() {
	*x = AddBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBook) ProtoMessage() {}

func (x *AddBook) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBook.ProtoReflect.Descriptor instead.
func (*AddBook) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{4}
}

func (x *AddBook) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddBook) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddBook) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

var File_books_proto protoreflect.FileDescriptor

var file_books_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x22, 0x48, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x04,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5e, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x32, 0xaf, 0x01,
	0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x00, 0x12,
	0x25, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x1a,
	0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x49, 0x44,
	0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x0f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x22, 0x00, 0x42,
	0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62,
	0x75, 0x72, 0x74, 0x61, 0x73, 0x6f, 0x76, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_proto_rawDescOnce sync.Once
	file_books_proto_rawDescData = file_books_proto_rawDesc
)

func file_books_proto_rawDescGZIP() []byte {
	file_books_proto_rawDescOnce.Do(func() {
		file_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_proto_rawDescData)
	})
	return file_books_proto_rawDescData
}

var file_books_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_books_proto_goTypes = []interface{}{
	(*Author)(nil),     // 0: api.Author
	(*Book)(nil),       // 1: api.Book
	(*Books)(nil),      // 2: api.Books
	(*ResponceID)(nil), // 3: api.ResponceID
	(*AddBook)(nil),    // 4: api.AddBook
}
var file_books_proto_depIdxs = []int32{
	0, // 0: api.Booker.GetBooks:input_type -> api.Author
	1, // 1: api.Booker.GetAuthor:input_type -> api.Book
	4, // 2: api.Booker.addBook:input_type -> api.AddBook
	0, // 3: api.Booker.addAuthor:input_type -> api.Author
	2, // 4: api.Booker.GetBooks:output_type -> api.Books
	0, // 5: api.Booker.GetAuthor:output_type -> api.Author
	3, // 6: api.Booker.addBook:output_type -> api.ResponceID
	3, // 7: api.Booker.addAuthor:output_type -> api.ResponceID
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_books_proto_init() }
func file_books_proto_init() {
	if File_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Books); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBook); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_proto_goTypes,
		DependencyIndexes: file_books_proto_depIdxs,
		MessageInfos:      file_books_proto_msgTypes,
	}.Build()
	File_books_proto = out.File
	file_books_proto_rawDesc = nil
	file_books_proto_goTypes = nil
	file_books_proto_depIdxs = nil
}
