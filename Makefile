gen:
	protoc --proto_path=proto proto/pcbook.proto --go_out=pb --go-grpc_out=pb --go_opt=Mproto/pcbook.proto=pb --go-grpc_opt=Mproto/pcbook.proto=pb

clean:
	rm pb/*.go

run:
	go run main.go