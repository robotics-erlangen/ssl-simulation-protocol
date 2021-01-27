// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_simulation_synchronous.proto

package sim

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request from the team to the simulator
type SimulationSyncRequest struct {
	// The simulation step [s] to perform
	SimStep *float32 `protobuf:"fixed32,1,opt,name=sim_step,json=simStep" json:"sim_step,omitempty"`
	// An optional simulator command
	SimulatorCommand *SimulatorCommand `protobuf:"bytes,2,opt,name=simulator_command,json=simulatorCommand" json:"simulator_command,omitempty"`
	// An optional robot control command
	RobotControl         *RobotControl `protobuf:"bytes,3,opt,name=robot_control,json=robotControl" json:"robot_control,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SimulationSyncRequest) Reset()         { *m = SimulationSyncRequest{} }
func (m *SimulationSyncRequest) String() string { return proto.CompactTextString(m) }
func (*SimulationSyncRequest) ProtoMessage()    {}
func (*SimulationSyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea856983e544df96, []int{0}
}

func (m *SimulationSyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimulationSyncRequest.Unmarshal(m, b)
}
func (m *SimulationSyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimulationSyncRequest.Marshal(b, m, deterministic)
}
func (m *SimulationSyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulationSyncRequest.Merge(m, src)
}
func (m *SimulationSyncRequest) XXX_Size() int {
	return xxx_messageInfo_SimulationSyncRequest.Size(m)
}
func (m *SimulationSyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulationSyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimulationSyncRequest proto.InternalMessageInfo

func (m *SimulationSyncRequest) GetSimStep() float32 {
	if m != nil && m.SimStep != nil {
		return *m.SimStep
	}
	return 0
}

func (m *SimulationSyncRequest) GetSimulatorCommand() *SimulatorCommand {
	if m != nil {
		return m.SimulatorCommand
	}
	return nil
}

func (m *SimulationSyncRequest) GetRobotControl() *RobotControl {
	if m != nil {
		return m.RobotControl
	}
	return nil
}

// Response to last SimulationSyncRequest
type SimulationSyncResponse struct {
	// List of detection frames for all cameras with the state after the simulation step in the request was performed
	Detection []*SSL_DetectionFrame `protobuf:"bytes,1,rep,name=detection" json:"detection,omitempty"`
	// An optional robot control response
	RobotControlResponse *RobotControlResponse `protobuf:"bytes,2,opt,name=robot_control_response,json=robotControlResponse" json:"robot_control_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SimulationSyncResponse) Reset()         { *m = SimulationSyncResponse{} }
func (m *SimulationSyncResponse) String() string { return proto.CompactTextString(m) }
func (*SimulationSyncResponse) ProtoMessage()    {}
func (*SimulationSyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea856983e544df96, []int{1}
}

func (m *SimulationSyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimulationSyncResponse.Unmarshal(m, b)
}
func (m *SimulationSyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimulationSyncResponse.Marshal(b, m, deterministic)
}
func (m *SimulationSyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulationSyncResponse.Merge(m, src)
}
func (m *SimulationSyncResponse) XXX_Size() int {
	return xxx_messageInfo_SimulationSyncResponse.Size(m)
}
func (m *SimulationSyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulationSyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimulationSyncResponse proto.InternalMessageInfo

func (m *SimulationSyncResponse) GetDetection() []*SSL_DetectionFrame {
	if m != nil {
		return m.Detection
	}
	return nil
}

func (m *SimulationSyncResponse) GetRobotControlResponse() *RobotControlResponse {
	if m != nil {
		return m.RobotControlResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*SimulationSyncRequest)(nil), "SimulationSyncRequest")
	proto.RegisterType((*SimulationSyncResponse)(nil), "SimulationSyncResponse")
}

func init() {
	proto.RegisterFile("ssl_simulation_synchronous.proto", fileDescriptor_ea856983e544df96)
}

var fileDescriptor_ea856983e544df96 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x6e, 0xea, 0x30,
	0x14, 0x86, 0x15, 0x18, 0xee, 0xbd, 0xe6, 0x22, 0x95, 0xb4, 0xa0, 0x14, 0x75, 0x88, 0xe8, 0xc2,
	0x42, 0xa2, 0x32, 0x54, 0x9d, 0x3a, 0x94, 0xaa, 0x4b, 0x99, 0xec, 0xad, 0x8b, 0x15, 0x8c, 0x0b,
	0x16, 0xb1, 0x4f, 0xea, 0xe3, 0x54, 0xe2, 0x4d, 0xfa, 0x04, 0x7d, 0xce, 0x2a, 0x60, 0xa0, 0x89,
	0x18, 0xfd, 0x7f, 0xdf, 0x39, 0xfa, 0x7d, 0x48, 0x8c, 0x98, 0x73, 0x54, 0xba, 0xcc, 0x33, 0xa7,
	0xc0, 0x70, 0xdc, 0x1a, 0xb1, 0xb6, 0x60, 0xa0, 0xc4, 0xa4, 0xb0, 0xe0, 0x60, 0x38, 0xac, 0x8c,
	0x4f, 0x85, 0x15, 0x5d, 0x4a, 0x27, 0x45, 0xe5, 0x79, 0x76, 0xdb, 0x98, 0xb6, 0xb0, 0x00, 0xc7,
	0xdf, 0xa5, 0x5c, 0x2e, 0x32, 0xb1, 0xf1, 0xd2, 0xe8, 0xac, 0x24, 0xc0, 0x38, 0x0b, 0xb9, 0x77,
	0x6e, 0x1a, 0x4e, 0x8d, 0x8e, 0xbe, 0x03, 0xd2, 0x67, 0x47, 0xc8, 0xb6, 0x46, 0x50, 0xf9, 0x51,
	0x4a, 0x74, 0xe1, 0x35, 0xf9, 0x8b, 0x4a, 0x73, 0x74, 0xb2, 0x88, 0x82, 0x38, 0x18, 0xb7, 0xe8,
	0x1f, 0x54, 0x9a, 0x39, 0x59, 0x84, 0x8f, 0xa4, 0xe7, 0x17, 0x82, 0xe5, 0x02, 0xb4, 0xce, 0xcc,
	0x32, 0x6a, 0xc5, 0xc1, 0xb8, 0x33, 0xed, 0x25, 0xec, 0x40, 0x66, 0x7b, 0x40, 0x2f, 0xb0, 0x91,
	0x84, 0x53, 0xd2, 0xad, 0x35, 0x8d, 0xda, 0xbb, 0xd9, 0x6e, 0x42, 0xab, 0x74, 0xb6, 0x0f, 0xe9,
	0x7f, 0xfb, 0xeb, 0x35, 0xfa, 0x0a, 0xc8, 0xa0, 0x59, 0x14, 0x0b, 0x30, 0x28, 0xc3, 0x3b, 0xf2,
	0xef, 0x78, 0xbd, 0x28, 0x88, 0xdb, 0xe3, 0xce, 0xf4, 0x32, 0x61, 0x6c, 0xce, 0x9f, 0x0f, 0xe9,
	0x8b, 0xcd, 0xb4, 0xa4, 0x27, 0x2b, 0x7c, 0x25, 0x83, 0x5a, 0x03, 0x6e, 0xfd, 0x32, 0xff, 0x8d,
	0x7e, 0xbd, 0x8a, 0x87, 0xf4, 0xca, 0x9e, 0x49, 0x9f, 0x1e, 0xde, 0xee, 0x57, 0xca, 0xad, 0xcb,
	0x45, 0x22, 0x40, 0xa7, 0xd5, 0xe0, 0xac, 0x2c, 0x26, 0x8c, 0xcd, 0x53, 0xc4, 0x7c, 0x72, 0x3a,
	0xfd, 0x64, 0x77, 0x72, 0x01, 0x79, 0x5a, 0x6c, 0x56, 0x29, 0x2a, 0xfd, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xe7, 0xeb, 0xe2, 0x73, 0x23, 0x02, 0x00, 0x00,
}
