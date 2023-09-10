gen:
	protoc -I ./proto/ \
  		-I include/googleapis -I include/grpc-gateway \
  		--go_out=paths=source_relative:./proto \
	  	--go-grpc_out=paths=source_relative:./proto  \
	  	--grpc-gateway_out=paths=source_relative:./proto \
		--openapiv2_out ./openapiv2 --openapiv2_opt logtostderr=true \
	   proto/users/v1/*.proto

swag:
	docker run -p 80:8080 \
   		-e SWAGGER_JSON=/openapiv2/users/v1/api.swagger.json \
   		-v /Users/ms/go/src/gostudy/protobuf_example/open_api_gen/openapiv2/:/openapiv2 \
		swaggerapi/swagger-ui


install:
	go get -u \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/rakyll/statik
