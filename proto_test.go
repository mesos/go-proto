package mesos

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/allocator"
	"github.com/mesos/go-proto/mesos/v1/executor"
	"github.com/mesos/go-proto/mesos/v1/maintenance"
	"github.com/mesos/go-proto/mesos/v1/master"
	"github.com/mesos/go-proto/mesos/v1/scheduler"
)

func newAgentID(id string) *mesos_v1.AgentID {
	return &mesos_v1.AgentID{
		Value: proto.String(id),
	}
}

func newFrameworkID(id string) *mesos_v1.FrameworkID {
	return &mesos_v1.FrameworkID{
		Value: proto.String(id),
	}
}

func newTaskID(id string) *mesos_v1.TaskID {
	return &mesos_v1.TaskID{
		Value: proto.String(id),
	}
}

// Test that we can import the mesos_v1 package.
func TestMesosProto(t *testing.T) {
	task := mesos_v1.TaskInfo{
		Name:    proto.String("check"),
		TaskId:  newTaskID("task-id-1"),
		AgentId: newAgentID("agent-id-1"),
	}

	data, err := proto.Marshal(&task)
	if err != nil {
		t.Errorf("marshalling TaskInfo: %s", err)
	}

	err = proto.Unmarshal(data, &task)
	if err != nil {
		t.Errorf("unmarshalling TaskInfo: %s", err)
	}
}

// Test that we can import the mesos_v1_executor package.
func TestExecutorProto(t *testing.T) {
	event := mesos_v1_executor.Event{
		Type: mesos_v1_executor.Event_UNKNOWN.Enum(),
	}

	data, err := proto.Marshal(&event)
	if err != nil {
		t.Errorf("marshalling Event: %s", err)
	}

	err = proto.Unmarshal(data, &event)
	if err != nil {
		t.Errorf("unmarshalling Event: %s", err)
	}
}

// Test that we can import the mesos_v1_agent package.
func TestAgentProto(t *testing.T) {
	call := mesos_v1_agent.Call{
		Type: mesos_v1_agent.Call_GET_HEALTH.Enum(),
	}

	data, err := proto.Marshal(&call)
	if err != nil {
		t.Errorf("marshalling Call: %s", err)
	}

	err = proto.Unmarshal(data, &call)
	if err != nil {
		t.Errorf("unmarshalling Call: %s", err)
	}
}

// Test that we can import the mesos_v1_allocator package.
func TestAllocatorProto(t *testing.T) {
	status := mesos_v1_allocator.InverseOfferStatus{
		Status: mesos_v1_allocator.InverseOfferStatus_ACCEPT.Enum(),
		Timestamp: &mesos_v1.TimeInfo{
			Nanoseconds: proto.Int64(0),
		},
		FrameworkId: newFrameworkID("framework-1"),
	}

	data, err := proto.Marshal(&status)
	if err != nil {
		t.Errorf("marshalling InverseOfferStatus: %s", err)
	}

	err = proto.Unmarshal(data, &status)
	if err != nil {
		t.Errorf("unmarshalling InverseOfferStatus: %s", err)
	}
}

// Test that we can import the mesos_v1_maintenance package.
func TestMaintenanceProto(t *testing.T) {
	window := mesos_v1_maintenance.Window{
		MachineIds: nil,
		Unavailability: &mesos_v1.Unavailability{
			Start: &mesos_v1.TimeInfo{
				Nanoseconds: proto.Int64(0),
			},
			Duration: &mesos_v1.DurationInfo{
				Nanoseconds: proto.Int64(0),
			},
		},
	}

	data, err := proto.Marshal(&window)
	if err != nil {
		t.Errorf("marshalling Window: %s", err)
	}

	err = proto.Unmarshal(data, &window)
	if err != nil {
		t.Errorf("unmarshalling Window: %s", err)
	}
}

// Test that we can import the mesos_v1_master package.
func TestMasterProto(t *testing.T) {
	call := mesos_v1_master.Call{
		Type: mesos_v1_master.Call_GET_HEALTH.Enum(),
	}

	data, err := proto.Marshal(&call)
	if err != nil {
		t.Errorf("marshalling Call: %s", err)
	}

	err = proto.Unmarshal(data, &call)
	if err != nil {
		t.Errorf("unmarshalling Call: %s", err)
	}
}

// Test that we can import the mesos_v1_scheduler package.
func TestSchedulerProto(t *testing.T) {
	call := mesos_v1_scheduler.Call{
		FrameworkId: newFrameworkID("framework-1"),
		Type:        mesos_v1_scheduler.Call_SUBSCRIBE.Enum(),
	}

	data, err := proto.Marshal(&call)
	if err != nil {
		t.Errorf("marshalling Call: %s", err)
	}

	err = proto.Unmarshal(data, &call)
	if err != nil {
		t.Errorf("unmarshalling Call: %s", err)
	}
}
