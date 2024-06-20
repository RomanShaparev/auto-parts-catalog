// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: car_model_service.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCarModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCarModelRequest) Reset() {
	*x = CreateCarModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_model_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarModelRequest) ProtoMessage() {}

func (x *CreateCarModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_model_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarModelRequest.ProtoReflect.Descriptor instead.
func (*CreateCarModelRequest) Descriptor() ([]byte, []int) {
	return file_car_model_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCarModelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCarModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCarModelRequest) Reset() {
	*x = GetCarModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_model_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarModelRequest) ProtoMessage() {}

func (x *GetCarModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_model_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarModelRequest.ProtoReflect.Descriptor instead.
func (*GetCarModelRequest) Descriptor() ([]byte, []int) {
	return file_car_model_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCarModelRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListCarModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarModels []*CarModel `protobuf:"bytes,1,rep,name=CarModels,proto3" json:"CarModels,omitempty"`
}

func (x *ListCarModelsResponse) Reset() {
	*x = ListCarModelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_model_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarModelsResponse) ProtoMessage() {}

func (x *ListCarModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_model_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarModelsResponse.ProtoReflect.Descriptor instead.
func (*ListCarModelsResponse) Descriptor() ([]byte, []int) {
	return file_car_model_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListCarModelsResponse) GetCarModels() []*CarModel {
	if x != nil {
		return x.CarModels
	}
	return nil
}

type DeleteCarModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCarModelRequest) Reset() {
	*x = DeleteCarModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_model_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarModelRequest) ProtoMessage() {}

func (x *DeleteCarModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_model_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarModelRequest.ProtoReflect.Descriptor instead.
func (*DeleteCarModelRequest) Descriptor() ([]byte, []int) {
	return file_car_model_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCarModelRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CarModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CarModel) Reset() {
	*x = CarModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_model_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarModel) ProtoMessage() {}

func (x *CarModel) ProtoReflect() protoreflect.Message {
	mi := &file_car_model_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarModel.ProtoReflect.Descriptor instead.
func (*CarModel) Descriptor() ([]byte, []int) {
	return file_car_model_service_proto_rawDescGZIP(), []int{4}
}

func (x *CarModel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_car_model_service_proto protoreflect.FileDescriptor

var file_car_model_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x65, 0x6e, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44,
	0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x43, 0x61, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a,
	0x08, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x90, 0x02,
	0x0a, 0x0f, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x35,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x17, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x43, 0x61, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x43, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x37, 0x5a, 0x35, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x74, 0x73, 0x2d, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_car_model_service_proto_rawDescOnce sync.Once
	file_car_model_service_proto_rawDescData = file_car_model_service_proto_rawDesc
)

func file_car_model_service_proto_rawDescGZIP() []byte {
	file_car_model_service_proto_rawDescOnce.Do(func() {
		file_car_model_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_model_service_proto_rawDescData)
	})
	return file_car_model_service_proto_rawDescData
}

var file_car_model_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_car_model_service_proto_goTypes = []interface{}{
	(*CreateCarModelRequest)(nil), // 0: gen.CreateCarModelRequest
	(*GetCarModelRequest)(nil),    // 1: gen.GetCarModelRequest
	(*ListCarModelsResponse)(nil), // 2: gen.ListCarModelsResponse
	(*DeleteCarModelRequest)(nil), // 3: gen.DeleteCarModelRequest
	(*CarModel)(nil),              // 4: gen.CarModel
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_car_model_service_proto_depIdxs = []int32{
	4, // 0: gen.ListCarModelsResponse.CarModels:type_name -> gen.CarModel
	0, // 1: gen.CarModelService.CreateCarModel:input_type -> gen.CreateCarModelRequest
	1, // 2: gen.CarModelService.GetCarModel:input_type -> gen.GetCarModelRequest
	5, // 3: gen.CarModelService.ListCarModels:input_type -> google.protobuf.Empty
	3, // 4: gen.CarModelService.DeleteCarModel:input_type -> gen.DeleteCarModelRequest
	4, // 5: gen.CarModelService.CreateCarModel:output_type -> gen.CarModel
	4, // 6: gen.CarModelService.GetCarModel:output_type -> gen.CarModel
	2, // 7: gen.CarModelService.ListCarModels:output_type -> gen.ListCarModelsResponse
	5, // 8: gen.CarModelService.DeleteCarModel:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_car_model_service_proto_init() }
func file_car_model_service_proto_init() {
	if File_car_model_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_model_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarModelRequest); i {
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
		file_car_model_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarModelRequest); i {
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
		file_car_model_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarModelsResponse); i {
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
		file_car_model_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarModelRequest); i {
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
		file_car_model_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarModel); i {
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
			RawDescriptor: file_car_model_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_model_service_proto_goTypes,
		DependencyIndexes: file_car_model_service_proto_depIdxs,
		MessageInfos:      file_car_model_service_proto_msgTypes,
	}.Build()
	File_car_model_service_proto = out.File
	file_car_model_service_proto_rawDesc = nil
	file_car_model_service_proto_goTypes = nil
	file_car_model_service_proto_depIdxs = nil
}
