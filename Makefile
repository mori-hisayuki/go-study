# Format all go files in the project using gofmt
format:
	find . -name '*.go' -not -path './.devcontainer/*' -exec gofmt -w {} + && swag init -g ./api/cmd/main.go

run:
	go run ./cmd/.
