/usr/local/bin/protoc \
    --go_out=../backend/laptiming/ --go_opt=paths=source_relative  --go-grpc_out=../backend/laptiming --go-grpc_opt=paths=source_relative \
    --cpp_out=../embedded_interface/protobuf/ --grpc_out=../embedded_interface/protobuf/ --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`\
    ./messages.proto