# Format all go files in the project using gofmt
format:
	find . -name '*.go' -not -path './.devcontainer/*' -exec gofmt -w {} +
