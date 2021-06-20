set GOARCH=amd64
set GOOS=linux
go build -o deploy/bin/auth services/auth/cmd/grpc/main.go
go build -o deploy/bin/billing services/billing/cmd/grpc/main.go
go build -o deploy/bin/entities services/entities/cmd/grpc/main.go
go build -o deploy/bin/users services/users/cmd/grpc/main.go
go build -o deploy/bin/api services/api/graphql/cmd/main.go
set GOOS=windows