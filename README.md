# Example of microservices architecture using [Golang](https://go.dev/) + [gRPC](https://grpc.io/)

### How to run?
1. clone the repo
2. `cd` into `go-microservices/auth`
3. run `go mod tidy`
4. run `go run main.go`
5. do steps 2-4 on `go-microservices/gateway`
6. open a browser `localhost:3000` (by default)
7. expected result: can see user data in the browser

### List of projects:
- [proto](https://github.com/vadymbarabanov/go-microservices/tree/main/proto) - a storage for proto files & generated code
- [gateway](https://github.com/vadymbarabanov/go-microservices/tree/main/gateway) - an entrypoint for regular http/REST clients e.g. browser. It's also a gRPC client
- [auth](https://github.com/vadymbarabanov/go-microservices/tree/main/auth) - an internal microservice that takes user info and returns a result. It's a gRPC server
