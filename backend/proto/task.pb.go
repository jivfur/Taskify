// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: backend/proto/task.proto

package taskify

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

// The Task message represents a task entity.
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId       int64  `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"` // Unique identifier for the task
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`   // Detailed description of the task
	Deadline     int64  `protobuf:"varint,4,opt,name=deadline,proto3" json:"deadline,omitempty"`        // Deadline timestamp for the task
	ExitCriteria string `protobuf:"bytes,5,opt,name=exitCriteria,proto3" json:"exitCriteria,omitempty"` // Exit criteria for completing the task
	Complete     bool   `protobuf:"varint,6,opt,name=complete,proto3" json:"complete,omitempty"`        // Status of task completion
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_backend_proto_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_backend_proto_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *Task) GetExitCriteria() string {
	if x != nil {
		return x.ExitCriteria
	}
	return ""
}

func (x *Task) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

// Request and Response messages
type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"` // The task to create or update
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	mi := &file_backend_proto_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_backend_proto_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"` // The task response after creation or update
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	mi := &file_backend_proto_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_backend_proto_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"` // Updated task response
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	mi := &file_backend_proto_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_backend_proto_task_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Indicates if the task was successfully deleted
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	mi := &file_backend_proto_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_backend_proto_task_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"` // List of tasks returned
}

func (x *ListTaskResponse) Reset() {
	*x = ListTaskResponse{}
	mi := &file_backend_proto_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskResponse) ProtoMessage() {}

func (x *ListTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskResponse.ProtoReflect.Descriptor instead.
func (*ListTaskResponse) Descriptor() ([]byte, []int) {
	return file_backend_proto_task_proto_rawDescGZIP(), []int{5}
}

func (x *ListTaskResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_backend_proto_task_proto protoreflect.FileDescriptor

var file_backend_proto_task_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x69, 0x66, 0x79, 0x22, 0xb2, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x69, 0x74,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x78, 0x69, 0x74, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x30, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x66, 0x79, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x31, 0x0a, 0x0c, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69,
	0x66, 0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x37, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x66, 0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x69, 0x66, 0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x32,
	0x81, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x69, 0x66, 0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x66, 0x79, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69,
	0x66, 0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x66, 0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x66, 0x79, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x69, 0x66, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x66, 0x79, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x69,
	0x66, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x66, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_proto_task_proto_rawDescOnce sync.Once
	file_backend_proto_task_proto_rawDescData = file_backend_proto_task_proto_rawDesc
)

func file_backend_proto_task_proto_rawDescGZIP() []byte {
	file_backend_proto_task_proto_rawDescOnce.Do(func() {
		file_backend_proto_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_proto_task_proto_rawDescData)
	})
	return file_backend_proto_task_proto_rawDescData
}

var file_backend_proto_task_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_backend_proto_task_proto_goTypes = []any{
	(*Task)(nil),               // 0: taskify.Task
	(*TaskRequest)(nil),        // 1: taskify.TaskRequest
	(*TaskResponse)(nil),       // 2: taskify.TaskResponse
	(*UpdateTaskResponse)(nil), // 3: taskify.UpdateTaskResponse
	(*DeleteTaskResponse)(nil), // 4: taskify.DeleteTaskResponse
	(*ListTaskResponse)(nil),   // 5: taskify.ListTaskResponse
}
var file_backend_proto_task_proto_depIdxs = []int32{
	0, // 0: taskify.TaskRequest.task:type_name -> taskify.Task
	0, // 1: taskify.TaskResponse.task:type_name -> taskify.Task
	0, // 2: taskify.UpdateTaskResponse.task:type_name -> taskify.Task
	0, // 3: taskify.ListTaskResponse.tasks:type_name -> taskify.Task
	1, // 4: taskify.TaskService.CreateTask:input_type -> taskify.TaskRequest
	1, // 5: taskify.TaskService.UpdateTask:input_type -> taskify.TaskRequest
	1, // 6: taskify.TaskService.DeleteTask:input_type -> taskify.TaskRequest
	1, // 7: taskify.TaskService.ListTask:input_type -> taskify.TaskRequest
	2, // 8: taskify.TaskService.CreateTask:output_type -> taskify.TaskResponse
	2, // 9: taskify.TaskService.UpdateTask:output_type -> taskify.TaskResponse
	4, // 10: taskify.TaskService.DeleteTask:output_type -> taskify.DeleteTaskResponse
	5, // 11: taskify.TaskService.ListTask:output_type -> taskify.ListTaskResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_backend_proto_task_proto_init() }
func file_backend_proto_task_proto_init() {
	if File_backend_proto_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_proto_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_proto_task_proto_goTypes,
		DependencyIndexes: file_backend_proto_task_proto_depIdxs,
		MessageInfos:      file_backend_proto_task_proto_msgTypes,
	}.Build()
	File_backend_proto_task_proto = out.File
	file_backend_proto_task_proto_rawDesc = nil
	file_backend_proto_task_proto_goTypes = nil
	file_backend_proto_task_proto_depIdxs = nil
}
