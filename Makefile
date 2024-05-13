run: |
	gofmt -w .
	go run ./cmd/main.go


lint:
	golangci-lint run ./... --timeout=2m -D staticcheck,govet