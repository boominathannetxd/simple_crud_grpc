// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/moviesapp.proto

package protos

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_moviesapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_moviesapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_moviesapp_proto_rawDescGZIP(), []int{0}
}

type MovieInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Isbn     string    `protobuf:"bytes,2,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Title    string    `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Director *Director `protobuf:"bytes,4,opt,name=director,proto3" json:"director,omitempty"`
}

func (x *MovieInfo) Reset() {
	*x = MovieInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_moviesapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieInfo) ProtoMessage() {}

func (x *MovieInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_moviesapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieInfo.ProtoReflect.Descriptor instead.
func (*MovieInfo) Descriptor() ([]byte, []int) {
	return file_protos_moviesapp_proto_rawDescGZIP(), []int{1}
}

func (x *MovieInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MovieInfo) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

func (x *MovieInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieInfo) GetDirector() *Director {
	if x != nil {
		return x.Director
	}
	return nil
}

type Director struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
}

func (x *Director) Reset() {
	*x = Director{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_moviesapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Director) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Director) ProtoMessage() {}

func (x *Director) ProtoReflect() protoreflect.Message {
	mi := &file_protos_moviesapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Director.ProtoReflect.Descriptor instead.
func (*Director) Descriptor() ([]byte, []int) {
	return file_protos_moviesapp_proto_rawDescGZIP(), []int{2}
}

func (x *Director) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Director) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_moviesapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_protos_moviesapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_protos_moviesapp_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_moviesapp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protos_moviesapp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protos_moviesapp_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_protos_moviesapp_proto protoreflect.FileDescriptor

var file_protos_moviesapp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61,
	0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x61, 0x70, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x76, 0x0a, 0x09,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70,
	0x70, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x22, 0x44, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x02, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x8c, 0x02, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x10, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x14, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e,
	0x49, 0x64, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x61, 0x70, 0x70, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0d, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e,
	0x49, 0x64, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x19, 0x5a, 0x17, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x63, 0x72, 0x75, 0x64, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_moviesapp_proto_rawDescOnce sync.Once
	file_protos_moviesapp_proto_rawDescData = file_protos_moviesapp_proto_rawDesc
)

func file_protos_moviesapp_proto_rawDescGZIP() []byte {
	file_protos_moviesapp_proto_rawDescOnce.Do(func() {
		file_protos_moviesapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_moviesapp_proto_rawDescData)
	})
	return file_protos_moviesapp_proto_rawDescData
}

var file_protos_moviesapp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_moviesapp_proto_goTypes = []interface{}{
	(*Empty)(nil),     // 0: moviesapp.Empty
	(*MovieInfo)(nil), // 1: moviesapp.MovieInfo
	(*Director)(nil),  // 2: moviesapp.Director
	(*Id)(nil),        // 3: moviesapp.Id
	(*Status)(nil),    // 4: moviesapp.Status
}
var file_protos_moviesapp_proto_depIdxs = []int32{
	2, // 0: moviesapp.MovieInfo.director:type_name -> moviesapp.Director
	0, // 1: moviesapp.Movie.GetMovies:input_type -> moviesapp.Empty
	3, // 2: moviesapp.Movie.GetMovie:input_type -> moviesapp.Id
	1, // 3: moviesapp.Movie.CreateMovie:input_type -> moviesapp.MovieInfo
	1, // 4: moviesapp.Movie.UpdateMovie:input_type -> moviesapp.MovieInfo
	3, // 5: moviesapp.Movie.DeleteMovie:input_type -> moviesapp.Id
	1, // 6: moviesapp.Movie.GetMovies:output_type -> moviesapp.MovieInfo
	1, // 7: moviesapp.Movie.GetMovie:output_type -> moviesapp.MovieInfo
	3, // 8: moviesapp.Movie.CreateMovie:output_type -> moviesapp.Id
	4, // 9: moviesapp.Movie.UpdateMovie:output_type -> moviesapp.Status
	4, // 10: moviesapp.Movie.DeleteMovie:output_type -> moviesapp.Status
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_moviesapp_proto_init() }
func file_protos_moviesapp_proto_init() {
	if File_protos_moviesapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_moviesapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_moviesapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieInfo); i {
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
		file_protos_moviesapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Director); i {
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
		file_protos_moviesapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_protos_moviesapp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_protos_moviesapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_moviesapp_proto_goTypes,
		DependencyIndexes: file_protos_moviesapp_proto_depIdxs,
		MessageInfos:      file_protos_moviesapp_proto_msgTypes,
	}.Build()
	File_protos_moviesapp_proto = out.File
	file_protos_moviesapp_proto_rawDesc = nil
	file_protos_moviesapp_proto_goTypes = nil
	file_protos_moviesapp_proto_depIdxs = nil
}
