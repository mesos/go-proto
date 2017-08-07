package mesos_v1_maintenance

//go:generate protoc --proto_path=.:../../.. --go_out=Mmesos/v1/mesos.proto=github.com/mesos/go-proto/mesos/v1,Mmesos/v1/allocator/allocator.proto=github.com/mesos/go-proto/mesos/v1/allocator:.  maintenance.proto
