uname := $(shell uname)

ifeq ($(uname), Linux)
	SED := sed -r -i
endif

ifeq ($(uname), Darwin)
	SED := sed -E -i ''
endif

ifeq ($(SED),)
$(error Unable to detect the correct options for editing files in-place with sed(1))
endif

.PHONY: default
default: help

.PHONY: help
help:
	@echo Mesos protobuf bindings for Go.
	@echo
	@echo Targets:
	@echo "    generate    - Re-generate all the protobuf bindings."
	@echo "    update      - Update the .proto files from \$$MESOS/include/mesos/v1"
	@echo "    check       - Run \"go test .\""
	@echo
	@echo "To regenerate the Go code, first install the protobuf tools:"
	@echo "    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}"
	@echo
	@echo "Make sure that \$$GOPATH/bin is in your \$$PATH and use go generate to"
	@echo "generate the Go code in each directory:"
	@echo
	@echo "Resources:"
	@echo "    https://developers.google.com/protocol-buffers/docs/gotutorial"
	@echo "    https://github.com/golang/protobuf"
	@echo
	@echo "The protobuf code here is generated with proto3. proto3 always"
	@echo "converts JSON field names to lowerCamelCase, but for historical"
	@echo "compatibility, the Mesos JSON serialization requires lower_snake_case"
	@echo "names. Since this is not regarded as a protobuf bug, we have to"
	@echo "rewrite the name in post."
	@echo
	@echo "See: https://github.com/golang/protobuf/issues/183"


.PHONY: generate
generate:
	@find . -name *.pb.go | xargs rm -f
	env PATH=$$GOPATH/bin:$$PATH go generate ./...
	@$(SED) '-es/(json=)([a-zA-Z_]+)(.*json:[^a-z]*)([a-z_]+)/\1\4\3\4/' $$(find . -name *.pb.go)

.PHONY: update
update:
	@for f in $$(cd $(MESOS)/include/mesos/v1 && find . -name \*.proto); do \
		mkdir -p $$(dirname mesos/v1/$$f); \
		cp $(MESOS)/include/mesos/v1/$$f mesos/v1/$$f; \
	done

.PHONY: check
check:
	@go test .
