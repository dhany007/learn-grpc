PROTO_DIR = proto
PACKAGE = github.com/dhany007/learn-grpc


proto-greet:
	protoc -Igreet/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. greet/${PROTO_DIR}/*.proto

start-server-greet:
	go run greet/server/main.go

start-client-greet:
	go run greet/client/main.go