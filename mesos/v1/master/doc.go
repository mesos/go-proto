package mesos_v1_master

//go:generate protoc --proto_path=.:../../.. --go_out=Mmesos/v1/mesos.proto=github.com/mesos/go-proto/mesos/v1,Mmesos/v1/maintenance/maintenance.proto=github.com/mesos/go-proto/mesos/v1/maintenance,Mmesos/v1/quota/quota.proto=github.com/mesos/go-proto/mesos/v1/quota:.  master.proto
