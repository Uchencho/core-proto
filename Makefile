
.PHONY: proto
proto:
	# remove previous proto
	rm -rf ./generated/*

    # grpc
	$(eval DIR=accounts)
	protoc --go_opt=paths=source_relative --go_out="generated" "${DIR}"/*.proto
	protoc --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative \
	 				   --go-grpc_out="generated" --go-grpc_opt=paths=source_relative "${DIR}"/grpc.proto
