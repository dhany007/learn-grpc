PROTO_DIR = proto
PACKAGE = github.com/dhany007/learn-grpc

build-greet:
	protoc -Igreet/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. greet/${PROTO_DIR}/*.proto
	go build -o bin/greet/server ./greet/server
	go build -o bin/greet/client ./greet/client

build-calculator:
	protoc -Icalculator/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. calculator/${PROTO_DIR}/*.proto
	go build -o bin/calculator/server ./calculator/server
	go build -o bin/calculator/client ./calculator/client

gen-ssh:
	chmod +x ssl/ssl.sh

evans:
	evans --host localhost --port 50052 --reflection repl

build-blog:
	protoc -Iblog/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. blog/${PROTO_DIR}/*.proto
	go build -o bin/blog/server ./blog/server
	go build -o bin/blog/client ./blog/client
