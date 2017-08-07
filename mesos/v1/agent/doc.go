package mesos_v1_agent

//go:generate protoc --proto_path=.:../../.. --go_out=Mmesos/v1/mesos.proto=github.com/mesos/go-proto/mesos/v1:. agent.proto
