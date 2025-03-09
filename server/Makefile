

.PHONY: generate
generate: ## Generate all (protocol buffers in this case)
	go generate ./...

.PHONY: install-deps
install-deps:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
