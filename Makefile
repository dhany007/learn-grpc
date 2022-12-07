PROTO_DIR = proto
PACKAGE = github.com/dhany007/learn-grpc


build-greet:
	protoc -Igreet/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. greet/${PROTO_DIR}/*.proto
	go build -o bin/greet/server ./greet/server
	go build -o bin/greet/client ./greet/client
