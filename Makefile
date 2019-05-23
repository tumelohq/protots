SRC_NO_VENDOR := $(shell find . -path -prune -o -name '*.go' -not -path './vendor/*' -not -path './proto/*')

.PHONY: go_fmt
go_fmt: ## Format go code
	@gofmt -w -s $(SRC_NO_VENDOR)
