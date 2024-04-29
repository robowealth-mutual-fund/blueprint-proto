.PHONY: pbgen
pbgen:
	protoc --proto_path=internal --proto_path=thirdParty \
		--go_out plugins=grpc:pkg/. --go_opt=paths=source_relative \
				v1/todo/message/request.proto \
				v1/todo/message/response.proto \
		--experimental_allow_proto3_optional

	protoc --proto_path=internal --proto_path=thirdParty \
		--go_out=plugins=grpc:pkg/. --go_opt=paths=source_relative \
	 	--grpc-gateway_out=pkg/. --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=swagger --openapiv2_opt=logtostderr=true,allow_merge=true,merge_file_name=todo \
				v1/todo/service.proto \
	 	--experimental_allow_proto3_optional

	protoc-go-inject-tag -input=pkg/v1/todo/message/request.pb.go