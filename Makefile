ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

gen_proto:
	@echo "Building docker image"
	docker build -t gen_proto -f gen.Dockerfile .
	@echo "Running docker to generate proto files"
	docker run -v "$(ROOT_DIR)"/src/proto-gen:/server/users/src/proto-gen gen_proto


gen_cert:
	@echo "Generating certificates"
	make cert -C src

gen: gen_proto gen_cert

gen_and_run:
	make gen
	docker compose -f docker-compose-dev.yaml up --build