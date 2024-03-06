run:
	go run main.go

check-install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger

swagger: check-install
	swagger generate spec -o ./swagger.yaml --scan-models
