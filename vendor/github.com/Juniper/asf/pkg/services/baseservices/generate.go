package baseservices

//go:generate -command gen_proto ./../../../tools/protoc.sh -I ../../../vendor/ -I ../../../vendor/github.com/gogo/protobuf/protobuf  -I. --gogofaster_out=plugins=grpc:.
//go:generate gen_proto base.proto
