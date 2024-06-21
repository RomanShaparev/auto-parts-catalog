// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: warehouse_position_service.proto

package gen

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

type CreateOrUpdateWarehousePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarehouseId         int32 `protobuf:"varint,1,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id,omitempty"`
	AutoPartComponentId int32 `protobuf:"varint,2,opt,name=auto_part_component_id,json=autoPartComponentId,proto3" json:"auto_part_component_id,omitempty"`
	Quantity            int32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CreateOrUpdateWarehousePositionRequest) Reset() {
	*x = CreateOrUpdateWarehousePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_position_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateWarehousePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateWarehousePositionRequest) ProtoMessage() {}

func (x *CreateOrUpdateWarehousePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_position_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateWarehousePositionRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateWarehousePositionRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_position_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrUpdateWarehousePositionRequest) GetWarehouseId() int32 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *CreateOrUpdateWarehousePositionRequest) GetAutoPartComponentId() int32 {
	if x != nil {
		return x.AutoPartComponentId
	}
	return 0
}

func (x *CreateOrUpdateWarehousePositionRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetWarehousePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarehouseId         int32 `protobuf:"varint,1,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id,omitempty"`
	AutoPartComponentId int32 `protobuf:"varint,2,opt,name=auto_part_component_id,json=autoPartComponentId,proto3" json:"auto_part_component_id,omitempty"`
}

func (x *GetWarehousePositionRequest) Reset() {
	*x = GetWarehousePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_position_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWarehousePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWarehousePositionRequest) ProtoMessage() {}

func (x *GetWarehousePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_position_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWarehousePositionRequest.ProtoReflect.Descriptor instead.
func (*GetWarehousePositionRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_position_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetWarehousePositionRequest) GetWarehouseId() int32 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *GetWarehousePositionRequest) GetAutoPartComponentId() int32 {
	if x != nil {
		return x.AutoPartComponentId
	}
	return 0
}

type WarehousePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarehouseId         int32 `protobuf:"varint,1,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id,omitempty"`
	AutoPartComponentId int32 `protobuf:"varint,2,opt,name=auto_part_component_id,json=autoPartComponentId,proto3" json:"auto_part_component_id,omitempty"`
	Quantity            int32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *WarehousePosition) Reset() {
	*x = WarehousePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_position_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehousePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehousePosition) ProtoMessage() {}

func (x *WarehousePosition) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_position_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehousePosition.ProtoReflect.Descriptor instead.
func (*WarehousePosition) Descriptor() ([]byte, []int) {
	return file_warehouse_position_service_proto_rawDescGZIP(), []int{2}
}

func (x *WarehousePosition) GetWarehouseId() int32 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *WarehousePosition) GetAutoPartComponentId() int32 {
	if x != nil {
		return x.AutoPartComponentId
	}
	return 0
}

func (x *WarehousePosition) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_warehouse_position_service_proto protoreflect.FileDescriptor

var file_warehouse_position_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x67, 0x65, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x26, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x50, 0x61, 0x72, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x75, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x50, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x87, 0x01,
	0x0a, 0x11, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x50, 0x61, 0x72, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xd4, 0x01, 0x0a, 0x18, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x37,
	0x5a, 0x35, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x74, 0x73, 0x2d, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warehouse_position_service_proto_rawDescOnce sync.Once
	file_warehouse_position_service_proto_rawDescData = file_warehouse_position_service_proto_rawDesc
)

func file_warehouse_position_service_proto_rawDescGZIP() []byte {
	file_warehouse_position_service_proto_rawDescOnce.Do(func() {
		file_warehouse_position_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouse_position_service_proto_rawDescData)
	})
	return file_warehouse_position_service_proto_rawDescData
}

var file_warehouse_position_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_warehouse_position_service_proto_goTypes = []interface{}{
	(*CreateOrUpdateWarehousePositionRequest)(nil), // 0: gen.CreateOrUpdateWarehousePositionRequest
	(*GetWarehousePositionRequest)(nil),            // 1: gen.GetWarehousePositionRequest
	(*WarehousePosition)(nil),                      // 2: gen.WarehousePosition
}
var file_warehouse_position_service_proto_depIdxs = []int32{
	0, // 0: gen.WarehousePositionService.CreateOrUpdateWarehousePosition:input_type -> gen.CreateOrUpdateWarehousePositionRequest
	1, // 1: gen.WarehousePositionService.GetWarehousePosition:input_type -> gen.GetWarehousePositionRequest
	2, // 2: gen.WarehousePositionService.CreateOrUpdateWarehousePosition:output_type -> gen.WarehousePosition
	2, // 3: gen.WarehousePositionService.GetWarehousePosition:output_type -> gen.WarehousePosition
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_warehouse_position_service_proto_init() }
func file_warehouse_position_service_proto_init() {
	if File_warehouse_position_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_warehouse_position_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateWarehousePositionRequest); i {
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
		file_warehouse_position_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWarehousePositionRequest); i {
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
		file_warehouse_position_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehousePosition); i {
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
			RawDescriptor: file_warehouse_position_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_warehouse_position_service_proto_goTypes,
		DependencyIndexes: file_warehouse_position_service_proto_depIdxs,
		MessageInfos:      file_warehouse_position_service_proto_msgTypes,
	}.Build()
	File_warehouse_position_service_proto = out.File
	file_warehouse_position_service_proto_rawDesc = nil
	file_warehouse_position_service_proto_goTypes = nil
	file_warehouse_position_service_proto_depIdxs = nil
}