gen:
	protoc -I . --go_out ./pb --go_opt paths=source_relative --go-grpc_out ./pb --go-grpc_opt paths=source_relative --proto_path=proto helloworld.proto
	protoc -I . --grpc-gateway_out ./pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true --proto_path=proto helloworld.proto
