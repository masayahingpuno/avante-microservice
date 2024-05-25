// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: assignment/v1/assignment.proto

package v1

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

// Define the Assignment message
type Assignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DueDate     string `protobuf:"bytes,4,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"` // ISO 8601 date format
}

func (x *Assignment) Reset() {
	*x = Assignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assignment) ProtoMessage() {}

func (x *Assignment) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assignment.ProtoReflect.Descriptor instead.
func (*Assignment) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{0}
}

func (x *Assignment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Assignment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Assignment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Assignment) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

// Request message for creating an assignment
type CreateAssignmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DueDate     string `protobuf:"bytes,3,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"` // ISO 8601 date format
}

func (x *CreateAssignmentRequest) Reset() {
	*x = CreateAssignmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssignmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssignmentRequest) ProtoMessage() {}

func (x *CreateAssignmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssignmentRequest.ProtoReflect.Descriptor instead.
func (*CreateAssignmentRequest) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAssignmentRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateAssignmentRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateAssignmentRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

// Response message for creating an assignment
type CreateAssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *CreateAssignmentResponse) Reset() {
	*x = CreateAssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssignmentResponse) ProtoMessage() {}

func (x *CreateAssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssignmentResponse.ProtoReflect.Descriptor instead.
func (*CreateAssignmentResponse) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAssignmentResponse) GetAssignment() *Assignment {
	if x != nil {
		return x.Assignment
	}
	return nil
}

// Request message for getting an assignment
type GetAssignmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAssignmentRequest) Reset() {
	*x = GetAssignmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssignmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssignmentRequest) ProtoMessage() {}

func (x *GetAssignmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssignmentRequest.ProtoReflect.Descriptor instead.
func (*GetAssignmentRequest) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssignmentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response message for getting an assignment
type GetAssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *GetAssignmentResponse) Reset() {
	*x = GetAssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssignmentResponse) ProtoMessage() {}

func (x *GetAssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssignmentResponse.ProtoReflect.Descriptor instead.
func (*GetAssignmentResponse) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{4}
}

func (x *GetAssignmentResponse) GetAssignment() *Assignment {
	if x != nil {
		return x.Assignment
	}
	return nil
}

// Request message for listing assignments
type ListAssignmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAssignmentsRequest) Reset() {
	*x = ListAssignmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssignmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssignmentsRequest) ProtoMessage() {}

func (x *ListAssignmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssignmentsRequest.ProtoReflect.Descriptor instead.
func (*ListAssignmentsRequest) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{5}
}

// Response message for listing assignments
type ListAssignmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assignments []*Assignment `protobuf:"bytes,1,rep,name=assignments,proto3" json:"assignments,omitempty"`
}

func (x *ListAssignmentsResponse) Reset() {
	*x = ListAssignmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssignmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssignmentsResponse) ProtoMessage() {}

func (x *ListAssignmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssignmentsResponse.ProtoReflect.Descriptor instead.
func (*ListAssignmentsResponse) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{6}
}

func (x *ListAssignmentsResponse) GetAssignments() []*Assignment {
	if x != nil {
		return x.Assignments
	}
	return nil
}

// Request message for updating an assignment
type UpdateAssignmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DueDate     string `protobuf:"bytes,4,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"` // ISO 8601 date format
}

func (x *UpdateAssignmentRequest) Reset() {
	*x = UpdateAssignmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAssignmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssignmentRequest) ProtoMessage() {}

func (x *UpdateAssignmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssignmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateAssignmentRequest) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAssignmentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAssignmentRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateAssignmentRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateAssignmentRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

// Response message for updating an assignment
type UpdateAssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *UpdateAssignmentResponse) Reset() {
	*x = UpdateAssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssignmentResponse) ProtoMessage() {}

func (x *UpdateAssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssignmentResponse.ProtoReflect.Descriptor instead.
func (*UpdateAssignmentResponse) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateAssignmentResponse) GetAssignment() *Assignment {
	if x != nil {
		return x.Assignment
	}
	return nil
}

// Request message for deleting an assignment
type DeleteAssignmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAssignmentRequest) Reset() {
	*x = DeleteAssignmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAssignmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssignmentRequest) ProtoMessage() {}

func (x *DeleteAssignmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssignmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteAssignmentRequest) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAssignmentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response message for deleting an assignment
type DeleteAssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteAssignmentResponse) Reset() {
	*x = DeleteAssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assignment_v1_assignment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssignmentResponse) ProtoMessage() {}

func (x *DeleteAssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assignment_v1_assignment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssignmentResponse.ProtoReflect.Descriptor instead.
func (*DeleteAssignmentResponse) Descriptor() ([]byte, []int) {
	return file_assignment_v1_assignment_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAssignmentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_assignment_v1_assignment_proto protoreflect.FileDescriptor

var file_assignment_v1_assignment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22,
	0x6f, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x22, 0x6c, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x55,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x7c, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x55, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x80, 0x04, 0x0a, 0x11, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x63, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x60, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x70, 0x70, 0x79,
	0x74, 0x72, 0x65, 0x65, 0x30, 0x39, 0x2f, 0x61, 0x76, 0x61, 0x6e, 0x74, 0x65, 0x2d, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_assignment_v1_assignment_proto_rawDescOnce sync.Once
	file_assignment_v1_assignment_proto_rawDescData = file_assignment_v1_assignment_proto_rawDesc
)

func file_assignment_v1_assignment_proto_rawDescGZIP() []byte {
	file_assignment_v1_assignment_proto_rawDescOnce.Do(func() {
		file_assignment_v1_assignment_proto_rawDescData = protoimpl.X.CompressGZIP(file_assignment_v1_assignment_proto_rawDescData)
	})
	return file_assignment_v1_assignment_proto_rawDescData
}

var file_assignment_v1_assignment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_assignment_v1_assignment_proto_goTypes = []interface{}{
	(*Assignment)(nil),               // 0: assignment.v1.Assignment
	(*CreateAssignmentRequest)(nil),  // 1: assignment.v1.CreateAssignmentRequest
	(*CreateAssignmentResponse)(nil), // 2: assignment.v1.CreateAssignmentResponse
	(*GetAssignmentRequest)(nil),     // 3: assignment.v1.GetAssignmentRequest
	(*GetAssignmentResponse)(nil),    // 4: assignment.v1.GetAssignmentResponse
	(*ListAssignmentsRequest)(nil),   // 5: assignment.v1.ListAssignmentsRequest
	(*ListAssignmentsResponse)(nil),  // 6: assignment.v1.ListAssignmentsResponse
	(*UpdateAssignmentRequest)(nil),  // 7: assignment.v1.UpdateAssignmentRequest
	(*UpdateAssignmentResponse)(nil), // 8: assignment.v1.UpdateAssignmentResponse
	(*DeleteAssignmentRequest)(nil),  // 9: assignment.v1.DeleteAssignmentRequest
	(*DeleteAssignmentResponse)(nil), // 10: assignment.v1.DeleteAssignmentResponse
}
var file_assignment_v1_assignment_proto_depIdxs = []int32{
	0,  // 0: assignment.v1.CreateAssignmentResponse.assignment:type_name -> assignment.v1.Assignment
	0,  // 1: assignment.v1.GetAssignmentResponse.assignment:type_name -> assignment.v1.Assignment
	0,  // 2: assignment.v1.ListAssignmentsResponse.assignments:type_name -> assignment.v1.Assignment
	0,  // 3: assignment.v1.UpdateAssignmentResponse.assignment:type_name -> assignment.v1.Assignment
	1,  // 4: assignment.v1.AssignmentService.CreateAssignment:input_type -> assignment.v1.CreateAssignmentRequest
	3,  // 5: assignment.v1.AssignmentService.GetAssignment:input_type -> assignment.v1.GetAssignmentRequest
	5,  // 6: assignment.v1.AssignmentService.ListAssignments:input_type -> assignment.v1.ListAssignmentsRequest
	7,  // 7: assignment.v1.AssignmentService.UpdateAssignment:input_type -> assignment.v1.UpdateAssignmentRequest
	9,  // 8: assignment.v1.AssignmentService.DeleteAssignment:input_type -> assignment.v1.DeleteAssignmentRequest
	2,  // 9: assignment.v1.AssignmentService.CreateAssignment:output_type -> assignment.v1.CreateAssignmentResponse
	4,  // 10: assignment.v1.AssignmentService.GetAssignment:output_type -> assignment.v1.GetAssignmentResponse
	6,  // 11: assignment.v1.AssignmentService.ListAssignments:output_type -> assignment.v1.ListAssignmentsResponse
	8,  // 12: assignment.v1.AssignmentService.UpdateAssignment:output_type -> assignment.v1.UpdateAssignmentResponse
	10, // 13: assignment.v1.AssignmentService.DeleteAssignment:output_type -> assignment.v1.DeleteAssignmentResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_assignment_v1_assignment_proto_init() }
func file_assignment_v1_assignment_proto_init() {
	if File_assignment_v1_assignment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_assignment_v1_assignment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Assignment); i {
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
		file_assignment_v1_assignment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAssignmentRequest); i {
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
		file_assignment_v1_assignment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAssignmentResponse); i {
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
		file_assignment_v1_assignment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssignmentRequest); i {
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
		file_assignment_v1_assignment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssignmentResponse); i {
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
		file_assignment_v1_assignment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssignmentsRequest); i {
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
		file_assignment_v1_assignment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssignmentsResponse); i {
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
		file_assignment_v1_assignment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAssignmentRequest); i {
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
		file_assignment_v1_assignment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAssignmentResponse); i {
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
		file_assignment_v1_assignment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAssignmentRequest); i {
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
		file_assignment_v1_assignment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAssignmentResponse); i {
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
			RawDescriptor: file_assignment_v1_assignment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_assignment_v1_assignment_proto_goTypes,
		DependencyIndexes: file_assignment_v1_assignment_proto_depIdxs,
		MessageInfos:      file_assignment_v1_assignment_proto_msgTypes,
	}.Build()
	File_assignment_v1_assignment_proto = out.File
	file_assignment_v1_assignment_proto_rawDesc = nil
	file_assignment_v1_assignment_proto_goTypes = nil
	file_assignment_v1_assignment_proto_depIdxs = nil
}
