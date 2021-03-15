

# https://github.com/gogo/protobuf
proto.install:
	go get -u -v github.com/gogo/protobuf/protoc-gen-gofast
	go get -u -v github.com/gogo/protobuf/proto
	go get -u -v github.com/gogo/protobuf/jsonpb
	go get -u -v github.com/gogo/protobuf/protoc-gen-gogo
	go get -u -v github.com/gogo/protobuf/gogoproto
