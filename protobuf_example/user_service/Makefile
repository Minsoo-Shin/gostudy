.PHONY: gen
gen:
	protoc api/v1/*.proto \
		-I include/googleapis -I include/grpc-gateway \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--grpc-gateway_out=paths=source_relative:. \
		--openapiv2_out ./swagger --openapiv2_opt logtostderr=true \
		--proto_path=.

#	protoc api/v1/*.proto \
#		--javascript_out=. \
#		--proto_path=.


install:
	go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
