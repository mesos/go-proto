# go-proto
Mesos protobuf bindings for Go.

`go-proto` contains just the generated
[protobuf](https://github.com/golang/protobuf) bindings for the v1 Mesos
[scheduler](http://mesos.apache.org/documentation/latest/scheduler-http-api/),
[executor](http://mesos.apache.org/documentation/latest/executor-http-api/)
and [operator](http://mesos.apache.org/documentation/latest/operator-http-api/)
APIs. The intention of `go-proto` is to provide a single package
that exports the protobuf definitions so that developers writing
Mesos API bindings can conveniently import them.
