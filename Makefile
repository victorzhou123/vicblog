# env info
PKG_DIR := .

define proto
	protoc -I. --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false common/controller/rpc/*.proto
endef

proto:
	$(call proto,$(PKG_DIR))