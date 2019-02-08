# Techshare Client
Sample project to benchmark REST API call vs gRPC call

## Command
### Request to server 
```
go run main.go -mode=<rest or grpc> -server_addr=<server address>
```

### Sample run call to REST API server
```
go run main.go -mode=rest -server_addr=http://127.0.0.1:10000
```

### Sample run call to gRPC Server
```
go run main.go -server_addr=localhost:11000
```