SHELL=/bin/bash -o pipefail

PROJECT_NAME := "mono"
PKG := "gitlab.com/hedge_app/$(PROJECT_NAME)"

commitid = master
prelink = "https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/"
annotation_end = "/third_party/googleapis/google/api/annotations.proto"
http_end = "/third_party/googleapis/google/api/http.proto"
annotation_link = $(prelink)$(commitid)$(annotation_end)
http_link = $(prelink)$(commitid)$(http_end)

SRC := $(shell find . -path -prune -o -name '*.go')
SRC_NO_VENDOR := $(shell find . -path -prune -o -name '*.go' -not -path './vendor/*' -not -path './proto/*')

location = "vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/"
http_out = $(location)http.proto
annotations_out = $(location)annotations.proto


.PHONY: proto_gen
proto_gen: ## Generate the proto files
	rm -rf proto/gen/go
	echo NOT CONFIGURED!
	@echo "GOPATH must be set correctly, it is currently" $(GOPATH)
	@mkdir -p $(location)
	@wget -q $(annotation_link) -O $(annotations_out)
	@wget -q $(http_link) -O $(http_out)
	cd idl && prototool generate
