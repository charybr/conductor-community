// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/task.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task_Status int32

const (
	Task_IN_PROGRESS                Task_Status = 0
	Task_CANCELED                   Task_Status = 1
	Task_FAILED                     Task_Status = 2
	Task_FAILED_WITH_TERMINAL_ERROR Task_Status = 3
	Task_COMPLETED                  Task_Status = 4
	Task_COMPLETED_WITH_ERRORS      Task_Status = 5
	Task_SCHEDULED                  Task_Status = 6
	Task_TIMED_OUT                  Task_Status = 7
	Task_READY_FOR_RERUN            Task_Status = 8
	Task_SKIPPED                    Task_Status = 9
)

var Task_Status_name = map[int32]string{
	0: "IN_PROGRESS",
	1: "CANCELED",
	2: "FAILED",
	3: "FAILED_WITH_TERMINAL_ERROR",
	4: "COMPLETED",
	5: "COMPLETED_WITH_ERRORS",
	6: "SCHEDULED",
	7: "TIMED_OUT",
	8: "READY_FOR_RERUN",
	9: "SKIPPED",
}
var Task_Status_value = map[string]int32{
	"IN_PROGRESS":                0,
	"CANCELED":                   1,
	"FAILED":                     2,
	"FAILED_WITH_TERMINAL_ERROR": 3,
	"COMPLETED":                  4,
	"COMPLETED_WITH_ERRORS":      5,
	"SCHEDULED":                  6,
	"TIMED_OUT":                  7,
	"READY_FOR_RERUN":            8,
	"SKIPPED":                    9,
}

func (x Task_Status) String() string {
	return proto.EnumName(Task_Status_name, int32(x))
}
func (Task_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_task_7843d3d2fd8c2dc8, []int{0, 0}
}

type Task struct {
	TaskType               string                    `protobuf:"bytes,1,opt,name=task_type,json=taskType" json:"task_type,omitempty"`
	Status                 Task_Status               `protobuf:"varint,2,opt,name=status,enum=conductor.proto.Task_Status" json:"status,omitempty"`
	InputData              map[string]*_struct.Value `protobuf:"bytes,3,rep,name=input_data,json=inputData" json:"input_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ReferenceTaskName      string                    `protobuf:"bytes,4,opt,name=reference_task_name,json=referenceTaskName" json:"reference_task_name,omitempty"`
	RetryCount             int32                     `protobuf:"varint,5,opt,name=retry_count,json=retryCount" json:"retry_count,omitempty"`
	Seq                    int32                     `protobuf:"varint,6,opt,name=seq" json:"seq,omitempty"`
	CorrelationId          string                    `protobuf:"bytes,7,opt,name=correlation_id,json=correlationId" json:"correlation_id,omitempty"`
	PollCount              int32                     `protobuf:"varint,8,opt,name=poll_count,json=pollCount" json:"poll_count,omitempty"`
	TaskDefName            string                    `protobuf:"bytes,9,opt,name=task_def_name,json=taskDefName" json:"task_def_name,omitempty"`
	ScheduledTime          int64                     `protobuf:"varint,10,opt,name=scheduled_time,json=scheduledTime" json:"scheduled_time,omitempty"`
	StartTime              int64                     `protobuf:"varint,11,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime                int64                     `protobuf:"varint,12,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	UpdateTime             int64                     `protobuf:"varint,13,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	StartDelayInSeconds    int32                     `protobuf:"varint,14,opt,name=start_delay_in_seconds,json=startDelayInSeconds" json:"start_delay_in_seconds,omitempty"`
	RetriedTaskId          string                    `protobuf:"bytes,15,opt,name=retried_task_id,json=retriedTaskId" json:"retried_task_id,omitempty"`
	Retried                bool                      `protobuf:"varint,16,opt,name=retried" json:"retried,omitempty"`
	Executed               bool                      `protobuf:"varint,17,opt,name=executed" json:"executed,omitempty"`
	CallbackFromWorker     bool                      `protobuf:"varint,18,opt,name=callback_from_worker,json=callbackFromWorker" json:"callback_from_worker,omitempty"`
	ResponseTimeoutSeconds int32                     `protobuf:"varint,19,opt,name=response_timeout_seconds,json=responseTimeoutSeconds" json:"response_timeout_seconds,omitempty"`
	WorkflowInstanceId     string                    `protobuf:"bytes,20,opt,name=workflow_instance_id,json=workflowInstanceId" json:"workflow_instance_id,omitempty"`
	WorkflowType           string                    `protobuf:"bytes,21,opt,name=workflow_type,json=workflowType" json:"workflow_type,omitempty"`
	TaskId                 string                    `protobuf:"bytes,22,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	ReasonForIncompletion  string                    `protobuf:"bytes,23,opt,name=reason_for_incompletion,json=reasonForIncompletion" json:"reason_for_incompletion,omitempty"`
	CallbackAfterSeconds   int64                     `protobuf:"varint,24,opt,name=callback_after_seconds,json=callbackAfterSeconds" json:"callback_after_seconds,omitempty"`
	WorkerId               string                    `protobuf:"bytes,25,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	OutputData             map[string]*_struct.Value `protobuf:"bytes,26,rep,name=output_data,json=outputData" json:"output_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	WorkflowTask           *WorkflowTask             `protobuf:"bytes,27,opt,name=workflow_task,json=workflowTask" json:"workflow_task,omitempty"`
	Domain                 string                    `protobuf:"bytes,28,opt,name=domain" json:"domain,omitempty"`
	InputMessage           *any.Any                  `protobuf:"bytes,29,opt,name=input_message,json=inputMessage" json:"input_message,omitempty"`
	OutputMessage          *any.Any                  `protobuf:"bytes,30,opt,name=output_message,json=outputMessage" json:"output_message,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                  `json:"-"`
	XXX_unrecognized       []byte                    `json:"-"`
	XXX_sizecache          int32                     `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_7843d3d2fd8c2dc8, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *Task) GetStatus() Task_Status {
	if m != nil {
		return m.Status
	}
	return Task_IN_PROGRESS
}

func (m *Task) GetInputData() map[string]*_struct.Value {
	if m != nil {
		return m.InputData
	}
	return nil
}

func (m *Task) GetReferenceTaskName() string {
	if m != nil {
		return m.ReferenceTaskName
	}
	return ""
}

func (m *Task) GetRetryCount() int32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *Task) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Task) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *Task) GetPollCount() int32 {
	if m != nil {
		return m.PollCount
	}
	return 0
}

func (m *Task) GetTaskDefName() string {
	if m != nil {
		return m.TaskDefName
	}
	return ""
}

func (m *Task) GetScheduledTime() int64 {
	if m != nil {
		return m.ScheduledTime
	}
	return 0
}

func (m *Task) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Task) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Task) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Task) GetStartDelayInSeconds() int32 {
	if m != nil {
		return m.StartDelayInSeconds
	}
	return 0
}

func (m *Task) GetRetriedTaskId() string {
	if m != nil {
		return m.RetriedTaskId
	}
	return ""
}

func (m *Task) GetRetried() bool {
	if m != nil {
		return m.Retried
	}
	return false
}

func (m *Task) GetExecuted() bool {
	if m != nil {
		return m.Executed
	}
	return false
}

func (m *Task) GetCallbackFromWorker() bool {
	if m != nil {
		return m.CallbackFromWorker
	}
	return false
}

func (m *Task) GetResponseTimeoutSeconds() int32 {
	if m != nil {
		return m.ResponseTimeoutSeconds
	}
	return 0
}

func (m *Task) GetWorkflowInstanceId() string {
	if m != nil {
		return m.WorkflowInstanceId
	}
	return ""
}

func (m *Task) GetWorkflowType() string {
	if m != nil {
		return m.WorkflowType
	}
	return ""
}

func (m *Task) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *Task) GetReasonForIncompletion() string {
	if m != nil {
		return m.ReasonForIncompletion
	}
	return ""
}

func (m *Task) GetCallbackAfterSeconds() int64 {
	if m != nil {
		return m.CallbackAfterSeconds
	}
	return 0
}

func (m *Task) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *Task) GetOutputData() map[string]*_struct.Value {
	if m != nil {
		return m.OutputData
	}
	return nil
}

func (m *Task) GetWorkflowTask() *WorkflowTask {
	if m != nil {
		return m.WorkflowTask
	}
	return nil
}

func (m *Task) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Task) GetInputMessage() *any.Any {
	if m != nil {
		return m.InputMessage
	}
	return nil
}

func (m *Task) GetOutputMessage() *any.Any {
	if m != nil {
		return m.OutputMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "conductor.proto.Task")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.Task.InputDataEntry")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.Task.OutputDataEntry")
	proto.RegisterEnum("conductor.proto.Task_Status", Task_Status_name, Task_Status_value)
}

func init() { proto.RegisterFile("model/task.proto", fileDescriptor_task_7843d3d2fd8c2dc8) }

var fileDescriptor_task_7843d3d2fd8c2dc8 = []byte{
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x9d, 0xf3, 0xe1, 0x8f, 0xeb, 0xd8, 0x56, 0x99, 0xc4, 0x61, 0x9c, 0xa4, 0x35, 0xb2, 0x65,
	0xf0, 0xc3, 0x60, 0x17, 0x69, 0x31, 0x74, 0xdd, 0x93, 0x63, 0x2b, 0xab, 0xb0, 0x24, 0x0e, 0x64,
	0x67, 0xc1, 0xf6, 0x22, 0x30, 0x12, 0xed, 0x0a, 0x96, 0x48, 0x8f, 0xa2, 0xd6, 0xfa, 0xc7, 0xed,
	0x3f, 0xed, 0x27, 0x0c, 0x24, 0x25, 0x35, 0x4b, 0x8b, 0x3d, 0xf5, 0x8d, 0x3c, 0xe7, 0xdc, 0xe3,
	0x7b, 0xaf, 0x78, 0xaf, 0xc1, 0x8a, 0x79, 0x40, 0xa3, 0x81, 0x24, 0xc9, 0xb2, 0xbf, 0x12, 0x5c,
	0x72, 0xd4, 0xf2, 0x39, 0x0b, 0x52, 0x5f, 0x72, 0x61, 0x80, 0x0e, 0x36, 0x92, 0x0f, 0x5c, 0x2c,
	0xe7, 0x11, 0xff, 0xf0, 0x49, 0xda, 0x39, 0x5e, 0x70, 0xbe, 0x88, 0xe8, 0x40, 0xdf, 0x1e, 0xd2,
	0xf9, 0x20, 0x91, 0x22, 0xf5, 0x65, 0xc6, 0x1e, 0x3e, 0x65, 0x09, 0x5b, 0x1b, 0xea, 0xf4, 0x9f,
	0x1d, 0xd8, 0x9a, 0x91, 0x64, 0x89, 0x8e, 0xa0, 0xa6, 0xfc, 0x3c, 0xb9, 0x5e, 0x51, 0x5c, 0xea,
	0x96, 0x7a, 0x35, 0xb7, 0xaa, 0x80, 0xd9, 0x7a, 0x45, 0xd1, 0x6b, 0x28, 0x27, 0x92, 0xc8, 0x34,
	0xc1, 0x1b, 0xdd, 0x52, 0xaf, 0x79, 0x7e, 0xdc, 0x7f, 0x92, 0x5a, 0x5f, 0x79, 0xf4, 0xa7, 0x5a,
	0xe3, 0x66, 0x5a, 0x34, 0x02, 0x08, 0xd9, 0x2a, 0x95, 0x5e, 0x40, 0x24, 0xc1, 0x9b, 0xdd, 0xcd,
	0x5e, 0xfd, 0xfc, 0xbb, 0x2f, 0x47, 0x3a, 0x4a, 0x37, 0x26, 0x92, 0xd8, 0x4c, 0x8a, 0xb5, 0x5b,
	0x0b, 0xf3, 0x3b, 0xea, 0xc3, 0xae, 0xa0, 0x73, 0x2a, 0x28, 0xf3, 0xa9, 0xa7, 0x33, 0x64, 0x24,
	0xa6, 0x78, 0x4b, 0x67, 0xf8, 0xac, 0xa0, 0x94, 0xcb, 0x0d, 0x89, 0x29, 0x7a, 0x01, 0x75, 0x41,
	0xa5, 0x58, 0x7b, 0x3e, 0x4f, 0x99, 0xc4, 0xdb, 0xdd, 0x52, 0x6f, 0xdb, 0x05, 0x0d, 0x8d, 0x14,
	0x82, 0x2c, 0xd8, 0x4c, 0xe8, 0x9f, 0xb8, 0xac, 0x09, 0x75, 0x44, 0x67, 0xd0, 0xf4, 0xb9, 0x10,
	0x34, 0x22, 0x32, 0xe4, 0xcc, 0x0b, 0x03, 0x5c, 0xd1, 0xee, 0x8d, 0x47, 0xa8, 0x13, 0xa0, 0x13,
	0x80, 0x15, 0x8f, 0xa2, 0xcc, 0xb8, 0xaa, 0xe3, 0x6b, 0x0a, 0x31, 0xbe, 0xa7, 0xd0, 0xd0, 0xe9,
	0x05, 0x74, 0x6e, 0x52, 0xac, 0x69, 0x93, 0xba, 0x02, 0xc7, 0x74, 0xae, 0x93, 0x3b, 0x83, 0x66,
	0xe2, 0xbf, 0xa7, 0x41, 0x1a, 0xd1, 0xc0, 0x93, 0x61, 0x4c, 0x31, 0x74, 0x4b, 0xbd, 0x4d, 0xb7,
	0x51, 0xa0, 0xb3, 0x30, 0xa6, 0xea, 0x97, 0x12, 0x49, 0x84, 0x34, 0x92, 0xba, 0x96, 0xd4, 0x34,
	0xa2, 0xe9, 0x43, 0xa8, 0x52, 0x96, 0xc5, 0xef, 0x68, 0xb2, 0x42, 0x99, 0x89, 0x7c, 0x01, 0xf5,
	0x74, 0x15, 0x10, 0x49, 0x0d, 0xdb, 0xd0, 0x2c, 0x18, 0x48, 0x0b, 0x5e, 0x41, 0xdb, 0x58, 0x07,
	0x34, 0x22, 0x6b, 0x2f, 0x64, 0x5e, 0x42, 0xd5, 0x17, 0x49, 0x70, 0x53, 0x17, 0xb4, 0xab, 0xd9,
	0xb1, 0x22, 0x1d, 0x36, 0x35, 0x14, 0xfa, 0x1e, 0x5a, 0xaa, 0x81, 0xa1, 0x4a, 0x5a, 0x95, 0x18,
	0x06, 0xb8, 0x65, 0x3a, 0x94, 0xc1, 0xaa, 0xfb, 0x4e, 0x80, 0x30, 0x54, 0x32, 0x00, 0x5b, 0xdd,
	0x52, 0xaf, 0xea, 0xe6, 0x57, 0xd4, 0x81, 0x2a, 0xfd, 0x48, 0xfd, 0x54, 0xd2, 0x00, 0x3f, 0xd3,
	0x54, 0x71, 0x47, 0x2f, 0x61, 0xcf, 0x27, 0x51, 0xf4, 0x40, 0xfc, 0xa5, 0x37, 0x17, 0x3c, 0xf6,
	0xd4, 0xfb, 0xa6, 0x02, 0x23, 0xad, 0x43, 0x39, 0x77, 0x29, 0x78, 0x7c, 0xaf, 0x19, 0xf4, 0x06,
	0xb0, 0xa0, 0xc9, 0x8a, 0xb3, 0xc4, 0xd4, 0xc9, 0x53, 0x59, 0x94, 0xb1, 0xab, 0xcb, 0x68, 0xe7,
	0xfc, 0xcc, 0xd0, 0x79, 0x25, 0x2f, 0x61, 0x2f, 0x9f, 0x1e, 0x2f, 0x64, 0x89, 0x24, 0xea, 0x55,
	0x85, 0x01, 0xde, 0xd3, 0xe5, 0xa0, 0x9c, 0x73, 0x32, 0xca, 0x09, 0xd0, 0xb7, 0xd0, 0x28, 0x22,
	0xf4, 0x6c, 0xec, 0x6b, 0xe9, 0x4e, 0x0e, 0xea, 0xf9, 0x38, 0x80, 0x4a, 0xde, 0x98, 0xb6, 0xa6,
	0xcb, 0xd2, 0x74, 0xe4, 0x47, 0x38, 0x10, 0x94, 0x24, 0x9c, 0x79, 0x73, 0x2e, 0xbc, 0x90, 0xf9,
	0x3c, 0x5e, 0x45, 0x54, 0x3d, 0x28, 0x7c, 0xa0, 0x85, 0xfb, 0x86, 0xbe, 0xe4, 0xc2, 0x79, 0x44,
	0xa2, 0xd7, 0xd0, 0x2e, 0x7a, 0x42, 0xe6, 0x92, 0x8a, 0xa2, 0x3e, 0xac, 0x3f, 0x69, 0xd1, 0xb1,
	0xa1, 0x22, 0xf3, 0xea, 0x8e, 0xa0, 0x66, 0x7a, 0xa7, 0x12, 0x39, 0x34, 0x33, 0x6c, 0x00, 0x27,
	0x40, 0x97, 0x50, 0xe7, 0xa9, 0x2c, 0xc6, 0xb1, 0xa3, 0xc7, 0xf1, 0xec, 0xcb, 0xe3, 0x38, 0xd1,
	0xc2, 0x4f, 0xf3, 0x08, 0xbc, 0x00, 0xd0, 0xc5, 0xe3, 0x86, 0x90, 0x64, 0x89, 0x8f, 0xba, 0xa5,
	0x5e, 0xfd, 0xfc, 0xe4, 0x33, 0xa7, 0xfb, 0xbc, 0x43, 0x24, 0x59, 0x3e, 0xea, 0x97, 0x5a, 0x36,
	0x6d, 0x28, 0x07, 0x3c, 0x26, 0x21, 0xc3, 0xc7, 0xa6, 0x5d, 0xe6, 0x86, 0x7e, 0x82, 0x86, 0xd9,
	0x18, 0x31, 0x4d, 0x12, 0xb2, 0xa0, 0xf8, 0x44, 0x7b, 0xef, 0xf5, 0xcd, 0x02, 0xeb, 0xe7, 0x0b,
	0xac, 0x3f, 0x64, 0x6b, 0x77, 0x47, 0x4b, 0xaf, 0x8d, 0x12, 0xfd, 0x0c, 0xcd, 0xac, 0xbc, 0x3c,
	0xf6, 0xf9, 0xff, 0xc4, 0x36, 0x8c, 0x36, 0x0b, 0xee, 0xcc, 0xa0, 0xf9, 0xdf, 0x0d, 0xa4, 0xb6,
	0xc4, 0x92, 0xae, 0xb3, 0x45, 0xa8, 0x8e, 0xe8, 0x07, 0xd8, 0xfe, 0x8b, 0x44, 0x29, 0xd5, 0x2b,
	0xb0, 0x7e, 0xde, 0xfe, 0xcc, 0xf7, 0x37, 0xc5, 0xba, 0x46, 0xf4, 0x76, 0xe3, 0x4d, 0xa9, 0x73,
	0x07, 0xad, 0x27, 0x8d, 0xfc, 0x1a, 0xb6, 0xa7, 0x7f, 0x97, 0xa0, 0x6c, 0x36, 0x2d, 0x6a, 0x41,
	0xdd, 0xb9, 0xf1, 0x6e, 0xdd, 0xc9, 0x2f, 0xae, 0x3d, 0x9d, 0x5a, 0xdf, 0xa0, 0x1d, 0xa8, 0x8e,
	0x86, 0x37, 0x23, 0xfb, 0xca, 0x1e, 0x5b, 0x25, 0x04, 0x50, 0xbe, 0x1c, 0x3a, 0xea, 0xbc, 0x81,
	0x9e, 0x43, 0xc7, 0x9c, 0xbd, 0x7b, 0x67, 0xf6, 0xce, 0x9b, 0xd9, 0xee, 0xb5, 0x73, 0x33, 0xbc,
	0xf2, 0x6c, 0xd7, 0x9d, 0xb8, 0xd6, 0x26, 0x6a, 0x40, 0x6d, 0x34, 0xb9, 0xbe, 0xbd, 0xb2, 0x67,
	0xf6, 0xd8, 0xda, 0x42, 0x87, 0xb0, 0x5f, 0x5c, 0x4d, 0x84, 0x16, 0x4e, 0xad, 0x6d, 0xa5, 0x9c,
	0x8e, 0xde, 0xd9, 0xe3, 0x3b, 0x65, 0x5c, 0x56, 0xd7, 0x99, 0x73, 0x6d, 0x8f, 0xbd, 0xc9, 0xdd,
	0xcc, 0xaa, 0xa0, 0x5d, 0x68, 0xb9, 0xf6, 0x70, 0xfc, 0xbb, 0x77, 0x39, 0x71, 0x3d, 0xd7, 0x76,
	0xef, 0x6e, 0xac, 0x2a, 0xaa, 0x43, 0x65, 0xfa, 0xab, 0x73, 0x7b, 0x6b, 0x8f, 0xad, 0xda, 0x05,
	0x81, 0x23, 0x9f, 0xc7, 0x7d, 0x46, 0xe5, 0x3c, 0x0a, 0x3f, 0x3e, 0x7d, 0x36, 0x17, 0x65, 0xf5,
	0x42, 0x6e, 0x1f, 0xfe, 0x78, 0xbb, 0x08, 0xe5, 0xfb, 0xf4, 0xa1, 0xef, 0xf3, 0x78, 0x90, 0x69,
	0x07, 0x85, 0x76, 0xe0, 0x47, 0x21, 0x65, 0x72, 0xb0, 0xe0, 0x0b, 0xb1, 0xf2, 0x1f, 0xe1, 0xfa,
	0x1f, 0xf2, 0xa1, 0xac, 0xad, 0x5e, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x54, 0xf9, 0x20, 0xc1,
	0x54, 0x07, 0x00, 0x00,
}
