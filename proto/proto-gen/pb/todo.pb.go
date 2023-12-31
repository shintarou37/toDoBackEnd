// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: proto/protofiles/todo.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Context  string `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Priority int32  `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *Todo) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type TodoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *TodoListRequest) Reset() {
	*x = TodoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListRequest) ProtoMessage() {}

func (x *TodoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListRequest.ProtoReflect.Descriptor instead.
func (*TodoListRequest) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TodoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo []*Todo `protobuf:"bytes,1,rep,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoListResponse) Reset() {
	*x = TodoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListResponse) ProtoMessage() {}

func (x *TodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListResponse.ProtoReflect.Descriptor instead.
func (*TodoListResponse) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodoListResponse) GetTodo() []*Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type TodoGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *TodoGetRequest) Reset() {
	*x = TodoGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoGetRequest) ProtoMessage() {}

func (x *TodoGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoGetRequest.ProtoReflect.Descriptor instead.
func (*TodoGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoGetRequest) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type TodoGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoGetResponse) Reset() {
	*x = TodoGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoGetResponse) ProtoMessage() {}

func (x *TodoGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoGetResponse.ProtoReflect.Descriptor instead.
func (*TodoGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{4}
}

func (x *TodoGetResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type TodoCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo   *Todo  `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *TodoCreateRequest) Reset() {
	*x = TodoCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoCreateRequest) ProtoMessage() {}

func (x *TodoCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoCreateRequest.ProtoReflect.Descriptor instead.
func (*TodoCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{5}
}

func (x *TodoCreateRequest) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

func (x *TodoCreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TodoCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoCreateResponse) Reset() {
	*x = TodoCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoCreateResponse) ProtoMessage() {}

func (x *TodoCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoCreateResponse.ProtoReflect.Descriptor instead.
func (*TodoCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{6}
}

func (x *TodoCreateResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type TodoUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo   *Todo  `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	TodoId string `protobuf:"bytes,2,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *TodoUpdateRequest) Reset() {
	*x = TodoUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoUpdateRequest) ProtoMessage() {}

func (x *TodoUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoUpdateRequest.ProtoReflect.Descriptor instead.
func (*TodoUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{7}
}

func (x *TodoUpdateRequest) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

func (x *TodoUpdateRequest) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type TodoUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoUpdateResponse) Reset() {
	*x = TodoUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protofiles_todo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoUpdateResponse) ProtoMessage() {}

func (x *TodoUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protofiles_todo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoUpdateResponse.ProtoReflect.Descriptor instead.
func (*TodoUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_protofiles_todo_proto_rawDescGZIP(), []int{8}
}

func (x *TodoUpdateResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

var File_proto_protofiles_todo_proto protoreflect.FileDescriptor

var file_proto_protofiles_todo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d,
	0x61, 0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x52, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x2a, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x32, 0x0a, 0x10, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x29, 0x0a, 0x0e, 0x54, 0x6f, 0x64, 0x6f, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x64,
	0x22, 0x31, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x22, 0x4c, 0x0a, 0x11, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x34, 0x0a, 0x12, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x4c, 0x0a, 0x11, 0x54, 0x6f, 0x64, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x6f, 0x64, 0x6f, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x12, 0x54, 0x6f, 0x64, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x32, 0xd1, 0x02, 0x0a, 0x0b,
	0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x4c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01,
	0x2a, 0x22, 0x0c, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_protofiles_todo_proto_rawDescOnce sync.Once
	file_proto_protofiles_todo_proto_rawDescData = file_proto_protofiles_todo_proto_rawDesc
)

func file_proto_protofiles_todo_proto_rawDescGZIP() []byte {
	file_proto_protofiles_todo_proto_rawDescOnce.Do(func() {
		file_proto_protofiles_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protofiles_todo_proto_rawDescData)
	})
	return file_proto_protofiles_todo_proto_rawDescData
}

var file_proto_protofiles_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_protofiles_todo_proto_goTypes = []interface{}{
	(*Todo)(nil),               // 0: main.Todo
	(*TodoListRequest)(nil),    // 1: main.TodoListRequest
	(*TodoListResponse)(nil),   // 2: main.TodoListResponse
	(*TodoGetRequest)(nil),     // 3: main.TodoGetRequest
	(*TodoGetResponse)(nil),    // 4: main.TodoGetResponse
	(*TodoCreateRequest)(nil),  // 5: main.TodoCreateRequest
	(*TodoCreateResponse)(nil), // 6: main.TodoCreateResponse
	(*TodoUpdateRequest)(nil),  // 7: main.TodoUpdateRequest
	(*TodoUpdateResponse)(nil), // 8: main.TodoUpdateResponse
}
var file_proto_protofiles_todo_proto_depIdxs = []int32{
	0,  // 0: main.TodoListResponse.todo:type_name -> main.Todo
	0,  // 1: main.TodoGetResponse.todo:type_name -> main.Todo
	0,  // 2: main.TodoCreateRequest.todo:type_name -> main.Todo
	0,  // 3: main.TodoCreateResponse.todo:type_name -> main.Todo
	0,  // 4: main.TodoUpdateRequest.todo:type_name -> main.Todo
	0,  // 5: main.TodoUpdateResponse.todo:type_name -> main.Todo
	3,  // 6: main.TodoService.Get:input_type -> main.TodoGetRequest
	1,  // 7: main.TodoService.List:input_type -> main.TodoListRequest
	5,  // 8: main.TodoService.Create:input_type -> main.TodoCreateRequest
	7,  // 9: main.TodoService.Update:input_type -> main.TodoUpdateRequest
	4,  // 10: main.TodoService.Get:output_type -> main.TodoGetResponse
	2,  // 11: main.TodoService.List:output_type -> main.TodoListResponse
	6,  // 12: main.TodoService.Create:output_type -> main.TodoCreateResponse
	8,  // 13: main.TodoService.Update:output_type -> main.TodoUpdateResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_protofiles_todo_proto_init() }
func file_proto_protofiles_todo_proto_init() {
	if File_proto_protofiles_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protofiles_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_proto_protofiles_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListRequest); i {
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
		file_proto_protofiles_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListResponse); i {
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
		file_proto_protofiles_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoGetRequest); i {
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
		file_proto_protofiles_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoGetResponse); i {
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
		file_proto_protofiles_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoCreateRequest); i {
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
		file_proto_protofiles_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoCreateResponse); i {
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
		file_proto_protofiles_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoUpdateRequest); i {
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
		file_proto_protofiles_todo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoUpdateResponse); i {
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
			RawDescriptor: file_proto_protofiles_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_protofiles_todo_proto_goTypes,
		DependencyIndexes: file_proto_protofiles_todo_proto_depIdxs,
		MessageInfos:      file_proto_protofiles_todo_proto_msgTypes,
	}.Build()
	File_proto_protofiles_todo_proto = out.File
	file_proto_protofiles_todo_proto_rawDesc = nil
	file_proto_protofiles_todo_proto_goTypes = nil
	file_proto_protofiles_todo_proto_depIdxs = nil
}
