//
//Copyright The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file was autogenerated by go-to-protobuf. Do not edit it manually!

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: k8s.io/api/events/v1beta1/generated.proto

package v1beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v11 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
// Events have a limited retention time and triggers and messages may evolve
// with time.  Event consumers should not rely on the timing of an event
// with a given Reason reflecting a consistent underlying trigger, or the
// continued existence of events with that Reason.  Events should be
// treated as informative, best-effort, supplemental data.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// eventTime is the time when this Event was first observed. It is required.
	EventTime *v1.MicroTime `protobuf:"bytes,2,opt,name=eventTime" json:"eventTime,omitempty"`
	// series is data about the Event series this event represents or nil if it's a singleton Event.
	// +optional
	Series *EventSeries `protobuf:"bytes,3,opt,name=series" json:"series,omitempty"`
	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	// This field cannot be empty for new Events.
	// +optional
	ReportingController *string `protobuf:"bytes,4,opt,name=reportingController" json:"reportingController,omitempty"`
	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`.
	// This field cannot be empty for new Events and it can have at most 128 characters.
	// +optional
	ReportingInstance *string `protobuf:"bytes,5,opt,name=reportingInstance" json:"reportingInstance,omitempty"`
	// action is what action was taken/failed regarding to the regarding object. It is machine-readable.
	// This field can have at most 128 characters.
	// +optional
	Action *string `protobuf:"bytes,6,opt,name=action" json:"action,omitempty"`
	// reason is why the action was taken. It is human-readable.
	// This field can have at most 128 characters.
	// +optional
	Reason *string `protobuf:"bytes,7,opt,name=reason" json:"reason,omitempty"`
	// regarding contains the object this Event is about. In most cases it's an Object reporting controller
	// implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because
	// it acts on some changes in a ReplicaSet object.
	// +optional
	Regarding *v11.ObjectReference `protobuf:"bytes,8,opt,name=regarding" json:"regarding,omitempty"`
	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers
	// a creation or deletion of related object.
	// +optional
	Related *v11.ObjectReference `protobuf:"bytes,9,opt,name=related" json:"related,omitempty"`
	// note is a human-readable description of the status of this operation.
	// Maximal length of the note is 1kB, but libraries should be prepared to
	// handle values up to 64kB.
	// +optional
	Note *string `protobuf:"bytes,10,opt,name=note" json:"note,omitempty"`
	// type is the type of this event (Normal, Warning), new types could be added in the future.
	// It is machine-readable.
	// +optional
	Type *string `protobuf:"bytes,11,opt,name=type" json:"type,omitempty"`
	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	// +optional
	DeprecatedSource *v11.EventSource `protobuf:"bytes,12,opt,name=deprecatedSource" json:"deprecatedSource,omitempty"`
	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	// +optional
	DeprecatedFirstTimestamp *v1.Time `protobuf:"bytes,13,opt,name=deprecatedFirstTimestamp" json:"deprecatedFirstTimestamp,omitempty"`
	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	// +optional
	DeprecatedLastTimestamp *v1.Time `protobuf:"bytes,14,opt,name=deprecatedLastTimestamp" json:"deprecatedLastTimestamp,omitempty"`
	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	// +optional
	DeprecatedCount *int32 `protobuf:"varint,15,opt,name=deprecatedCount" json:"deprecatedCount,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_events_v1beta1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Event) GetEventTime() *v1.MicroTime {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *Event) GetSeries() *EventSeries {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *Event) GetReportingController() string {
	if x != nil && x.ReportingController != nil {
		return *x.ReportingController
	}
	return ""
}

func (x *Event) GetReportingInstance() string {
	if x != nil && x.ReportingInstance != nil {
		return *x.ReportingInstance
	}
	return ""
}

func (x *Event) GetAction() string {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return ""
}

func (x *Event) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *Event) GetRegarding() *v11.ObjectReference {
	if x != nil {
		return x.Regarding
	}
	return nil
}

func (x *Event) GetRelated() *v11.ObjectReference {
	if x != nil {
		return x.Related
	}
	return nil
}

func (x *Event) GetNote() string {
	if x != nil && x.Note != nil {
		return *x.Note
	}
	return ""
}

func (x *Event) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *Event) GetDeprecatedSource() *v11.EventSource {
	if x != nil {
		return x.DeprecatedSource
	}
	return nil
}

func (x *Event) GetDeprecatedFirstTimestamp() *v1.Time {
	if x != nil {
		return x.DeprecatedFirstTimestamp
	}
	return nil
}

func (x *Event) GetDeprecatedLastTimestamp() *v1.Time {
	if x != nil {
		return x.DeprecatedLastTimestamp
	}
	return nil
}

func (x *Event) GetDeprecatedCount() int32 {
	if x != nil && x.DeprecatedCount != nil {
		return *x.DeprecatedCount
	}
	return 0
}

// EventList is a list of Event objects.
type EventList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// items is a list of schema objects.
	Items []*Event `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (x *EventList) Reset() {
	*x = EventList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventList) ProtoMessage() {}

func (x *EventList) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventList.ProtoReflect.Descriptor instead.
func (*EventList) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_events_v1beta1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *EventList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EventList) GetItems() []*Event {
	if x != nil {
		return x.Items
	}
	return nil
}

// EventSeries contain information on series of events, i.e. thing that was/is happening
// continuously for some time.
type EventSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// count is the number of occurrences in this series up to the last heartbeat time.
	Count *int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
	LastObservedTime *v1.MicroTime `protobuf:"bytes,2,opt,name=lastObservedTime" json:"lastObservedTime,omitempty"`
}

func (x *EventSeries) Reset() {
	*x = EventSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSeries) ProtoMessage() {}

func (x *EventSeries) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSeries.ProtoReflect.Descriptor instead.
func (*EventSeries) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_events_v1beta1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *EventSeries) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *EventSeries) GetLastObservedTime() *v1.MicroTime {
	if x != nil {
		return x.LastObservedTime
	}
	return nil
}

var File_k8s_io_api_events_v1beta1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_api_events_v1beta1_generated_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6b, 0x38, 0x73,
	0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x22, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x36, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x06, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x4d, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3e, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x41, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x67, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x64, 0x65,
	0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x10, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x18, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x72, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x18, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x72, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x64, 0x0a, 0x17, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x17, 0x64, 0x65,
	0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x8f, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31,
}

var (
	file_k8s_io_api_events_v1beta1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_api_events_v1beta1_generated_proto_rawDescData = file_k8s_io_api_events_v1beta1_generated_proto_rawDesc
)

func file_k8s_io_api_events_v1beta1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_api_events_v1beta1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_api_events_v1beta1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_api_events_v1beta1_generated_proto_rawDescData)
	})
	return file_k8s_io_api_events_v1beta1_generated_proto_rawDescData
}

var file_k8s_io_api_events_v1beta1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_k8s_io_api_events_v1beta1_generated_proto_goTypes = []interface{}{
	(*Event)(nil),               // 0: k8s.io.api.events.v1beta1.Event
	(*EventList)(nil),           // 1: k8s.io.api.events.v1beta1.EventList
	(*EventSeries)(nil),         // 2: k8s.io.api.events.v1beta1.EventSeries
	(*v1.ObjectMeta)(nil),       // 3: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.MicroTime)(nil),        // 4: k8s.io.apimachinery.pkg.apis.meta.v1.MicroTime
	(*v11.ObjectReference)(nil), // 5: k8s.io.api.core.v1.ObjectReference
	(*v11.EventSource)(nil),     // 6: k8s.io.api.core.v1.EventSource
	(*v1.Time)(nil),             // 7: k8s.io.apimachinery.pkg.apis.meta.v1.Time
	(*v1.ListMeta)(nil),         // 8: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
}
var file_k8s_io_api_events_v1beta1_generated_proto_depIdxs = []int32{
	3,  // 0: k8s.io.api.events.v1beta1.Event.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	4,  // 1: k8s.io.api.events.v1beta1.Event.eventTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.MicroTime
	2,  // 2: k8s.io.api.events.v1beta1.Event.series:type_name -> k8s.io.api.events.v1beta1.EventSeries
	5,  // 3: k8s.io.api.events.v1beta1.Event.regarding:type_name -> k8s.io.api.core.v1.ObjectReference
	5,  // 4: k8s.io.api.events.v1beta1.Event.related:type_name -> k8s.io.api.core.v1.ObjectReference
	6,  // 5: k8s.io.api.events.v1beta1.Event.deprecatedSource:type_name -> k8s.io.api.core.v1.EventSource
	7,  // 6: k8s.io.api.events.v1beta1.Event.deprecatedFirstTimestamp:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	7,  // 7: k8s.io.api.events.v1beta1.Event.deprecatedLastTimestamp:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	8,  // 8: k8s.io.api.events.v1beta1.EventList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	0,  // 9: k8s.io.api.events.v1beta1.EventList.items:type_name -> k8s.io.api.events.v1beta1.Event
	4,  // 10: k8s.io.api.events.v1beta1.EventSeries.lastObservedTime:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.MicroTime
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_k8s_io_api_events_v1beta1_generated_proto_init() }
func file_k8s_io_api_events_v1beta1_generated_proto_init() {
	if File_k8s_io_api_events_v1beta1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventList); i {
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
		file_k8s_io_api_events_v1beta1_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventSeries); i {
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
			RawDescriptor: file_k8s_io_api_events_v1beta1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_api_events_v1beta1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_api_events_v1beta1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_api_events_v1beta1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_api_events_v1beta1_generated_proto = out.File
	file_k8s_io_api_events_v1beta1_generated_proto_rawDesc = nil
	file_k8s_io_api_events_v1beta1_generated_proto_goTypes = nil
	file_k8s_io_api_events_v1beta1_generated_proto_depIdxs = nil
}
