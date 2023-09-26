# Build a CRUD API with gRPC and Golang
Steps ($:represents terminal commands)

`$ mkdir simple_crud_grpc`
`$ cd simple_crud_grpc`
`$ mkdir protos`
`$ touch moviesapp.proto`

# Update movies.proto
`$ cd ..`

`$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
`$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

`$ go mod init moviesapp.com/grpc`
`$ export PATH="$PATH:$(go env GOPATH)/bin"`

# Generate the proto files 
`$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/moviesapp.proto`

# Update the packages 
`$ go mod tidy`

`$ touch server/main.go`
`$ touch client/main.go`

# Update server/main.go 
`$ gofmt -s -w server/main.go`
`$ go run server/main.go`
`$ go run client/main.go`
