gen:
	protoc --proto_path=./proto  --go_out=plugins=grpc:pb  proto/*.proto

clean:
	del pb\*.go

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

server1-tls:
	go run cmd/server/main.go -port 50051 -tls

server2-tls:
	go run cmd/server/main.go -port 50052 -tls

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

client-tls:
	go run cmd/client/main.go -address 0.0.0.0:8080 -tls

test:
	go test -cover -race ./...s

evans:
	evans -r repl -p 8080

cert:
	cd cert; ./gen.sh; cd ..

.PHONY: clean gen server client test cert