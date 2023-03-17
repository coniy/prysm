// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: proto/prysm/v1alpha1/blobs.proto

package eth

import (
	github_com_prysmaticlabs_prysm_v4_consensus_types_primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	v1 "github.com/prysmaticlabs/prysm/v4/proto/engine/v1"
	_ "github.com/prysmaticlabs/prysm/v4/proto/eth/ext"
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

type BlobsSidecar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeaconBlockRoot []byte                                                            `protobuf:"bytes,1,opt,name=beacon_block_root,json=beaconBlockRoot,proto3" json:"beacon_block_root,omitempty" ssz-size:"32"`
	BeaconBlockSlot github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot `protobuf:"varint,2,opt,name=beacon_block_slot,json=beaconBlockSlot,proto3" json:"beacon_block_slot,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot"`
	Blobs           []*v1.Blob                                                        `protobuf:"bytes,3,rep,name=blobs,proto3" json:"blobs,omitempty" ssz-max:"16"`
	AggregatedProof []byte                                                            `protobuf:"bytes,4,opt,name=aggregated_proof,json=aggregatedProof,proto3" json:"aggregated_proof,omitempty" ssz-size:"48"`
}

func (x *BlobsSidecar) Reset() {
	*x = BlobsSidecar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_v1alpha1_blobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobsSidecar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobsSidecar) ProtoMessage() {}

func (x *BlobsSidecar) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_v1alpha1_blobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobsSidecar.ProtoReflect.Descriptor instead.
func (*BlobsSidecar) Descriptor() ([]byte, []int) {
	return file_proto_prysm_v1alpha1_blobs_proto_rawDescGZIP(), []int{0}
}

func (x *BlobsSidecar) GetBeaconBlockRoot() []byte {
	if x != nil {
		return x.BeaconBlockRoot
	}
	return nil
}

func (x *BlobsSidecar) GetBeaconBlockSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if x != nil {
		return x.BeaconBlockSlot
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(0)
}

func (x *BlobsSidecar) GetBlobs() []*v1.Blob {
	if x != nil {
		return x.Blobs
	}
	return nil
}

func (x *BlobsSidecar) GetAggregatedProof() []byte {
	if x != nil {
		return x.AggregatedProof
	}
	return nil
}

var File_proto_prysm_v1alpha1_blobs_proto protoreflect.FileDescriptor

var file_proto_prysm_v1alpha1_blobs_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0,
	0x02, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x12,
	0x32, 0x0a, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02,
	0x33, 0x32, 0x52, 0x0f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x71, 0x0a, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x45,
	0x82, 0xb5, 0x18, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79,
	0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x0f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x42,
	0x06, 0x92, 0xb5, 0x18, 0x02, 0x31, 0x36, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x31,
	0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x34, 0x38,
	0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x42, 0x95, 0x01, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x0a, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xca, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68,
	0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_prysm_v1alpha1_blobs_proto_rawDescOnce sync.Once
	file_proto_prysm_v1alpha1_blobs_proto_rawDescData = file_proto_prysm_v1alpha1_blobs_proto_rawDesc
)

func file_proto_prysm_v1alpha1_blobs_proto_rawDescGZIP() []byte {
	file_proto_prysm_v1alpha1_blobs_proto_rawDescOnce.Do(func() {
		file_proto_prysm_v1alpha1_blobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_prysm_v1alpha1_blobs_proto_rawDescData)
	})
	return file_proto_prysm_v1alpha1_blobs_proto_rawDescData
}

var file_proto_prysm_v1alpha1_blobs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_prysm_v1alpha1_blobs_proto_goTypes = []interface{}{
	(*BlobsSidecar)(nil), // 0: ethereum.eth.v1alpha1.BlobsSidecar
	(*v1.Blob)(nil),      // 1: ethereum.engine.v1.Blob
}
var file_proto_prysm_v1alpha1_blobs_proto_depIdxs = []int32{
	1, // 0: ethereum.eth.v1alpha1.BlobsSidecar.blobs:type_name -> ethereum.engine.v1.Blob
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_prysm_v1alpha1_blobs_proto_init() }
func file_proto_prysm_v1alpha1_blobs_proto_init() {
	if File_proto_prysm_v1alpha1_blobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_prysm_v1alpha1_blobs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobsSidecar); i {
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
			RawDescriptor: file_proto_prysm_v1alpha1_blobs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_prysm_v1alpha1_blobs_proto_goTypes,
		DependencyIndexes: file_proto_prysm_v1alpha1_blobs_proto_depIdxs,
		MessageInfos:      file_proto_prysm_v1alpha1_blobs_proto_msgTypes,
	}.Build()
	File_proto_prysm_v1alpha1_blobs_proto = out.File
	file_proto_prysm_v1alpha1_blobs_proto_rawDesc = nil
	file_proto_prysm_v1alpha1_blobs_proto_goTypes = nil
	file_proto_prysm_v1alpha1_blobs_proto_depIdxs = nil
}
