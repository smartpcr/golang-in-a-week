module monitoring/server

go 1.20

require monitoring/proto v0.1.3

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.13.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/grpc v1.57.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace monitoring/proto => ../proto
