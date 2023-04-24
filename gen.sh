protoc -I${GOOGLEAPIS_DIR} -I${K8S_DIR} --include_imports --include_source_info \
	--descriptor_set_out=protos.pb \
	$(find ${K8S_DIR}/k8s.io -name "*.proto")
