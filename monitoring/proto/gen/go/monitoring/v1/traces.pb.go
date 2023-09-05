// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: monitoring/v1/traces.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Span struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId      string                 `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId       string                 `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	ParentSpanId string                 `protobuf:"bytes,3,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	Name         string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	StartTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *Span) Reset() {
	*x = Span{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_v1_traces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_v1_traces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_monitoring_v1_traces_proto_rawDescGZIP(), []int{0}
}

func (x *Span) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *Span) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *Span) GetParentSpanId() string {
	if x != nil {
		return x.ParentSpanId
	}
	return ""
}

func (x *Span) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Span) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

type CreateSpanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId      string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId       string `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	ParentSpanId string `protobuf:"bytes,3,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateSpanRequest) Reset() {
	*x = CreateSpanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_v1_traces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSpanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSpanRequest) ProtoMessage() {}

func (x *CreateSpanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_v1_traces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSpanRequest.ProtoReflect.Descriptor instead.
func (*CreateSpanRequest) Descriptor() ([]byte, []int) {
	return file_monitoring_v1_traces_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSpanRequest) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *CreateSpanRequest) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *CreateSpanRequest) GetParentSpanId() string {
	if x != nil {
		return x.ParentSpanId
	}
	return ""
}

func (x *CreateSpanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateSpanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId      string                 `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId       string                 `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	ParentSpanId string                 `protobuf:"bytes,3,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	Name         string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	StartTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *CreateSpanResponse) Reset() {
	*x = CreateSpanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_v1_traces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSpanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSpanResponse) ProtoMessage() {}

func (x *CreateSpanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_v1_traces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSpanResponse.ProtoReflect.Descriptor instead.
func (*CreateSpanResponse) Descriptor() ([]byte, []int) {
	return file_monitoring_v1_traces_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSpanResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *CreateSpanResponse) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *CreateSpanResponse) GetParentSpanId() string {
	if x != nil {
		return x.ParentSpanId
	}
	return ""
}

func (x *CreateSpanResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSpanResponse) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

type GetTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *GetTraceRequest) Reset() {
	*x = GetTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_v1_traces_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTraceRequest) ProtoMessage() {}

func (x *GetTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_v1_traces_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTraceRequest.ProtoReflect.Descriptor instead.
func (*GetTraceRequest) Descriptor() ([]byte, []int) {
	return file_monitoring_v1_traces_proto_rawDescGZIP(), []int{3}
}

func (x *GetTraceRequest) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

type GetTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId string  `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Spans   []*Span `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
}

func (x *GetTraceResponse) Reset() {
	*x = GetTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_v1_traces_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTraceResponse) ProtoMessage() {}

func (x *GetTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_v1_traces_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTraceResponse.ProtoReflect.Descriptor instead.
func (*GetTraceResponse) Descriptor() ([]byte, []int) {
	return file_monitoring_v1_traces_proto_rawDescGZIP(), []int{4}
}

func (x *GetTraceResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *GetTraceResponse) GetSpans() []*Span {
	if x != nil {
		return x.Spans
	}
	return nil
}

type Traces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceIds []string `protobuf:"bytes,1,rep,name=trace_ids,json=traceIds,proto3" json:"trace_ids,omitempty"`
}

func (x *Traces) Reset() {
	*x = Traces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_v1_traces_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Traces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Traces) ProtoMessage() {}

func (x *Traces) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_v1_traces_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Traces.ProtoReflect.Descriptor instead.
func (*Traces) Descriptor() ([]byte, []int) {
	return file_monitoring_v1_traces_proto_rawDescGZIP(), []int{5}
}

func (x *Traces) GetTraceIds() []string {
	if x != nil {
		return x.TraceIds
	}
	return nil
}

var File_monitoring_v1_traces_proto protoreflect.FileDescriptor

var file_monitoring_v1_traces_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x04, 0x53, 0x70,
	0x61, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xbd, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x68, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x05,
	0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e,
	0x52, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x22, 0x25, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x32, 0xc4,
	0x02, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x73, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x30,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x2e, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x22, 0x00, 0x42, 0xfd, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63,
	0x72, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x53, 0x4d, 0x54, 0xaa, 0x02, 0x1d, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63,
	0x72, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1d, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63,
	0x72, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x29, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63,
	0x72, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x20, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x63, 0x72, 0x3a, 0x3a, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitoring_v1_traces_proto_rawDescOnce sync.Once
	file_monitoring_v1_traces_proto_rawDescData = file_monitoring_v1_traces_proto_rawDesc
)

func file_monitoring_v1_traces_proto_rawDescGZIP() []byte {
	file_monitoring_v1_traces_proto_rawDescOnce.Do(func() {
		file_monitoring_v1_traces_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitoring_v1_traces_proto_rawDescData)
	})
	return file_monitoring_v1_traces_proto_rawDescData
}

var file_monitoring_v1_traces_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_monitoring_v1_traces_proto_goTypes = []interface{}{
	(*Span)(nil),                  // 0: smartpcr.monitoring.traces.v1.Span
	(*CreateSpanRequest)(nil),     // 1: smartpcr.monitoring.traces.v1.CreateSpanRequest
	(*CreateSpanResponse)(nil),    // 2: smartpcr.monitoring.traces.v1.CreateSpanResponse
	(*GetTraceRequest)(nil),       // 3: smartpcr.monitoring.traces.v1.GetTraceRequest
	(*GetTraceResponse)(nil),      // 4: smartpcr.monitoring.traces.v1.GetTraceResponse
	(*Traces)(nil),                // 5: smartpcr.monitoring.traces.v1.Traces
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_monitoring_v1_traces_proto_depIdxs = []int32{
	6, // 0: smartpcr.monitoring.traces.v1.Span.start_time:type_name -> google.protobuf.Timestamp
	6, // 1: smartpcr.monitoring.traces.v1.CreateSpanResponse.start_time:type_name -> google.protobuf.Timestamp
	0, // 2: smartpcr.monitoring.traces.v1.GetTraceResponse.spans:type_name -> smartpcr.monitoring.traces.v1.Span
	1, // 3: smartpcr.monitoring.traces.v1.TracesService.CreateSpan:input_type -> smartpcr.monitoring.traces.v1.CreateSpanRequest
	3, // 4: smartpcr.monitoring.traces.v1.TracesService.GetTrace:input_type -> smartpcr.monitoring.traces.v1.GetTraceRequest
	7, // 5: smartpcr.monitoring.traces.v1.TracesService.GetAllTraces:input_type -> google.protobuf.Empty
	2, // 6: smartpcr.monitoring.traces.v1.TracesService.CreateSpan:output_type -> smartpcr.monitoring.traces.v1.CreateSpanResponse
	4, // 7: smartpcr.monitoring.traces.v1.TracesService.GetTrace:output_type -> smartpcr.monitoring.traces.v1.GetTraceResponse
	5, // 8: smartpcr.monitoring.traces.v1.TracesService.GetAllTraces:output_type -> smartpcr.monitoring.traces.v1.Traces
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_monitoring_v1_traces_proto_init() }
func file_monitoring_v1_traces_proto_init() {
	if File_monitoring_v1_traces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitoring_v1_traces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Span); i {
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
		file_monitoring_v1_traces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSpanRequest); i {
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
		file_monitoring_v1_traces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSpanResponse); i {
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
		file_monitoring_v1_traces_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTraceRequest); i {
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
		file_monitoring_v1_traces_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTraceResponse); i {
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
		file_monitoring_v1_traces_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Traces); i {
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
			RawDescriptor: file_monitoring_v1_traces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitoring_v1_traces_proto_goTypes,
		DependencyIndexes: file_monitoring_v1_traces_proto_depIdxs,
		MessageInfos:      file_monitoring_v1_traces_proto_msgTypes,
	}.Build()
	File_monitoring_v1_traces_proto = out.File
	file_monitoring_v1_traces_proto_rawDesc = nil
	file_monitoring_v1_traces_proto_goTypes = nil
	file_monitoring_v1_traces_proto_depIdxs = nil
}
